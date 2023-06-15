package main

import (
	"errors"
	"fmt"
)

type FuncaoAplicar func(int) int

func aplicarFuncao(slice []int, funcao FuncaoAplicar) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}

func dobrarValor(valor int) int {
	return valor * 2
}

func main() {
	valores := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncao(valores, dobrarValor)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)
}
