package config

import (
	"io"
	"os"
	"runtime"
	"time"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/spf13/afero"
)

var (
	configSections = []string{"logger", "http", "performance", "deception", "ssh"}
	defNoColor     = false
)

var defOpts = map[string]map[string]interface{}{
	"logger": {
		"debug":               true,
		"trace":               false,
		"nocolor":             defNoColor,
		"use_date_filename":   true,
		"docker_logging":      false,
		"console_time_format": time.Kitchen,
	},
	"http": {
		"use_unix_socket":         false,
		"unix_socket_path":        "/var/run/hellpot",
		"unix_socket_permissions": "0666",
		"bind_addr":               "127.0.0.1",
		"bind_port":               "8080",
		"real_ip_header":          "X-Real-IP",

		"router": map[string]interface{}{
			"catchall":   false,
			"makerobots": true,
			"paths": []string{
				"wp-login.php",
				"wp-login",
			},
		},
		"uagent_string_blacklist": []string{
			"Cloudflare-Traffic-Manager",
		},
	},
	"performance": {
		"restrict_concurrency": false,
		"max_workers":          256,
	},
	"deception": {
		"server_name": "nginx",
	},
}

func gen(memfs afero.Fs) {
	var (
		dat []byte
		err error
		f   afero.File
	)
	if dat, err = snek.Marshal(toml.Parser()); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if f, err = memfs.Create("config.toml"); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	var n int
	if n, err = f.Write(dat); err != nil || n != len(dat) {
		if err == nil {
			err = io.ErrShortWrite
		}
		println(err.Error())
		os.Exit(1)
	}
	println("Default config written to config.toml")
	os.Exit(0)
}

func setDefaults() {
	memfs := afero.NewMemMapFs()
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		defNoColor = true
	}
	for _, def := range configSections {
		for key, val := range defOpts[def] {
			if _, ok := val.(map[string]interface{}); !ok {
				if err := snek.Set(def+"."+key, val); err != nil {
					println(err.Error())
					os.Exit(1)
				}
				continue
			}
			for k, v := range val.(map[string]interface{}) {
				if err := snek.Set(def+"."+key+"."+k, v); err != nil {
					println(err.Error())
					os.Exit(1)
				}
			}
			continue
		}
	}

	if GenConfig {
		gen(memfs)
	}
}
