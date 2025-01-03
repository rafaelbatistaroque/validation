package validation

import "fmt"

type LengthRule struct {
	Length int
}

func (r *LengthRule) Validate(value string, fieldName string) error {
	if len(value) != r.Length {
		return fmt.Errorf(Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), fieldName, r.Length)
	}
	return nil
}
