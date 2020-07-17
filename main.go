package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	v := make([]float64, 81)
	v = []float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, //your sudoku matrix
		0, 0, 0, 0, 0, 3, 0, 8, 5,
		0, 0, 1, 0, 2, 0, 0, 0, 0,

		0, 0, 0, 5, 0, 7, 0, 0, 0,
		0, 0, 4, 0, 0, 0, 1, 0, 0,
		0, 9, 0, 0, 0, 0, 0, 0, 0,

		5, 0, 0, 0, 0, 0, 0, 7, 3,
		0, 0, 2, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 4, 0, 0, 0, 9,
	}
	// Create a new matrix
	A := mat.NewDense(9, 9, v)
	M := matrix{a: A}
	fmt.Println("All possible solution:")

	solve(&M)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

type matrix struct {
	a *mat.Dense
}

func (m *matrix) check(y, x int, num float64) bool {
	M := m.a
	for i := 0; i < 9; i++ {
		if M.At(y, i) == num {
			return false
		}
		if M.At(i, x) == num {
			return false
		}
	}
	x0 := x / 3 * 3
	y0 := y / 3 * 3
	for xi := x0; xi < (x0 + 3); xi++ {
		for yi := y0; yi < (y0 + 3); yi++ {
			if M.At(yi, xi) == num {
				return false
			}
		}
	}
	return true
}

func solve(m *matrix) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if m.a.At(y, x) == 0 {
				for i := 1; i < 10; i++ {
					if m.check(y, x, float64(i)) {
						m.a.Set(y, x, float64(i))
						solve(m)
						m.a.Set(y, x, float64(0))
					}
				}
				return
			}
		}
	}
	matPrint(m.a)
	return
}
