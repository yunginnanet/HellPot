package config

import (
	"bytes"
	"runtime"
	"time"

	"github.com/knadh/koanf/parsers/toml"
)

var Defaults = &Preset{val: defOpts}

func init() {
	Defaults.IO = &PresetIO{p: Defaults}
}

type Preset struct {
	val map[string]interface{}
	IO  *PresetIO
}

type PresetIO struct {
	p   *Preset
	buf *bytes.Buffer
}

func (pre *Preset) ReadBytes() ([]byte, error) {
	return toml.Parser().Marshal(pre.val) //nolint:wrapcheck
}

func (shim *PresetIO) Read(p []byte) (int, error) {
	if shim.buf != nil && shim.buf.Len() > 0 {
		return shim.buf.Read(p) //nolint:wrapcheck
	}
	data, err := shim.p.ReadBytes()
	if err != nil {
		return 0, err
	}
	if shim.buf == nil {
		shim.buf = bytes.NewBuffer(data)
	}
	return shim.buf.Read(p) //nolint:wrapcheck
}

func (pre *Preset) Read() (map[string]interface{}, error) {
	return pre.val, nil
}

var defOpts = map[string]interface{}{
	"logger": map[string]interface{}{
		"debug":               true,
		"trace":               false,
		"nocolor":             runtime.GOOS == "windows",
		"noconsole":           false,
		"use_date_filename":   false,
		"docker_logging":      false,
		"rsyslog_address":     "",
		"console_time_format": time.Kitchen,
	},
	"http": map[string]interface{}{
		"use_unix_socket":         false,
		"unix_socket_path":        "/var/run/hellpot",
		"unix_socket_permissions": "0666",
		"bind_addr":               "127.0.0.1",
		"bind_port":               int64(8080), //nolint:gomnd
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
	"performance": map[string]interface{}{
		"restrict_concurrency": false,
		"max_workers":          256, //nolint:gomnd
	},
	"deception": map[string]interface{}{
		"server_name": "nginx",
	},
	"bespoke": map[string]interface{}{
		"grimoire_file":   "",
		"enable_grimoire": false,
	},
}
