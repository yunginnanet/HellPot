package heffalump

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func gz(data []byte) string {
	gz, _ := gzip.NewReader(bytes.NewReader(data))
	out, _ := ioutil.ReadAll(gz)
	return string(out)
}

func b64d(str string) []byte {
	var data []byte
	data, _ = base64.StdEncoding.DecodeString(str)
	return data
}
