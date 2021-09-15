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
)

const hellpot = "H4sIAAAAAAACA8WXzU7CQBDH730FLnPw2EBlUWN8AY/eiQdEgkQLRosJiYceOPTaJeEBeRILy352ZrtFiQmQTTszO/P/zWxL1Ble9O8GSQpiwdLddl0tL/fXdttit+Wmhb5cdIbs5ng5SfdfMAw2MhpU4dyIhCu2ySHS0RjgHPZcLnRWfrfNwVRpRReHCrOWn4LSW5WhnZM0QqQVMbRjidWI7sK9yCovXbdhkPuJK0FJAO0tw0sFVDicWWEpUS/UF8iyR3PjeM9wEigXy3qk6pOLa0zdZkzcLrHsuINQN4HSucnNldkP6HRPGq0mSTQqp2CWQXDIKmxibaaMOzwsvKVoPxAY90aCIdTQ0sPFzRZ2DhS/I94RrcCe6Le2Tz1x1G4w9SjIbsX6yM1p7evNgEK1z1wsrRIVFkULkiMIHkxtpx3V4JqPnHpPGvPtezwUdDef7smNu5gc9eYLVt4Y2OpIy52j0Q1U26cpMxIU0mtiCvWDA0QClJq+icdSCdND7Lj5laZtY3BD7D/Gy+UsHH99L3gNB29pnzKnIg2kA+dQoySy/LeITXNDvCoHJ+DOhQc+hDxNcNClAh2ArwaaqAp/MQqq/0ze+BxQD9zQbiEAev7xYLlKM2EKoZSaO7WvlGgxgnhLncWLJQYHgPCAXhGPUaXNl/RKpIbdajHYL9hhN3tGIIpgqJyns6ybjd+7z7OPyTjrvY5Wi8/e/eTt7WGRHcwfhbFIq7J+WT51x4u0t1rOp7P5fDSfZLb9N7B+zAYxu4rZdXzLoh+7kUdDLBAAAA=="

func b64d(str string) []byte {
	var data []byte
	data, _ = base64.StdEncoding.DecodeString(str)
	return data
}
func rc(s []string) string {
	return strings.TrimSpace(s[ru()%uint32(len(s))])
}
func process(in string) (s string) {
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
