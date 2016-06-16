package sanitize

import "regexp"

// Map of regexes
var regexes map[string]*regexp.Regexp

func init() {

	if regexes == nil {
		regexes = make(map[string]*regexp.Regexp)
	}

	regexes["email"] = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	regexes["password"] = regexp.MustCompile(`^([a-zA-Z0-9@._%+-]){8,}$`)
}

// ParseEmail Function validates given email adress with regex
func ParseEmail(email string) (valid bool) {

	valid = false

	if regexes["email"].MatchString(email) {
		valid = true
	}

	return valid
}

// ParsePassword Function validates given password with regex
func ParsePassword(password string) (valid bool) {

	valid = false

	if regexes["password"].MatchString(password) {
		valid = true
	}

	return valid
}
