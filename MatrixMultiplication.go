package main

import (
	"fmt"
)

func MatrixMultiplicationM(a, b [][]int) {
	var c [3][3]int
	//multiplication of matrix
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			for k := 0; k < len(a[0]); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	//prints the result
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Print(c[i][j])
		}
		fmt.Println("")
	}
}
func main() {
	a := [][]int{
		[]int{1, 1, 1},
		[]int{2, 2, 2},
		[]int{3, 3, 3}}
	b := [][]int{
		[]int{1, 1, 1},
		[]int{2, 2, 2},
		[]int{3, 3, 3}}
	MatrixMultiplicationM(a, b)
}
