package extra

import (
	"runtime"
	"github.com/yunginnanet/HellPot/config"
)

// Banner prints out our banner (using spooky magic)
func Banner() {
	if runtime.GOOS == "windows" || config.NoColor {
		println(config.Title + " " + config.Version)
		println(" ")
		return
	}
	PrintBanner()
}
