package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/spf13/afero"
)

func init() {
	var err error
	if home, err = os.UserHomeDir(); err != nil {
		panic(err)
	}
	if len(defOpts) == 0 {
		panic("default options map is empty")
	}
	defOpts["logger"]["directory"] = path.Join(home, ".local", "share", Title, "logs")
	prefConfigLocation = path.Join(home, ".config", Title)
}

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
	target := fmt.Sprintf("%s.toml", Title)
	if err := snek.SafeWriteConfigAs("config.toml"); err != nil {
		print(err.Error())
		os.Exit(1)
	}
	var f afero.File
	var err error
	f, err = memfs.Open("config.toml")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	nf, err := os.Create(target) // #nosec G304
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if _, err = io.Copy(nf, f); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	print("default configuration successfully written to " + target)
	os.Exit(0)
}

func setDefaults() {
	memfs := afero.NewMemMapFs()
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		defNoColor = true
	}
	for _, def := range configSections {
		snek.SetDefault(def, defOpts[def])
	}
	if GenConfig {
		snek.SetFs(memfs)
		gen(memfs)
	}
}
