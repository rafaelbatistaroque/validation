package validation

import "errors"

var (
	Err_PARAMETER_SHOULD_BE_GREATHER_THAN_ZERO  = errors.New("parameter %s should be greather than zero")
	Err_PARAMETER_NOT_EMPTY                     = errors.New("parameter %s cannot be null or empty")
	Err_PARAMETER_LENGHT_INVALID                = errors.New("parameter %s should have at least %d characters")
	Err_PARAMETER_LENGHT_MIN_INVALID            = errors.New("parameter %s should have at least %d characters")
	Err_PARAMETER_LENGHT_MAX_INVALID            = errors.New("parameter %s should have at most %d characters")
	Err_PARAMETER_INVALID                       = errors.New("parameter %s is invalid")
	Err_PARAMETER_SHOULD_HAVE_LOWER_CHARACTER   = errors.New("parameter %s should have at least one lower case character")
	Err_PARAMETER_SHOULD_HAVE_UPPER_CHARACTER   = errors.New("parameter %s should have at least one upper case character")
	Err_PARAMETER_SHOULD_HAVE_DIGIT_CHARACTER   = errors.New("parameter %s should have at least one digit character")
	Err_PARAMETER_SHOULD_HAVE_EXACTLY_CHARACTER = errors.New("parameter %s must have exactly %d characters")
	Err_PARAMETER_SHOULD_BE_SOCIAL_URL_INVALID  = errors.New("parameter %s is not a valid social media URL")
	Err_PARAMETER_EMAIL_INVALID                 = errors.New("parameter %s is not a valid email address")
	Err_PARAMETER_SHOULD_BE_WITHIN_LIST         = errors.New("parameter %s should be one of the allowed values: %v")
	Err_PARAMETER_SHOULD_BE_OF_LIST             = errors.New("parameter %s is invalid")
)
