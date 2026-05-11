package differenceofsquares

func SquareOfSum(n int) int {
    var sum = ((1 + n) * n) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
    var sum = 0
    for i := 0; i <= n; i++ {
        sum += i * i
    }
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
