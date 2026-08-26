package main

import "fmt"

func UrutanDeploy(jumlahModul int, ketergantungan [][]int) []int {
	result := []int{}
	if len(ketergantungan) == 0 {
		for i := 0; i < jumlahModul; i++ {
			result = append(result, i)
		}
		return result
	}
	
	return result
}

func main() {
	jumlahModul := 3
	ketergantungan := [][]int{}
	fmt.Println(UrutanDeploy(jumlahModul, ketergantungan))
}
