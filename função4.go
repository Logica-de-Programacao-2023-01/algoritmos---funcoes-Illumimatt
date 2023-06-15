package main

func segundoMaiorValor(slice []int) int {
	if len(slice) < 2 {
		panic("O slice precisa ter pelo menos 2 elementos.")
	}

	maior := slice[0]
	segundoMaior := slice[1]

	if segundoMaior > maior {
		maior, segundoMaior = segundoMaior, maior
	}

	for i := 2; i < len(slice); i++ {
		if slice[i] > maior {
			segundoMaior = maior
			maior = slice[i]
		} else if slice[i] > segundoMaior && slice[i] != maior {
			segundoMaior = slice[i]
		}
	}

	return segundoMaior
}
