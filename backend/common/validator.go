package common

import (
	"net/mail"
	"net/url"
	"regexp"
)

func ValidateEmptyString(s string) bool {
	return s == ""
}

func ValidatePhone(s string) bool {
	if len(s) != 10 && len(s) != 11 {
		return false
	}

	pattern := `^[0-9]+$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(s)
}

func ValidateUrl(s string) bool {
	u, err := url.ParseRequestURI(s)

	if err != nil || u.Scheme == "" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}

func ValidateEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func ValidateNotNilId(id *string) bool {
	if id == nil || len(*id) == 0 || len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func ValidateId(id *string) bool {
	if id == nil || len(*id) == 0 {
		return true
	}
	if len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func ValidateNegativeNumberInt(number int) bool {
	return number < 0
}
func ValidateNegativeNumberFloat(number float32) bool {
	return number < 0
}

func ValidateNotPositiveNumberInt(number int) bool {
	return number <= 0
}

func ValidatePassword(pass *string) bool {
	return pass != nil && len(*pass) >= 6
}

func ValidatePositiveNumberInt(number int) bool {
	return number > 0
}
