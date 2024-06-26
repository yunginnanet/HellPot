package config

import (
	"sync"

	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/internal/logger"
)

// Parameters represents the configuration for HellPot.
type Parameters struct {
	HTTP    HTTP                  `koanf:"http"`
	Logger  *logger.Configuration `koanf:"logger"`
	Bespoke Customization         `koanf:"bespoke"`
	Perf    Performance           `koanf:"performance"`
	Liar    Deception             `koanf:"deception"`

	IdleHands DevilsPlaythings `koanf:"experimental"`

	source        *koanf.Koanf    `koanf:"-"`
	logger        *zerolog.Logger `koanf:"-"`
	UsingDefaults bool            `koanf:"-"`
}

var once = &sync.Once{}

func (p *Parameters) GetLogger() *zerolog.Logger {
	once.Do(func() {
		p.logger = logger.GetLoggerOnce()
	})
	return p.logger
}

type Deception struct {
	// FakeServerName is our configured value for the "Server: " response header when serving HTTP clients
	FakeServerName string `koanf:"fake_server_name"`
}

type Performance struct {
	ConcurrencyCap bool `koanf:"limit_concurrency"`
	MaxWorkers     int  `koanf:"max_workers"`
}

// UnixSocket represents the configuration for the Unix socket.
type UnixSocket struct {
	// UnixSocketPath is the path to the Unix socket that HellPot will listen on if UseUnixSocket is set to true.
	UnixSocketPath string `koanf:"unix_socket_path"`
	// UseUnixSocket determines if we will listen for HTTP connections on a unix socket.
	UseUnixSocket bool `koanf:"use_unix_socket"`
	// UnixSocketPermissions are the octal permissions for the Unix socket.
	UnixSocketPermissions uint32 `koanf:"unix_socket_permissions"`
}

// Router represents the configuration for the HTTP router.
type Router struct {
	// Paths are defined via our toml configuration file. These are the paths that HellPot will present for "robots.txt"
	//    These are also the paths that HellPot will respond for. Other paths will throw a warning and will serve a 404.
	Paths       []string    `koanf:"paths"`
	CatchAll    bool        `koanf:"catchall"`
	MakeRobots  bool        `koanf:"makerobots"`
	ClientRules ClientRules `koanf:"client_rules"`
}

// HTTP represents the configuration for the HTTP server.
type HTTP struct {
	Bind string `koanf:"bind_addr"`
	Port int64  `koanf:"bind_port"`
	// ProxiedIPHeader is the HTTP Header containing the original IP of the client
	// for usage by traditional reverse Proxy deployments.
	ProxiedIPHeader string           `koanf:"proxied_ip_header"`
	Router          Router           `koanf:"router"`
	UnixSocket      UnixSocket       `koanf:"unix_socket"`
	ClientRules     ClientRules      `koanf:"client_rules"`
	Experimental    DevilsPlaythings `koanf:"experimental"`
}

// DevilsPlaythings - nothing to see here, move along.
type DevilsPlaythings struct {
	// POSTMimicry when true will cause HellPot to respond to POST requests to the configured roads to hell
	// with the content of the POST request entangled within the response. (Experimental)
	POSTMimicry bool `koanf:"post_mimicry"`
}

// Customization represents the configuration for the customizations.
type Customization struct {
	CustomHeffalump bool   `koanf:"custom_heffalump"`
	Grimoire        string `koanf:"grimoire"`
}
