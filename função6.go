package main

import (
	"errors"
	"strings"
)

func concatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	concatenado := strings.Join(slice, ",")
	return concatenado, nil
}
