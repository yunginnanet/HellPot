package config

import (
	"bytes"
	"fmt"
	"regexp"
)

type ClientRules struct {
	// See: https://github.com/yunginnanet/HellPot/issues/23
	UseragentDisallowStrings  []string         `koanf:"user_agent_disallow_strings"`
	useragentDisallowStrBytes [][]byte         `koanf:"-"`
	UseragentDisallowRegex    []string         `koanf:"user_agent_disallow_regex"`
	useragentDisallowRegex    []*regexp.Regexp `koanf:"-"`
}

func NewClientRules(strs []string, regex []string) (*ClientRules, error) {
	if strs == nil && regex == nil {
		return &ClientRules{}, nil
	}
	if regex == nil {
		regex = make([]string, 0)
	}
	if strs == nil {
		strs = make([]string, 0)
	}
	cr := &ClientRules{
		UseragentDisallowStrings: strs,
		UseragentDisallowRegex:   regex,
	}
	return cr, cr.compile()
}

func (c *ClientRules) compile() error {
	dupes := make(map[string]struct{})
	for _, v := range c.UseragentDisallowRegex {
		if v == "" {
			continue
		}
		if _, ok := dupes[v]; ok {
			continue
		}
		dupes[v] = struct{}{}
		var compd *regexp.Regexp
		var err error
		if compd, err = regexp.Compile(v); err != nil {
			return fmt.Errorf("failed to compile regex '%s': %w", v, err)
		}
		c.useragentDisallowRegex = append(c.useragentDisallowRegex, compd)
	}

	newStrs := make([]string, 0)
	for _, v := range c.UseragentDisallowStrings {
		if v == "" {
			continue
		}
		if _, ok := dupes[v]; ok {
			continue
		}
		dupes[v] = struct{}{}
		newStrs = append(newStrs, v)
	}
	c.UseragentDisallowStrings = newStrs
	c.useragentDisallowStrBytes = make([][]byte, len(c.UseragentDisallowStrings))
	for i, v := range c.UseragentDisallowStrings {
		c.useragentDisallowStrBytes[i] = []byte(v)
	}
	return nil
}

func (c *ClientRules) MatchUseragent(ua []byte) bool {
	for _, v := range c.useragentDisallowRegex {
		if v.Match(ua) {
			return true
		}
	}
	for _, v := range c.useragentDisallowStrBytes {
		if bytes.Contains(ua, v) {
			return true
		}
	}
	return false
}
