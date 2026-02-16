package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
    been := make(map[rune]bool)
    for _, symbol := range word {
        if symbol == '-' || symbol == ' ' {
            continue
        }
        if been[symbol] {
            return false
        }
        been[symbol] = true
    }
    return true
}
