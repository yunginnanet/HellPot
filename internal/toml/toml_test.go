package toml

import (
	"errors"
	"strings"
	"testing"
)

var (
	wanted        = "----- Wanted -----"
	wantedErr     = "----- Wanted Err -----"
	actual        = "----- Actual -----"
	actualErr     = "----- Actual Err -----"
	divider       = strings.Repeat("-", len(wanted))
	divErr        = strings.Repeat("-", len(wantedErr))
	testResFmt    = "\n" + wanted + "\n%s\n" + actual + "\n%s\n" + divider + "\n"
	testResErrFmt = "\n" + wantedErr + "\n%s\n" + actualErr + "\n%s\n" + divErr + "\n"
)

type Test struct {
	Name            string
	UnmarshalInput  []byte
	UnmarshalOutput any
	MarshalInput    any
	MarshalOutput   []byte
	wantError       error
}

type Geets struct {
	Geeters string `toml:"geeters"`
}
type McGeeParameters struct {
	Username string `toml:"username"`
	SkipTag  string `toml:"-"`
	SubGeet  *Geets `toml:"sub_geet"`
}
type YeetersonParameters struct {
	ServerName string          `toml:"server_name"`
	DenyList   []string        `toml:"deny_list"`
	PortNumber int             `toml:"port_number"`
	YeetMode   bool            `toml:"yeet_mode"`
	McGee      McGeeParameters `toml:"mcgeet"`
	unexported string          `toml:"unexported"` //golint:unused
}
type Parameters struct {
	Yeeterson YeetersonParameters `toml:"yeet"`
	McGee     McGeeParameters     `toml:"mcgee"`
	SkipMe    string              `toml:"-"`

	skipMe string `toml:"skip_me"` //golint:unused
}

var simpleYeeterson = YeetersonParameters{
	ServerName: "yeeterson",
	DenyList:   []string{"yeet", "yeeterson", "yeeterson.com"},
	PortNumber: 8080,
	YeetMode:   true,
}

//nolint:funlen
func TestMarshalTOML(t *testing.T) {
	test1 := Test{
		Name: "simple",
		MarshalInput: Parameters{
			Yeeterson: simpleYeeterson,
			McGee: McGeeParameters{
				Username: "mcgee",
			},
		},
		MarshalOutput: []byte(`[yeet]
server_name = "yeeterson"
deny_list = ["yeet", "yeeterson", "yeeterson.com"]
port_number = 8080
yeet_mode = true

[mcgee]
username = "mcgee"`),
	}

	test2 := Test{
		Name: "with empty string, negative number, spaced strings, and punctuation",
		MarshalInput: Parameters{
			Yeeterson: YeetersonParameters{
				ServerName: "",
				DenyList:   []string{"yeet it."},
				PortNumber: -5,
				YeetMode:   false,
			},
			McGee: McGeeParameters{
				Username: "the yeet guy",
				SkipTag:  "skip me",
			},
			SkipMe: "skip me",
		},
		MarshalOutput: []byte(`[yeet]
server_name = ""
deny_list = ["yeet it."]
port_number = -5
yeet_mode = false

[mcgee]
username = "the yeet guy"`),
	}

	yeetersonWithChild := YeetersonParameters{
		ServerName: simpleYeeterson.ServerName,
		DenyList:   simpleYeeterson.DenyList,
		PortNumber: simpleYeeterson.PortNumber,
		YeetMode:   simpleYeeterson.YeetMode,
		McGee: McGeeParameters{
			Username: "Yeeterson McGeeterson",
		},
	}

	test3 := Test{
		Name: "with sub-structs",
		MarshalInput: Parameters{
			Yeeterson: yeetersonWithChild,
			McGee: McGeeParameters{
				Username: "mcgee",
			},
		},
		MarshalOutput: []byte(`[yeet]
server_name = "yeeterson"
deny_list = ["yeet", "yeeterson", "yeeterson.com"]
port_number = 8080
yeet_mode = true

[yeet.mcgeet]
username = "Yeeterson McGeeterson"

[mcgee]
username = "mcgee"`),
	}

	test4 := Test{
		Name: "with empty structs",
		MarshalInput: Parameters{
			Yeeterson: YeetersonParameters{},
			McGee: McGeeParameters{
				Username: "mcgeets",
				SubGeet:  &Geets{},
			},
		},
		MarshalOutput: []byte(`[mcgee]
username = "mcgeets"`),
	}

	test5 := Test{
		Name: "with pointer struct",
		MarshalInput: Parameters{
			Yeeterson: YeetersonParameters{},
			McGee: McGeeParameters{
				Username: "geetsies",
				SubGeet:  &Geets{Geeters: "geets"},
			},
		},
		MarshalOutput: []byte(`[mcgee]
username = "geetsies"

[mcgee.sub_geet]
geeters = "geets"`),
	}

	tests := []Test{test1, test2, test3, test4, test5}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output, err := MarshalTOML(test.MarshalInput)
			if !errors.Is(err, test.wantError) {
				errWantString := "nil"
				if test.wantError != nil {
					errWantString = test.wantError.Error()
				}
				t.Errorf(testResErrFmt, errWantString, err)
			}
			if string(output) != string(test.MarshalOutput) {
				t.Errorf(testResFmt, test.MarshalOutput, output)
			}
		})
	}
}
