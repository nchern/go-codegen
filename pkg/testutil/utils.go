package testutil

import (
	"go/format"
	"io/ioutil"
	"os"
)

func FormatSrc(src string) string {
	res, err := format.Source([]byte(src))
	if err != nil {
		panic(err)
	}
	return string(res)
}

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
