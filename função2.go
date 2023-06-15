package main

import (
	"strings"
)

func contarVogais(str string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	lowerStr := strings.ToLower(str)
	count := 0

	for _, char := range lowerStr {
		if contains(vogais, string(char)) {
			count++
		}
	}

	return count
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
