package validate

import (
	errors2 "github.com/DenisKhanov/PrivateKeeper/pkg/errors"
	"regexp"
)

// CheckEmail checks if the email adheres to the standard format
func CheckEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	//compiling regex
	re := regexp.MustCompile(emailRegex)

	// check email at regex
	if !re.MatchString(email) {
		return errors2.ErrEmail
	}
	return nil
}
