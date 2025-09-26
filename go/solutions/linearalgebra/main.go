package main

import (
	"fmt"
	"linearalgebra/matrix"
)

func main() {
	fmt.Println("vim-go")
	sum := matrix.Oscar_Add(5, 9)
	fmt.Println("\n\n\n\nsum: ", sum, "\n\n")

	var v matrix.Vector
	v.Len = 19
	fmt.Println(v.Len)
	v.PrintOscar()

	var A, B matrix.Matrix

	A.Numrows = 2
	A.Numcols = 2
	B.Numrows = 2
	B.Numcols = 2
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			A.Data[i][j] = float64((i + 2) * 2)
			B.Data[i][j] = float64((i + 3) * 2)
		}
	}

	C := A.Multiply(B)

	fmt.Println("\n\nexperiment: ", C.Data[1][1])
}
