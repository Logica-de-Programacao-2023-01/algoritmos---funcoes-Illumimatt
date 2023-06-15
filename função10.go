package main

import (
	"errors"
	"sort"
)

func ordenarValores(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	valoresOrdenados := make([]int, len(slice))
	copy(valoresOrdenados, slice)

	sort.Ints(valoresOrdenados)
	return valoresOrdenados, nil
}
