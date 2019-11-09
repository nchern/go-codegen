package testutil

import (
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

// FormatSrc formats go source
func FormatSrc(src string) string {
	res, err := format.Source([]byte(src))
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(res))
}

// CreateGoFile creates temp file with given go code
func CreateGoFile(srcText string) *os.File {
	file, err := ioutil.TempFile("/tmp", "generic")
	if err != nil {
		panic(err.Error())
	}
	if _, err = file.WriteString(srcText); err != nil {
		panic(err)
	}
	return file
}
