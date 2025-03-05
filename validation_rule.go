package validation

import (
	"errors"
	"strings"
)

type ValidationRule interface {
	Validate(value string, fieldName string) error
}

type Rule struct {
	errors []string
}

func (i *Rule) ApplyRules(value, fieldName string, rules ...ValidationRule) *Rule {
	for _, rule := range rules {
		if err := rule.Validate(value, fieldName); err != nil {
			i.errors = append(i.errors, err.Error())
		}
	}

	return i
}

func (i *Rule) GetResult() Validatable {
	return i
}

func (i *Rule) IsInvalid() bool {
	return len(i.errors) > 0
}

func (i *Rule) GetErrors() error {
	if len(i.errors) == 0 {
		return nil
	}

	return errors.New(strings.Join(i.errors, "\n"))
}

func (i *Rule) AddError(err error) {
	i.errors = append(i.errors, err.Error())
}
