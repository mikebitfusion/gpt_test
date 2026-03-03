package gpt

import "strings"

func String(fields []string) string {
	return strings.Join(fields, "|")
}
