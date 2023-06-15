package main

import (
	"errors"
	"strings"
)

func extrairPalavras(texto string) ([]string, error) {
	if len(texto) == 0 {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Split(texto, " ")
	return palavras, nil
}
