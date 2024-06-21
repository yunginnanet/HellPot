package config

import "testing"

func TestCompileRules(t *testing.T) {
	if _, err := NewClientRules(nil, nil); err != nil {
		t.Error(err)
	}
	rules, err := NewClientRules([]string{"test", "test"}, nil)
	if err != nil {
		t.Error(err)
	}
	if len(rules.UseragentDisallowStrings) != 1 {
		t.Error("expected 1 got", len(rules.UseragentDisallowStrings))
	}
	if rules.UseragentDisallowStrings[0] != "test" {
		t.Error("expected test got", rules.UseragentDisallowStrings[0])
	}
	rules, err = NewClientRules(
		[]string{"yeeterson", "", "", "", "yeeterson", "mc", "mc", "geeterson"},
		[]string{"^y..terson$", "^mc", "^mc", "^g..ters.n$"},
	)
	if err != nil {
		t.Error(err)
	}
	if len(rules.useragentDisallowRegex) != 3 {
		t.Error("expected 3 got", len(rules.useragentDisallowRegex))
	}
	if len(rules.UseragentDisallowStrings) != 3 {
		t.Error("expected 3 got", len(rules.UseragentDisallowStrings))
	}
	if !rules.MatchUseragent("yeeterson") {
		t.Error("expected true got false")
	}
	if !rules.MatchUseragent("mc") {
		t.Error("expected true got false")
	}
	if !rules.MatchUseragent("yooterson") {
		t.Error("expected true got false")
	}
	if !rules.MatchUseragent("gooters%n") {
		t.Error("expected true got false")
	}
	if rules.MatchUseragent("yootersongooterson") {
		t.Error("expected false got true")
	}
}
