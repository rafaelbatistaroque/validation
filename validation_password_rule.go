package validation

import (
	"fmt"
	"unicode"
)

type PasswordRule struct {
	MinLength    int
	RequireLower bool
	RequireUpper bool
	RequireDigit bool
}

func (r *PasswordRule) Validate(value string, fieldName string) error {
	if len(value) < r.MinLength {
		return fmt.Errorf(Err_PARAMETER_LENGHT_INVALID.Error(), fieldName, r.MinLength)
	}

	var hasLower, hasUpper, hasDigit bool

	for _, char := range value {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if r.RequireLower && !hasLower {
		return fmt.Errorf(Err_PARAMETER_SHOULD_HAVE_LOWER_CHARACTER.Error(), fieldName)
	}

	if r.RequireUpper && !hasUpper {
		return fmt.Errorf(Err_PARAMETER_SHOULD_HAVE_UPPER_CHARACTER.Error(), fieldName)
	}

	if r.RequireDigit && !hasDigit {
		return fmt.Errorf(Err_PARAMETER_SHOULD_HAVE_DIGIT_CHARACTER.Error(), fieldName)
	}

	return nil
}
