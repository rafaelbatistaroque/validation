package validation

type EqualRule struct {
	Equal string
}

func (r *EqualRule) Validate(value string, fieldName string) error {
	if r.Equal != value {
		return Err_PARAMETER_INVALID
	}

	return nil
}
