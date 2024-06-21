package config

import (
	"bytes"
	"testing"
)

func TestDefaults(t *testing.T) {
	t.Run("ReadBytes", func(t *testing.T) {
		t.Parallel()
		bs, err := Defaults.ReadBytes()
		if err != nil {
			t.Fatal(err)
		}
		if len(bs) == 0 {
			t.Fatal("expected non-empty byte slice")
		}
		total := 0
		for _, needle := range []string{
			"logger",
			"http",
			"performance",
			"deception",
		} {
			total += bytes.Count(bs, []byte(needle)) + 3 // name plus brackets and newline
			if !bytes.Contains(bs, []byte(needle)) {
				t.Errorf("expected %q in byte slice", needle)
			}
		}
		if len(bs) <= total {
			t.Errorf("default byte slice seems too short to contain any default values")
		}
	})
	t.Run("Read", func(t *testing.T) {
		t.Parallel()
		m, err := Defaults.Read()
		if err != nil {
			t.Fatal(err)
		}
		if len(m) == 0 {
			t.Fatal("expected non-empty map")
		}
		for _, needle := range []string{
			"logger",
			"http",
			"performance",
			"deception",
		} {
			if _, ok := m[needle]; !ok {
				t.Errorf("expected %q in map", needle)
			}
			if m[needle] == nil {
				t.Errorf("expected non-nil value for %q", needle)
			}
		}
	})
}
