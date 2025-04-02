package toml

import (
	"bytes"
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

func check(test1 test, t *testing.T) {
	t.Helper()
	t.Logf("running test: %s", test1.Name)
	testData := test1.Marshaled
	testTarget := &TestParameters{}
	d := newDecoder(testData, testTarget)
	if err := d.readData(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if testTarget.Yeeterson.ServerName == "yeeterson" {
		t.Fatalf("shouldn't be applied yet: %s", testTarget.Yeeterson.ServerName)
	}

	if err := d.handleTables(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// spew.Dump(d.tables)

	if testTarget.Yeeterson.ServerName != test1.Unmarshaled.(TestParameters).Yeeterson.ServerName {
		t.Errorf("expected '%s', got '%s'",
			test1.Unmarshaled.(TestParameters).Yeeterson.ServerName,
			testTarget.Yeeterson.ServerName,
		)
	}

	expectLen := len(test1.Unmarshaled.(TestParameters).Yeeterson.DenyList)

	if expectLen != len(test1.Unmarshaled.(TestParameters).Yeeterson.DenyList) {
		t.Fatalf("expected %d, got %d",
			len(test1.Unmarshaled.(TestParameters).Yeeterson.DenyList),
			len(testTarget.Yeeterson.DenyList))
	}

	if expectLen > 0 && testTarget.Yeeterson.DenyList[0] != test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[0] {
		t.Errorf("expected '%s', got '%s'",
			test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[0],
			testTarget.Yeeterson.DenyList[0],
		)
	}

	if expectLen > 1 && testTarget.Yeeterson.DenyList[1] != test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[1] {
		t.Errorf("expected '%s', got '%s'",
			test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[1],
			testTarget.Yeeterson.DenyList[1],
		)
	}

	if expectLen > 2 && testTarget.Yeeterson.DenyList[2] != test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[2] {
		t.Errorf("expected '%s', got '%s'",
			test1.Unmarshaled.(TestParameters).Yeeterson.DenyList[2],
			testTarget.Yeeterson.DenyList[2],
		)
	}

	if testTarget.Yeeterson.PortNumber != test1.Unmarshaled.(TestParameters).Yeeterson.PortNumber {
		t.Errorf("expected %d, got %d",
			test1.Unmarshaled.(TestParameters).Yeeterson.PortNumber,
			testTarget.Yeeterson.PortNumber,
		)
	}

	if test1.Unmarshaled.(TestParameters).McGee.SubGeet != nil {
		if testTarget.McGee.SubGeet == nil {
			t.Fatalf("expected sub_geet, got none")
		}
		if testTarget.McGee.SubGeet.Geeters != test1.Unmarshaled.(TestParameters).McGee.SubGeet.Geeters {
			t.Errorf("expected '%s', got '%s'",
				test1.Unmarshaled.(TestParameters).McGee.SubGeet.Geeters,
				testTarget.McGee.SubGeet.Geeters,
			)
		}
	}

	m, e := MarshalTOML(testTarget)
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}

	if !bytes.Equal(m, testData) {
		t.Errorf("expected %s, got %s", testData, m)
	}
}

func TestDecoder_handleTables(t *testing.T) {
	check(test1, t)
	check(test2, t)
	check(test3, t)
	check(test4, t)
}
