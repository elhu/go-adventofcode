package math

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GaussianElimination(mat [][]float64) [][]float64 {
	for i := 0; i < len(mat); i++ {
		factor := mat[i][i]
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] /= factor
		}
		for j := i + 1; j < len(mat); j++ {
			factor = mat[j][i]
			for k := 0; k < len(mat[j]); k++ {
				mat[j][k] -= factor * mat[i][k]
			}
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			factor := mat[j][i]
			for k := 0; k < len(mat[j]); k++ {
				mat[j][k] -= factor * mat[i][k]
			}
		}
	}
	return mat
}
