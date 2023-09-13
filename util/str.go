package util

import "strings"

func TrimSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}