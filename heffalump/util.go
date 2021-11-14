package heffalump

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func gz(data []byte) string {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	out, err := io.ReadAll(gz)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func b64d(str string) []byte {
	var data []byte
	data, _ = base64.StdEncoding.DecodeString(str)
	return data
}
