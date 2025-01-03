package validation

import (
	"fmt"
	"strings"
)

type RequiredRule struct{}

func (r *RequiredRule) Validate(value string, fieldName string) error {
	var invalid bool

	switch v := any(value).(type) {
	case string:
		invalid = strings.TrimSpace(v) == ""
	case []string:
		invalid = len(v) == 0
	}

	if invalid {
		return fmt.Errorf(Err_PARAMETER_NOT_EMPTY.Error(), fieldName)
	}

	return nil
}
