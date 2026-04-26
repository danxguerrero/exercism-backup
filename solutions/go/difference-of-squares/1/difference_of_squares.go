package differenceofsquares
import (
    "math"
)



func SquareOfSum(n int) int {
    sum := 0
	for i := range n {
        sum += (i + 1)
    }

    squared := int(math.Pow(float64(sum), 2))
    
    return squared
}

func SumOfSquares(n int) int {
	sum := 0
    for i := range n {
        sum += int(math.Pow(float64(i + 1), 2))
    }

    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
