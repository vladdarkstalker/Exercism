package acronym

import "unicode"

func Abbreviate(s string) string {
    result := ""

    nextword := true
    for _, element := range s {
        if element == ' ' || element == '-' || element == '_' {
            nextword = true
        } else if nextword {
            result += string(unicode.ToUpper(element))
            nextword = false
        }
    }
    
	return result
}
