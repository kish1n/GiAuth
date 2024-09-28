package requests

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func newDecodeError(what string, err error) error {
	return validation.Errors{
		what: fmt.Errorf("decode requests %s: %w", what, err),
	}
}
