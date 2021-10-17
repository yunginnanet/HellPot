package extra

import (
	"os"
	"runtime"

	"github.com/yunginnanet/HellPot/config"
)

func bannerFail(errs ...error) {
	println("failed printing banner, consider using --nocolor")
	for _, err := range errs {
		if err != nil {
			println(err.Error())
		}
	}
	os.Exit(1)
}

// Banner prints out our banner (using spooky magic)
func Banner() {
	if runtime.GOOS == "windows" || config.NoColor {
		println(config.Title + " " + config.Version)
		println(" ")
		return
	}
	PrintBanner()
}
