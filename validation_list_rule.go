package validation

import (
	"fmt"
)

type ListRule struct {
	Contains    []string
	NotContains []string
}

func (r *ListRule) Validate(value string, fieldName string) error {
	if len(r.Contains) > 0 && !isInList(value, r.Contains) {
		return fmt.Errorf(Err_PARAMETER_SHOULD_BE_WITHIN_LIST.Error(), fieldName, r.Contains)
	}

	if len(r.NotContains) > 0 && isInList(value, r.NotContains) {
		return fmt.Errorf(Err_PARAMETER_SHOULD_BE_OF_LIST.Error(), fieldName)
	}

	return nil
}

func isInList(value string, list []string) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}
