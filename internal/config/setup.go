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

func normalizeMap(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		ogk := k
		k = strings.ToLower(k)

		var sslice []string
		var sliceOK bool

		if sslice, sliceOK = v.([]string); !sliceOK {
			goto justLower
		}
		for i, s := range sslice {
			sslice[i] = strings.ToLower(s)
		}
		slices.Sort(sslice)
		m[k] = sslice
	justLower:
		if k != ogk {
			delete(m, ogk)
		}
	}
	return m
}

func (p *Parameters) merge(ogk *koanf.Koanf, newk *koanf.Koanf, friendlyName string) error {
	if ogk == nil {
		panic("original koanf is nil")
	}
	if newk == nil {
		return nil
	}
	dirty := false

	newKeys := normalizeMap(newk.All())

	if len(newk.All()) == 0 || len(newKeys) == 0 {
		return nil
	}

	valIsEmpty := func(v any) bool {
		switch v.(type) {
		case string:
			return v.(string) == ""
		case []string:
			return len(v.([]string)) == 0
		default:
			return false
		}
	}

	for k, v := range newKeys {
		ogv := ogk.Get(k)
		if ogk.Exists(k) && valIsEmpty(ogv) && !valIsEmpty(v) {
			println("setting newer value for key " + k + " to " + fmt.Sprintf("%v", v) + " from " + friendlyName)
			if err := ogk.Set(k, v); err != nil {
				panic(fmt.Sprintf("failed to set key %s: %v", k, err))
			}
			dirty = true
			continue
		}

		if _, hasDefault := Defaults.val[k]; !hasDefault {
			continue
		}

		if ogv == Defaults.val[k] && v != ogv {
			if err := ogk.Set(k, v); err != nil {
				panic(fmt.Sprintf("failed to set key %s: %v", k, err))
			}
			dirty = true
		}
	}

	if !dirty {
		return nil
	}

	println("found configuration overrides in " + friendlyName)

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

	if err := flagsK.Load(flags.ProviderWithValue(CLIFlags, ".", parseCLISlice), nil); err != nil {
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
	}

	if source == nil {
		p.UsingDefaults = true
	}

	if err := p.LoadFlags(k); err != nil {
		return nil, err
	}

	if err := p.LoadEnv(k); err != nil {
		return nil, err
	}

	if err := k.Unmarshal("", p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// for some reason the logger config pointer is not getting populated, do it manually
	p.Logger = logger.Configuration{}
	for k, v := range k.All() {
		if strings.HasPrefix(k, "logger.") {
			println("setting logger config key " + k + " to " + fmt.Sprintf("%v", v))
			if err := p.Logger.Set(k, v); err != nil {
				return nil, fmt.Errorf("failed to set logger config: %w", err)
			}
		}
	}

	return p, nil
}
