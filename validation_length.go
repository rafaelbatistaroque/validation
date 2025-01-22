package validation

import "fmt"

type LengthRule struct {
	Exactly int
	Max     int
	Min     int
}

func (r *LengthRule) Validate(value string, fieldName string) error {
	length := len(value)

	if r.Exactly > 0 && length != r.Exactly {
		return fmt.Errorf(Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER.Error(), fieldName, r.Exactly)
	}

	if r.Min > 0 && length < r.Min {
		return fmt.Errorf(Err_PARAMETER_LENGHT_MIN_INVALID.Error(), fieldName, r.Min)
	}

	if r.Max > 0 && length > r.Max {
		return fmt.Errorf(Err_PARAMETER_LENGHT_MAX_INVALID.Error(), fieldName, r.Max)
	}

	return nil
}
