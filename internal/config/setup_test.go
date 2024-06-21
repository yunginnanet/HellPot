package config

import (
	"bytes"
	"testing"
)

func TestSetup(t *testing.T) {
	t.Run("Success", SetupSuccess)
	t.Run("NoFailureOnNilSource", SetupNoFailureOnNilSource)
	t.Run("FailureOnReadConfig", SetupFailureOnReadConfig)
}

func SetupSuccess(t *testing.T) {
	source := bytes.NewBufferString(`
[http]
port = 55
bind_addr = "5.5.5.5"

[http.router]
catchall = true
makerobots = false
`)

	params, err := Setup(source)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if params == nil {
		t.Fatal("Expected params to be not nil")
	}

	if params.source.Get("http.port") != int64(55) {
		t.Errorf("Expected 55, got (%T) %v", params.source.Get("http.port"), params.source.Get("http.port"))
	}
	if params.HTTP.Port != int64(55) {
		t.Errorf("Expected 55, got %v", params.HTTP.Port)
	}
	if params.source.Get("http.bind_addr") != "5.5.5.5" {
		t.Errorf("Expected 5.5.5.5, got %v", params.source.Get("http.bind_addr"))
	}
	if params.HTTP.Bind != "5.5.5.5" {
		t.Errorf("Expected 5.5.5.5, got %v", params.HTTP.Bind)
	}
	if params.source.Get("http.router.catchall") != true {
		t.Errorf("Expected true, got %v", params.source.Get("http.router.catchall"))
	}
	if params.HTTP.Router.CatchAll != true {
		t.Errorf("Expected true, got %v", params.HTTP.Router.CatchAll)
	}

}

func SetupNoFailureOnNilSource(t *testing.T) {
	params, err := Setup(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if params == nil {
		t.Fatal("Expected params to be not nil")
	}

	t.Run("DefaultsWithNilSource", func(t *testing.T) {
		for _, needle := range []string{
			"logger",
			"http",
			"performance",
			"deception",
			"http.router",
			"http.router.paths",
		} {
			if params.source.Get(needle) == nil {
				t.Errorf("Expected %q in map", needle)
			}
		}

		// nolint:forcetypeassert
		if params.HTTP.Port != Defaults.val["http"].(map[string]interface{})["port"].(int64) {
			t.Errorf("Expected %v, got %v",
				// nolint:forcetypeassert
				Defaults.val["http"].(map[string]interface{})["port"].(int64), params.HTTP.Port,
			)
		}
		// nolint:forcetypeassert
		if params.HTTP.Bind != Defaults.val["http"].(map[string]interface{})["bind_addr"].(string) {
			t.Errorf("Expected %v, got %v",
				// nolint:forcetypeassert
				Defaults.val["http"].(map[string]interface{})["bind_addr"].(string), params.HTTP.Bind,
			)
		}
		if params.HTTP.Router.CatchAll !=
			// nolint:forcetypeassert
			Defaults.val["http"].(map[string]interface{})["router"].(map[string]interface{})["catchall"].(bool) {
			t.Errorf("Expected %v, got %v",
				// nolint:forcetypeassert
				Defaults.val["http"].(map[string]interface{})["router"].(map[string]interface{})["catchall"].(bool),
				params.HTTP.Router.CatchAll,
			)
		}
		if len(params.HTTP.Router.Paths) !=
			// nolint:forcetypeassert
			len(Defaults.val["http"].(map[string]interface{})["router"].(map[string]interface{})["paths"].([]string)) {
			t.Errorf("Expected %v, got %v",
				// nolint:forcetypeassert
				Defaults.val["http"].(map[string]interface{})["router"].(map[string]interface{})["paths"].([]string),
				params.HTTP.Router.Paths,
			)
		}
	})
}

func SetupFailureOnReadConfig(t *testing.T) {
	source := bytes.NewBufferString(`{eeeeeeeeeeeeeeeeeeEE: 1}`)

	params, err := Setup(source)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if params != nil {
		t.Error("Expected params to be nil")
	}
}
