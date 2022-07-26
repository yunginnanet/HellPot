package extra

import (
	crip "crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"git.tcp.direct/kayos/common/squish"

	"github.com/yunginnanet/HellPot/config"
)

const hellpot = "H4sIAAAAAAACA8VXvW7bQAze9QpZOGQUZNXntBD6Ahm7Gx1cx0jdRnKRKAUCdPDgQavOgB/QTxLZ1P3oRJ5Obo0CtnE5feSR30fylOhmfjv9PEtzwIXIj4dds/xw2jsequNB2gizXd3Mxad2O81PX7AAe+UNGneuR8aUOuTsqQUDXAMv1cJE5Tfbn6GaKz45kpid+lQc3zoNY5zmEUEt+jCGNZUjeYr0StZYmbwtwNavuCaUFWA8MjxVIImjNas6TPQT9Tnq4MnYJF0zkhVU4rLvqflscU/ox0Lg45qKTjoSmiLQPA+ZuTT7BbrckpfWKMkUquTErIPEYbPoKjamy6SjR0feGssPUMYTCDWEnrR8c0m7hJ2B4jekK2KUsBfa7bpTD0ftnmKPE9nN2IzcLc99vxhIUbszlwqrJoklpQWlI6AeQh9nDHXj2ldOvyat/vZdDxVfzZdbSuspRUe/+IKZtxq2GWlbZzS6jnrnDEXGCkXUGnahuTgAA+DY9HU8FUoYH3ji/q84HetDWmT/Y3ml6oX21/eCNzB46+6UuVTSQHXgGmzUTJT/zeNQ3zCvysEBuH3hER9CbhNa6FoLHSBfT2gmK/rFKCj/K1nTfcBduKHVwgjo+Y+HilXBEAqhKg1X6lQzMaIF6ZK6ipVILR0Awh16SWy9KsxvZXWbL34oGpNmMcPNdYFmiE40+qV9cg4Logjm2uXjukzK5a/kYf28WpaTn4u3zcvkfvX09GVTnuFfEYzBNujvr9+S5SafvL0Wj+uiWBSrsov/I6axmMXiLhYf40zE2TTOZnF2F2fNn2n0DpcvBxhQEAAA"

func rc(s []string) string {
	return strings.TrimSpace(s[ru()%uint32(len(s))])
}

func process(in string) (s string) {
	var v = strings.Split(config.Version, "")
	maj := v[0]
	min := v[2]
	smin := v[4]
	defl8, _ := squish.UnpackStr(in)
	sp := strings.Split(defl8, "|")
	s = sp[0]
	if smin == "" {
		s = strings.ReplaceAll(s, "$1;40m.", "$1;40m")
	}
	c := strings.Split(sp[1], ",")
	cproc := func(in, num string) (inr string) {
		inr = in
		tor := fmt.Sprintf("$%s", num)
		for n := strings.Count(inr, tor); n > 0; n-- {
			inr = strings.Replace(inr, tor, rc(c), 1)
		}
		return
	}
	for n := 1; n < 5; n++ {
		s = cproc(s, fmt.Sprintf("%d", n))
	}

	s = strings.ReplaceAll(s, "$maj", maj)
	s = strings.ReplaceAll(s, "$min", min)
	s = strings.ReplaceAll(s, "$smin", smin)
	return
}

func ru() uint32 {
	b := make([]byte, 8192)
	if _, err := crip.Read(b); err != nil {
		bannerFail(err)
	}
	return binary.LittleEndian.Uint32(b)
}

// PrintBanner prints our entropic banner
func PrintBanner() {
	time.Sleep(5 * time.Millisecond)
	defer time.Sleep(5 * time.Millisecond)
	println("\n" + process(hellpot) + "\n\n")
}
