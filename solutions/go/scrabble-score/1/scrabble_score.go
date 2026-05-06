package scrabblescore

import "strings"

func Score(word string) int {
    score := 0
	for _, el := range strings.ToLower(word) {
        if el == 'k' {
            score += 5
        } else if el == 'd' || el == 'g' {
            score += 2
        } else if el == 'j' || el == 'x' {
            score += 8
        } else if el == 'q' || el == 'z' {
            score += 10
        } else if el == 'b' || el == 'c' || el == 'm' || el == 'p' {
            score += 3
        } else if el == 'f' || el == 'h' || el == 'v' || el == 'w' || el == 'y' {
            score += 4
        } else {
            score += 1
        }
    }
    return score
}
