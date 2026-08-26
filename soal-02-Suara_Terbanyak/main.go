package main

import (
	"fmt"
)

func Mayoritas(suara []int) int {
	if len(suara) == 1 {
		return suara[0]
	}
	kandidat := -1
	jumlahSuara := 0
	for i := 0; i < len(suara); i++ {
		if jumlahSuara == 0 {
			kandidat = suara[i]

			jumlahSuara = 1
		} else if suara[i] == kandidat {
			jumlahSuara++
		} else {
			jumlahSuara--
		}
	}
	jumlahSuara = 0
	for i := 0; i < len(suara); i++ {
		if suara[i] == kandidat {
			jumlahSuara++
		}
	}
	if float64(jumlahSuara) > float64(len(suara))/2 {
		return kandidat
	}
	return kandidat
}

func main() {
	votes := []int{3, 2, 3}
	fmt.Println((Mayoritas(votes)))
}
