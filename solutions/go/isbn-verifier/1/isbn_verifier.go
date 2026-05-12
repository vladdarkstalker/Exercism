package isbnverifier

func IsValidISBN(isbn string) bool {
    if len(isbn) < 10 || isbn == "" {
        return false
    }
    var sum int
    for i, j := 0, 0; i < len(isbn); i++ {
        if isbn[i] == '-' {
            continue
        } else {
            if isbn[i] == 'X' && j == 9 {
                sum += 10 * (10 - j)
            } else {
                sum += int(isbn[i] - '0') * (10 - j)
            }
            j++
            if j >= 11 {
                return false
            }
        }
    }
	return sum % 11 == 0
}
