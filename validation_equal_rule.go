package validation

type EqualRule struct {
	Equal any
}

func (r *EqualRule) Validate(value any, fieldName string) error {
	if r.Equal == value {
		return Err_PARAMETER_INVALID
	}

	return nil
}
