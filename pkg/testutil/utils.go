package testutil

import (
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func formatSrc(src string) string {
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

// AssertCodeIsSame asserts if two code snippets are equivalent
func AssertCodeIsSame(t *testing.T, expected string, actual string) {
	assert.Equal(t, formatSrc(expected), formatSrc(actual))
}
