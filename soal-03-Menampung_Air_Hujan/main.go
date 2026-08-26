package main

import (
	"fmt"
)

func TampungAir(tinggi []int) int {
	if len(tinggi) == 1 {
		return 0
	}
	tertampung := 0
	kiri := 0
	kanan := len(tinggi) - 1
	tinggiKiri := tinggi[0]
	tinggiKanan := tinggi[len(tinggi)-1]
	for kiri != kanan {
		if tinggiKiri <= tinggiKanan {
			kiri++
			if tinggiKiri < tinggi[kiri] {
				tinggiKiri = tinggi[kiri]
			}
			tertampung += tinggiKiri - tinggi[kiri]
		} else {
			kanan--
			if tinggiKanan < tinggi[kanan] {
				tinggiKanan = tinggi[kanan]
			}
			tertampung += tinggiKanan - tinggi[kanan]
		}
	}
	return tertampung
}

func main() {
	buildings := []int{1}
	fmt.Println(TampungAir(buildings))
}
