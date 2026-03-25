package luhn

import "unicode"

func Valid(id string) bool {
    len_id := len(id)
    if len_id <= 0 {
        return false;
    }
    slice_runes := []rune(id)
    digits := make([]int, 0)
    for _, i := range slice_runes {
        if unicode.IsSpace(i) {
            continue
        }
        if !unicode.IsDigit(i) {
            return false
        }
        digits = append(digits, int(i-'0'))
    }
    if len(digits) == 1 {
        return false
    }
    sum := 0
    var is_doubling bool
    if len(digits) % 2 == 0 {
        is_doubling = true
    } else {
        is_doubling = false
    }
    for _, digit := range digits {
        if is_doubling {
            digit *= 2
            if digit > 9 {
              digit -= 9
          }
        }
        sum += digit
        is_doubling = !is_doubling
    }
    if sum % 10 == 0 {
        return true
    } else {
        return false
    }
}
