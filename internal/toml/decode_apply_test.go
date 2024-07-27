package toml

import (
	"testing"
)

//nolint:funlen
func TestDecoder_readData(t *testing.T) {
	testData := test1.Marshaled
	d := newDecoder(testData, nil)
	if err := d.readData(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(d.tables) == 0 {
		t.Fatalf("expected tables, got none")
	}
	if d.tables["yeet"] == nil {
		t.Errorf("expected table 'yeet', got none")
	}

	var ok bool

	if _, ok = d.tables["yeet"].(map[string]interface{}); !ok {
		t.Errorf("expected table 'yeet' to be a map, got %T", d.tables["yeet"])
	}
	if _, ok = d.tables["yeet"].(map[string]interface{})["server_name"]; !ok {
		t.Errorf("expected table 'yeet.server_name', got none")
	}
	var sname string
	if sname, ok = d.tables["yeet"].(map[string]interface{})["server_name"].(string); !ok {
		//nolint:forcetypeassert
		t.Errorf("expected table 'yeet.server_name' to be a string, got %T",
			d.tables["yeet"].(map[string]interface{})["server_name"])
	}
	if sname != "yeeterson" {
		t.Errorf("expected table 'yeet.server_name' to be 'yeeterson', got %s", sname)
	}
	if _, ok = d.tables["yeet"].(map[string]interface{})["deny_list"]; !ok {
		t.Errorf("expected table 'yeet.deny_list', got none")
	}
	if _, ok = d.tables["yeet"].(map[string]interface{})["deny_list"].([]string); !ok {
		//nolint:forcetypeassert
		t.Errorf("expected table 'yeet.deny_list' to be a []string, got %T",
			d.tables["yeet"].(map[string]interface{})["deny_list"])
	}
	if _, ok = d.tables["yeet"].(map[string]interface{})["port_number"]; !ok {
		t.Errorf("expected table 'yeet.port_number', got none")
	}
	var pn int64
	if pn, ok = d.tables["yeet"].(map[string]interface{})["port_number"].(int64); !ok {
		//nolint:forcetypeassert
		t.Errorf("expected table 'yeet.port_number' to be an int, got %T",
			d.tables["yeet"].(map[string]interface{})["port_number"])
	}
	if pn != 8080 {
		t.Errorf("expected table 'yeet.port_number' to be 8080, got %d", pn)
	}
	if _, ok = d.tables["yeet"].(map[string]interface{})["yeet_mode"]; !ok {
		t.Errorf("expected table 'yeet.yeet_mode', got none")
	}
	var ym bool
	if ym, ok = d.tables["yeet"].(map[string]interface{})["yeet_mode"].(bool); !ok {
		//nolint:forcetypeassert
		t.Errorf("expected table 'yeet.yeet_mode' to be a bool, got %T",
			d.tables["yeet"].(map[string]interface{})["yeet_mode"])
	}
	if !ym {
		t.Errorf("expected table 'yeet.yeet_mode' to be true, got false")
	}
	if _, ok = d.tables["mcgee"]; !ok {
		t.Errorf("expected table 'mcgee', got none")
	}
	if _, ok = d.tables["mcgee"].(map[string]interface{})["username"]; !ok {
		t.Errorf("expected table 'mcgee.username', got none")
	}
	var uname string
	if uname, ok = d.tables["mcgee"].(map[string]interface{})["username"].(string); !ok {
		//nolint:forcetypeassert
		t.Errorf("expected table 'mcgee.username' to be a string, got %T",
			d.tables["mcgee"].(map[string]interface{})["username"])
	}
	if uname != "mcgee" {
		t.Errorf("expected table 'mcgee.username' to be 'mcgee', got %s", uname)
	}
}
