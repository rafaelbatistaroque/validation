package validation

import (
	"fmt"
	"regexp"
)

type EmailRule struct{}

func (r *EmailRule) Validate(value string, fieldName string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	if !re.MatchString(value) {
		return fmt.Errorf("'%s' is not a valid email address", fieldName)
	}

	return nil
}
