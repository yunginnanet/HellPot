package config

const (
	// Version roughly represents the applications current version.
	Version = "0.3"
	// Title is the name of the application used throughout the configuration process.
	Title = "HellPot"
)

var (
	// BannerOnly when toggled causes HellPot to only print the banner and version then exit.
	BannerOnly = false
	// GenConfig when toggled causes HellPot to write its default config to the cwd and then exit.
	GenConfig = false
	// NoColor when true will disable the banner and any colored console output.
	NoColor bool
)

// "http"
var (
	// HTTPBind is defined via our toml configuration file. It is the address that HellPot listens on.
	HTTPBind string
	// HTTPPort is defined via our toml configuration file. It is the port that HellPot listens on.
	HTTPPort string
	// Paths are defined via our toml configuration file. These are the paths that HellPot will present for "robots.txt"
	//       These are also the paths that HellPot will respond for. Other paths will throw a warning and will serve a 404.
	Paths []string

	// UseUnixSocket determines if we will listen for HTTP connections on a unix socket.
	UseUnixSocket bool

	// UnixSocketPath is defined via our toml configuration file. It is the path of the socket HellPot listens on
	// if UseUnixSocket, also defined via our toml configuration file, is set to true.
	UnixSocketPath        = ""
	UnixSocketPermissions uint32
)

// "performance"
var (
	RestrictConcurrency bool
	MaxWorkers          int
)

// "deception"
var (
	// FakeServerName is our configured value for the "Server: " response header when serving HTTP clients
	FakeServerName string
)
