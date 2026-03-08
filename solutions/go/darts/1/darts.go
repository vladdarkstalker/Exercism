package darts

import "math"

func Score(x, y float64) int {
    score := 0
    radius := math.Pow(math.Pow(x, 2) + math.Pow(y, 2), 0.5)

	if radius <= 1.0 {
        score = 10
    } else if radius <= 5.0 {
        score = 5
    } else if radius <= 10.0 {
        score = 1
    }

    return score
}
