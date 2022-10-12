package validate

import "fmt"

func ValidateCreds(username, password string) error {

	if err := checkUsername(username); err != nil {
		return err
	}
	if err := checkPassword(password); err != nil {
		return err
	}

	return nil
}

func checkUsername(name string) error {

	for _, letter := range name {
		if !isNumeric(letter) && !isAlpha(letter) {
			return fmt.Errorf("username must contain only lowercase letters and digits")
		}
	}
	return nil
}

func checkPassword(pass string) error {

	var countDigit, countLower, countUpper int
	if len(pass) < 6 {
		return fmt.Errorf("password must be at least 6 characters in length")
	}

	for _, letter := range pass {
		if isNumeric(letter) {
			countDigit++
			continue
		}
		if isAlpha(letter) {
			countLower++
			continue
		}
		if letter >= 65 && letter <= 90 {
			countUpper++
			continue
		}
		if !isSpecialChar(letter) {
			return fmt.Errorf("password must contain at least 1 digit, 1 uppercase and 1 lowercase letter and special characters")
		}
	}
	if countDigit < 1 || countLower < 1 || countUpper < 1 {
		return fmt.Errorf("password must contain at least 1 digit, 1 uppercase and 1 lowercase letter")
	}
	return nil
}

func isNumeric(letter rune) bool {
	return letter >= 48 && letter <= 57
}

func isAlpha(letter rune) bool {
	return letter >= 97 && letter <= 122
}

func isSpecialChar(letter rune) bool {

	special := "~!@#$%^&*_-+=`|\\(){}[]:;\"'<>,.?/"
	for _, char := range special {
		if char == letter {
			return true
		}
	}
	return false
}
