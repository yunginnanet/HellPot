package extra

import (
	"bytes"
	"compress/gzip"
	crip "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/yunginnanet/HellPot/config"
)

const hellpot = "H4sIAAAAAAACA8VXvW7bQAze9QpZOGQUbNXnJBDyAh27Gx1c10icRHLRKgUCdPDgQavOgB/QT1JZ9P3oRJ5OTowAjnE5feSR30fy5Ohqdj25nyYZ4EJkh/22Xn457h325WEvbYTZLq9m4u60nWTHP7AAO+UNaneuR8aUOqTxdAIDXAIv1cJE5TfbNVDNFZ8cScxWfUqOb52GMU6yiKAWfRjDisqRPEV6JautTN4WYONXXBPKCjAcGZ4qkMTRmpUtJrqJ+hy18GRskq4ZyQoqcdn1VH82uCf0YyHwcUVFJx0JTRFonvvMXJr9Ap1vyUtrlGQKVXJiVkHisFm0FRvSZdLRoyVvheUHKOMRhBpCR1q+uaRdws5A8RvSFTFI2DPttu2ph6N2R7HHiexmbEbuhue+WwykqO2ZS4VVkcSS0oLSEVAPoY8zhrpx7SunW5NWf/uuh5Kv5vMtpfWUoqNbfMHMWw1bj7SNMxpdR51z+iJjhSJqDbvQXByAAXBs+jqeCiWMDzxx9y5Oh/qQFtkfLK9UvXD69r3g9Qzeqj1lzpU0UB24BBsVE+WneezrG+ZVOTgAty884kPIbUILXWmhA+TrCM1kRb8YBeV/IWu6D7gLN7RaGAE9v3ioWBUMoRCqUn+lTjQTA1qQLqmLWInE0gEg3KGXxJNXhfmrrK6z+ZOicVQvpri5yk2DNKvmK4pgpl08rIpRsfg1+rn6vVwU4+f52/rP+Ovy5eXbumjg3xGMwdXox9cfo8U6G7+95g+rPJ/ny6KN/ycmsZjG4iYWt3Eq4nQSp9M4vYnT+t8k+g/TQ9elQBAAAA=="

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
	s = strings.ReplaceAll(s, "$maj", maj)
	s = strings.ReplaceAll(s, "$min", min)
	return
}
func gz(data []byte) string {
	gz, err1 := gzip.NewReader(bytes.NewReader(data))
	out, err2 := io.ReadAll(gz)
	if err1 != nil || err2 != nil {
		bannerFail(err1, err2)
	}
	return string(out)
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
	println("\n" + process(hellpot) + "\n\n")
}
