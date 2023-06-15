package main

func calcularMedia(slice []int) float64 {
	total := 0
	for _, valor := range slice {
		total += valor
	}
	media := float64(total) / float64(len(slice))
	return media
}
