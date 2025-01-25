package validation

import (
	"fmt"
	"regexp"
)

type SocialURLRule struct{}

func (r *SocialURLRule) Validate(value string, fieldName string) error {

	regex := `^https://[a-zA-Z0-9]+([a-zA-Z0-9.-]*[a-zA-Z0-9])?(/[a-zA-Z0-9._~:/?#@!$&'()*+,;=%-]*)?$`

	re := regexp.MustCompile(regex)

	if !re.MatchString(value) {
		return fmt.Errorf(Err_PARAMETER_SHOULD_BE_SOCIAL_URL_INVALID.Error(), fieldName)
	}

	return nil
}
