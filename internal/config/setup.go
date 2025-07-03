package config

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	flags "github.com/knadh/koanf/providers/basicflag"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"

	"github.com/yunginnanet/HellPot/internal/logger"
)

type readerProvider struct {
	source io.Reader
}

func (r *readerProvider) ReadBytes() ([]byte, error) {
	return io.ReadAll(r.source) //nolint:wrapcheck
}

func (r *readerProvider) Read() (map[string]interface{}, error) {
	b, err := r.ReadBytes()
	if err != nil {
		return nil, err
	}
	return toml.Parser().Unmarshal(b) //nolint:wrapcheck
}

func (p *Parameters) merge(ogk *koanf.Koanf, newk *koanf.Koanf, friendlyName string) error {
	if ogk == nil {
		panic("original koanf is nil")
	}
	if newk == nil {
		return nil
	}
	dirty := false

	newKeys := newk.All()
	ogKeys := ogk.All()

	if len(newk.All()) == 0 || len(newKeys) == 0 {
		// println("no configuration overrides found in " + friendlyName)
		return nil
	}

	valIsEmpty := func(v any) bool {
		if v == nil {
			return true
		}
		switch v.(type) {
		case string:
			return v.(string) == ""
		case []string:
			return len(v.([]string)) == 0
		default:
			return false
		}
	}

	for k, v := range ogKeys {
		// println("checking key " + k + " for overrides from " + friendlyName)
		if strings.HasPrefix(k, "logger") && !valIsEmpty(newKeys[k]) {
			lg := p.Logger
			lgptr := &lg
			// println("setting logger config key " + k + " to " + fmt.Sprintf("%v", newKeys[k]) + " from " + friendlyName)
			if err := lgptr.Set(k, newKeys[k]); err != nil {
				panic(fmt.Sprintf("failed to set logger config: %v", err))
			}
			p.Logger = *lgptr
			dirty = true
			continue
		}
		if valIsEmpty(v) && !valIsEmpty(newKeys[k]) {
			// println("setting newer value for key " + k + " to " + fmt.Sprintf("%v", newKeys[k]) + " from " + friendlyName)
			if err := ogk.Set(k, newKeys[k]); err != nil {
				panic(fmt.Sprintf("failed to set key %s: %v", k, err))
			}
			dirty = true
			continue
		}

		if _, hasDefault := Defaults.val[k]; !hasDefault {
			continue
		}

		if v == Defaults.val[k] && newKeys[k] != Defaults.val[k] {
			if err := ogk.Set(k, v); err != nil {
				panic(fmt.Sprintf("failed to set key %s: %v", k, err))
			}
			dirty = true
		}

	}

	if !dirty {
		return nil
	}

	// println("found configuration overrides in " + friendlyName)

	if err := ogk.Merge(newk); err != nil {
		return fmt.Errorf("failed to merge env config: %w", err)
	}

	return nil
}

func (p *Parameters) LoadEnv(k *koanf.Koanf) error {
	envK := koanf.New(".")

	envErr := envK.Load(env.Provider("HELLPOT_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "HELLPOT_")
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, "__", " ")
		s = strings.ReplaceAll(s, "_", ".")
		s = strings.ReplaceAll(s, " ", "_")
		return s
	}), nil)

	if envErr != nil {
		return fmt.Errorf("failed to load env: %w", envErr)
	}

	if err := p.merge(k, envK, "environment variables"); err != nil {
		return err
	}

	return nil
}

func parseCLISlice(key string, value string) (string, interface{}) {
	if _, ok := slicePtrs[key]; !ok {
		return key, value
	}
	split := strings.Split(value, ",")
	slices.Sort(split)
	return key, split
}

func (p *Parameters) LoadFlags(k *koanf.Koanf) error {
	flagsK := koanf.New(".")

	if err := flagsK.Load(flags.ProviderWithValue(CLIFlags, "-", parseCLISlice, k), nil); err != nil {
		return fmt.Errorf("failed to load flags: %w", err)
	}

	if err := p.merge(k, flagsK, "cli arguments"); err != nil {
		return err
	}

	return nil
}

func Setup(source io.Reader) (*Parameters, error) {
	k := koanf.New(".")

	if err := k.Load(Defaults, nil); err != nil {
		return nil, fmt.Errorf("failed to load defaults: %w", err)
	}

	if source != nil {
		if err := k.Load(&readerProvider{source}, toml.Parser()); err != nil {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	p := &Parameters{
		source: k,
		Logger: logger.Configuration{},
	}

	if source == nil {
		p.UsingDefaults = true
	}

	if err := p.LoadEnv(k); err != nil {
		return nil, err
	}

	if err := p.LoadFlags(k); err != nil {
		return nil, err
	}

	if err := k.Unmarshal("", p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	for k, v := range k.All() {
		if strings.HasPrefix(k, "logger.") {
			// println("setting logger config key " + k + " to " + fmt.Sprintf("%v", v))
			if err := p.Logger.Set(k, v); err != nil {
				return nil, fmt.Errorf("failed to set logger config: %w", err)
			}
		}
	}

	return p, nil
}
