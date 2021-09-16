package extra

import (
	"bytes"
	"compress/gzip"
	crip "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/yunginnanet/HellPot/config"
)

const hellpot = "H4sIAAAAAAACA8WXMU/rMBDH93wFlhsYozbUBRTxBd7IXjGUUkGBpIgXnoT0hg4dssaV+gH7SUhzje04d45TqJBKZZy78/n/u7PT4GxyProZRwngQCS77bocXuzndtt8t5WmhZ7Ozybi+jAdJfs/MAw2dTQow9kRGVdqkSrSwRjgFPayHuis3G6bylRpxW+OFGZdf3JOb7UN7RwlASEtxtCOBbVHchXpRFZ66X0bBis3cSUoC6C/pf9WgRSOZpY3lGhv1BWoYU/mJumakSxQicN2pPKzwjmhHguBjwsqO2kh1EWgdO5ys2V2Azrek0erSTKFKjmYhRccdhdNYn26TFo8GngLLD9AjHsjZAgttHxzSbOErQPF7UhXRC+wR/qtm6ceHrUbSj0Osr1jfeSueO3bxUBCbZ65VFoFKSyJFmqOgDyEWk47qsY1r5x2TRr97boecr6aj/eUxlNKjnbxeStvNGx5pK2so9EO1FqnKzMWFFFr2IX64gBMgFPT1fFUKn564Iqbb2naN4Y0xP5hvLLuhcO36wWv4+AtmqfMsUg96cAp1CiYLH8tYlffMK/K3gnYfeGADz63CQ26UKA98LVAM7uiX4y89n8ib7oPuAvXt1oYgI5fPFSutRmagi+l7kodKSV6tCBdUifxEpHBAcA/oFPEQ1Sl07/a7TyZPtc6DsrBGCcXabWoWr36CgKYqBCPi2yQzd4GD4v3+Swbvkw/l3+Hf+avr7fLrDK/Q2NMrrR++rgfzJbJ8PMjfVyk6TSdZ037/2IUinEoLkNxFcYijEdhPA7jyzAu/42CLxeiz4pAEAAA"

func b64d(str string) []byte {
	var data []byte
	data, _ = base64.StdEncoding.DecodeString(str)
	return data
}
func rc(s []string) string {
	return strings.TrimSpace(s[ru()%uint32(len(s))])
}
func process(in string) (s string) {
	var v = strings.Split(config.Version, "")
	maj := v[0]
	min := v[2]
	sp := strings.Split(gz(b64d(in)), "|")
	s = sp[0]
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
	s = strings.Replace(s, "$maj", maj, -1)
	s = strings.Replace(s, "$min", min, -1)
	return
}
func gz(data []byte) string {
	gz, _ := gzip.NewReader(bytes.NewReader(data))
	out, _ := ioutil.ReadAll(gz)
	return string(out)
}
func ru() uint32 {
	b := make([]byte, 8192)
	crip.Read(b)
	return binary.LittleEndian.Uint32(b)
}

// PrintBanner prints our entropic banner
func PrintBanner() {
	println("\n" + process(hellpot) + "\n\n")
}
