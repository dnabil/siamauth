package siamauth

import "strings"

func trimSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}