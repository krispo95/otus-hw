package Fibonacci

import (
	"math"
)

func CountFibonacciIteration(a int) int {
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	}
	prev := 0
	curr := 1

	for i := 2; i <= a; i++ {
		prev, curr = curr, curr+prev
	}
	return curr
}

func CountFibonacciGold(a int) int {
	f := (1 + math.Sqrt(5)) / 2
	fpow := CountPower(f, a)
	fn := fpow/math.Sqrt(5) + 0.5
	return int(fn)
}

func CountPower(base float64, power int) float64 {
	res := float64(1)
	for power > 1 {
		if power%2 == 1 {
			res *= base
		}
		base *= base
		power /= 2
	}
	if power > 0 {
		res *= base
	}

	return res
}

func CountFibonacciMatrix(a int) int {
	if a == 1 {
		return 0
	} else if a == 2 {
		return 1
	}
	matr := [2][2]int{{1, 1}, {1, 0}}
	res := [2][2]int{{1, 0}, {0, 1}}
	for a > 1 {
		if a%2 == 1 {
			res = MatrixMultiplication(res, matr)
		}
		matr = MatrixMultiplication(matr, matr)
		a /= 2
	}
	if a > 0 {
		res = MatrixMultiplication(res, matr)
	}
	return res[1][0]
}

func MatrixMultiplication(m1 [2][2]int, m2 [2][2]int) [2][2]int {
	res := [2][2]int{}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2); j++ {
			res[i][j] = 0
			for k := 0; k < 2; k++ {
				res[i][j] = res[i][j] + (m1[i][k] * m2[k][j])
			}
		}
	}
	return res
}
