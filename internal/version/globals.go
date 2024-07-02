package version

import (
	"runtime/debug"
)

const HP = "HellPot"

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
