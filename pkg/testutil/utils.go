// Package testutil provides various utils to test generated code
package testutil

import (
	"fmt"
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
		panic(fmt.Sprintf("'%s' while formatting:\n%s", err, src))
	}
	return strings.TrimSpace(string(res))
}

// CreateGoFile creates temp file with given go code
func CreateGoFile(srcText string) *os.File {
	file, err := ioutil.TempFile("/tmp", "generic")
	if err != nil {
		panic(err)
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

// Text takes lines as varargs and returns a string with mulitline text.
// Just some sugar for writing tests with multiline inputs
func Text(lines ...string) string {
	return strings.Join(lines, "\n")
}
