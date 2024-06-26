package config

import (
	"fmt"
	"io"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
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

	envK := koanf.New(".")

	envErr := envK.Load(env.Provider("HELLPOT_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "HELLPOT_")
		s = strings.ToLower(s)
		s = strings.ReplaceAll(s, "__", " ")
		s = strings.ReplaceAll(s, "_", ".")
		s = strings.ReplaceAll(s, " ", "_")
		return s
	}), nil)

	if envErr == nil && envK != nil && len(envK.All()) > 0 {
		if err := k.Merge(envK); err != nil {
			return nil, fmt.Errorf("failed to merge env config: %w", err)
		}
	}

	p := &Parameters{
		source: k,
	}

	if err := k.Unmarshal("", p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return p, nil
}
