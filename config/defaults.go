package config

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	if home, err = os.UserHomeDir(); err != nil {
		panic(err)
	}
	defOpts["logger"]["directory"] = home + "/.config/" + Title + "/logs/"

}

var (
	configSections = []string{"logger", "http", "performance", "deception", "ssh"}
	defNoColor     = false
)

var defOpts = map[string]map[string]interface{}{
	"logger": {
		"debug":             true,
		"trace":             false,
		"nocolor":           defNoColor,
		"use_date_filename": true,
	},
	"http": {
		"use_unix_socket":         false,
		"unix_socket_path":        "/var/run/hellpot",
		"unix_socket_permissions": "0666",
		"bind_addr":               "127.0.0.1",
		"bind_port":               "8080",
		"router": map[string]interface{}{
			"catchall":   false,
			"makerobots": true,
			"paths": []string{
				"wp-login.php",
				"wp-login",
			},
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

func setDefaults() {
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		snek.SetDefault("logger.directory", "./logs/")
		defNoColor = true
	}

	for _, def := range configSections {
		snek.SetDefault(def, defOpts[def])
	}

	if GenConfig {
		if err = snek.SafeWriteConfigAs("./config.toml"); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

}
