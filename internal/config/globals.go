package config

import "runtime/debug"

// Title is the name of the application used throughout the configuration process.
const Title = "HellPot"

var Version = "dev"

func init() {
	if Version != "dev" {
		return
	}
	binInfo := make(map[string]string)
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, v := range info.Settings {
		binInfo[v.Key] = v.Value
	}
	if gitrev, ok := binInfo["vcs.revision"]; ok {
		Version = gitrev[:7]
	}
}

var (
	// BannerOnly when toggled causes HellPot to only print the banner and version then exit.
	BannerOnly = false
	// GenConfig when toggled causes HellPot to write its default config to the cwd and then exit.
	GenConfig = false
	// NoColor when true will disable the banner and any colored console output.
	NoColor bool
	// MakeRobots when false will not respond to requests for robots.txt.
	MakeRobots bool
	// CatchAll when true will cause HellPot to respond to all paths.
	// Note that this will override MakeRobots.
	CatchAll bool
)

// "http"
var (
	// HTTPBind is defined via our toml configuration file. It is the address that HellPot listens on.
	HTTPBind string
	// HTTPPort is defined via our toml configuration file. It is the port that HellPot listens on.
	HTTPPort string
	// HeaderName is defined via our toml configuration file. It is the HTTP Header containing the original IP of the client,
	// in traditional reverse Proxy deplyoments.
	HeaderName string

	// Paths are defined via our toml configuration file. These are the paths that HellPot will present for "robots.txt"
	//       These are also the paths that HellPot will respond for. Other paths will throw a warning and will serve a 404.
	Paths []string

	// UseUnixSocket determines if we will listen for HTTP connections on a unix socket.
	UseUnixSocket bool

	// UnixSocketPath is defined via our toml configuration file. It is the path of the socket HellPot listens on
	// if UseUnixSocket, also defined via our toml configuration file, is set to true.
	UnixSocketPath        = ""
	UnixSocketPermissions uint32

	// UseragentBlacklistMatchers contains useragent matches checked for with strings.Contains() that
	// prevent HellPot from firing off.
	// See: https://github.com/yunginnanet/HellPot/issues/23
	UseragentBlacklistMatchers []string
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
