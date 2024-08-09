package services

import "unicode"

func CalculateStrongPasswordSteps(password string) int {
	length := len(password)
	if length < 6 {
		return 6 - length
	} else if length > 20 {
		return length - 20
	}

	hasLower, hasUpper, hasDigit := false, false, false
	repeats := 0
	for i := 0; i < length; i++ {
		if unicode.IsLower(rune(password[i])) {
			hasLower = true
		} else if unicode.IsUpper(rune(password[i])) {
			hasUpper = true
		} else if unicode.IsDigit(rune(password[i])) {
			hasDigit = true
		}

		if i >= 2 && password[i] == password[i-1] && password[i-1] == password[i-2] {
			repeats++
		}
	}

	missingTypes := 0
	if !hasLower {
		missingTypes++
	}
	if !hasUpper {
		missingTypes++
	}
	if !hasDigit {
		missingTypes++
	}

	return max(missingTypes, repeats)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
