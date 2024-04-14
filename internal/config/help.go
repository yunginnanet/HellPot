package config

import (
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

type help struct {
	title, version string
	usage          map[int][]string
	out            io.Writer
}

var CLI = help{
	title:   Title,
	version: Version,
	usage: map[int][]string{
		0: {0: "-c, --config", 1: "<file>", 2: "Specify config file"},
		1: {0: "--nocolor", 1: "disable color and banner"},
		2: {0: "--banner", 1: "show banner + version and exit"},
		3: {0: "--genconfig", 1: "write default config to " + Title + ".toml then exit"},
		4: {0: "-g, --grimoire", 1: "<file>", 2: "Specify a custom file used for text generation"},
		5: {0: "-h,--help", 1: "show this help and exit"},
	},
	out: os.Stdout,
}

func (cli help) secondColStart(index int) (max int) {
	l := cli.firstColEnd() + 2
	if len(cli.usage[index]) > 2 && cli.usage[index][2] != "" {
		l -= len(cli.usage[index][1])
	}
	if l > max {
		max = l
	}
	return max
}

func (cli help) firstColEnd() (max int) {
	for n := range cli.usage {
		l := len(cli.usage[n][0])
		if l > max {
			max = l
		}
	}
	return max
}

func (cli help) stdout(s ...string) {
	for _, v := range s {
		_, _ = cli.out.Write([]byte(v))
	}
}

func (cli help) lb(n int) {
	for n > 0 {
		cli.stdout("\n")
		n--
	}
}

func (cli help) printUsage() {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		os.Exit(1)
	}
	cli.header()

	for n := 0; n < len(cli.usage); n++ {
		line := &strings.Builder{}
		buf := &strings.Builder{}
		usageAt := 1
		tlen := cli.secondColStart(n)
		switch {
		case cli.usage[n][0] == "":
			cli.lb(1)
		case cli.usage[n][1] == "":
			cli.stdout(cli.usage[n][0])
			cli.lb(2)
		case len(cli.usage[n]) > 2 && cli.usage[n][2] != "":
			tlen = cli.firstColEnd() - len(cli.usage[n][1])
			usageAt = 2
			fallthrough
		default:
			buf.WriteString(cli.usage[n][0])
		}
		if tlen < 0 {
			tlen = 2
		}
		tab := strings.Repeat(" ", tlen)
		line.WriteString(" ")
		if buf.Len() < cli.firstColEnd() {
			line.WriteString(strings.Repeat(" ", cli.firstColEnd()-buf.Len()))
		}
		if usageAt == 2 {
			buf.WriteString(strings.Repeat(" ", tlen/2))
			buf.WriteString(cli.usage[n][1])
		}
		buf.WriteString(tab)
		buf.Write([]byte(" (" + cli.usage[n][usageAt] + ")"))
		buf.Write([]byte{'\n'})
		line.Write([]byte(buf.String()))
		cli.stdout(line.String())
	}
	os.Exit(0)

}

func (cli help) header() {
	cli.stdout("\n")
	s := &strings.Builder{}
	s.Write([]byte(cli.title))
	s.Write([]byte(" v["))
	s.Write([]byte(cli.version))
	s.Write([]byte("]"))
	tab := cli.firstColEnd() - (s.Len() % 2) + 1
	if tab > 0 {
		cli.stdout(strings.Repeat(" ", tab))
	}
	cli.stdout(s.String())
	cli.lb(2)
}
