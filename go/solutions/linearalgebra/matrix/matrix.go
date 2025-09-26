package matrix

import "fmt"

type Vector struct {
	Data []float64
	Len  int
}

type Matrix struct {
	Data    [][]float64
	Numrows int
	Numcols int
}

func (left Matrix) Multiply(right Matrix) {
	var product Matrix
	product.Numcols = right.Numcols
	product.Numrows = left.Numrows
	//var scalarSum float64
	for i := 0; i < left.Numrows; i++ {
		//scalarSum = 0
		for j := 0; j < right.Numcols; j++ {
			for k := 0; k < left.Numcols; k++ {
				product.Data[i][j] += left.Data[i][k] * right.Data[k][j]
			}
		}
	}
}

func (Vector) PrintOscar() {
	fmt.Println("Oscar!!")
}

func Oscar_Add(a, b int) int {
	return a + b
}
