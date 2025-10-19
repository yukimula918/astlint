package util

import "strings"

func IsBlankString(s string) bool {
	return strings.TrimSpace(s) == ""
}
