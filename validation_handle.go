package validation

type Validatable interface {
	IsInvalid() bool
	GetErrors() error
}
