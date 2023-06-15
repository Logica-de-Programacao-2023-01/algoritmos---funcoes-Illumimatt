package main

import (
	"errors"
)

func filtrarNumerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	numerosPares := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			numerosPares = append(numerosPares, num)
		}
	}

	return numerosPares, nil
}
