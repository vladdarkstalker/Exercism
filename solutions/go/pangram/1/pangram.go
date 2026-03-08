package pangram

import "strings"

func IsPangram(input string) bool {
	if input == "" {
        return false
    }

    input = strings.ToLower(input)
    pangram := make(map[rune]bool, 26)

    for _, element := range input {
        if element >= 'a' && element <= 'z' {
            pangram[element] = true
        }
    }
    
    return len(pangram) == 26
}
