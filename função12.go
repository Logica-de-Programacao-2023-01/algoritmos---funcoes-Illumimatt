package main

import (
	"errors"
	"strings"
	"unicode"
)

func stringsComLetraMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	var resultado strings.Builder
	for _, palavra := range slice {
		if len(palavra) > 0 && unicode.IsUpper([]rune(palavra)[0]) {
			resultado.WriteString(palavra)
			resultado.WriteString(" ")
		}
	}

	return resultado.String(), nil
}
