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
	status := make([]int, jumlahModul)
	dependencies := [][]int{}
	fmt.Println(status)
	for i := 0; i < jumlahModul; i++ {
		
	}
	fmt.Println(dependencies)
	return result
}

func main() {
	jumlahModul := 4
	ketergantungan := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(UrutanDeploy(jumlahModul, ketergantungan))
}
