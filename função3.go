package main

import (
	"strings"
)

func concatenarStrings(slice []string) string {
	concatenado := strings.Join(slice, "")
	return concatenado
}
