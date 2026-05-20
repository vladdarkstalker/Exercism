package phonenumber

import (
    "errors"
    "fmt"
)

func Number(phoneNumber string) (string, error) {
	var result string
    for _, el := range phoneNumber {
        if el >= '0' && el <= '9' {
            result += string(el)
        }
    }
    if len(result) == 11 && result[0] == '1' {
        result = result[1:]
    }
    if len(result) != 10 {
        return "", errors.New("invalid phone number")
    }
    if result[0] < '2' || result[0] > '9' {
        return "", errors.New("invalid phone number")
    }
    if result[3] < '2' || result[3] > '9' {
        return "", errors.New("invalid phone number")
    }
    return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
    cleaned, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return cleaned[:3], nil
}

func Format(phoneNumber string) (string, error) {
    cleaned, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", cleaned[:3], cleaned[3:6], cleaned[6:]), nil
}
