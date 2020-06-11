package impl

import "strings"

func text(lines ...string) string {
	return strings.Join(lines, "\n")
}
