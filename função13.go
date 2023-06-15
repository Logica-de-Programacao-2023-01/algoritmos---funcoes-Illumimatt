package main

import (
	"errors"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número é negativo")
	}

	if numero == 0 {
		return 0, nil
	}

	soma := 0
	for numero > 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}
