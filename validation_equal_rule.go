package validation

import "fmt"

type EqualRule struct {
	Equal string
}

func (r *EqualRule) Validate(value string, fieldName string) error {
	if r.Equal != value {
		return fmt.Errorf(Err_PARAMETER_INVALID.Error(), fieldName)
	}

	return nil
}
