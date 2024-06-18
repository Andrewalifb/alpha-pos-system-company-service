package utils

import "regexp"

func IsValidEmail(email string) bool {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(?:\.[a-zA-Z]{2,})?$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Match the email address against the regex pattern
	return regex.MatchString(email)
}
