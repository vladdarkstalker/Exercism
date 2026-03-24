package bob

import (
    "unicode"
    "strings"
)

func Hey(remark string) string {
    trimmedRemark := strings.TrimSpace(remark)
    if trimmedRemark == "" {
        return "Fine. Be that way!"
    }
    
    hasLetters := false
    allUpper := true
    for _, r := range trimmedRemark {
        if unicode.IsLetter(r) {
            hasLetters = true
            if !unicode.IsUpper(r) {
                allUpper = false
            }
        }
    }
    
    isQuestion := strings.HasSuffix(trimmedRemark, "?")
    switch {
        case hasLetters && allUpper && isQuestion:
        	return "Calm down, I know what I'm doing!"
        case isQuestion:
        	return "Sure."
        case allUpper && hasLetters:
        	return "Whoa, chill out!"
        default:
        	return "Whatever."
    }
}
