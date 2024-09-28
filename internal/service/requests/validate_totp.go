package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kish1n/GiAuth/resources"
)

func NewValidateTotp(r *http.Request) (req resources.ValidateTotpResponse, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/type": validation.Validate(req.Data.Type, validation.Required, validation.In(resources.VALIDATE_TOTP)),
		"data/code": validation.Validate(req.Data.Attributes.Code, validation.Required),
	}.Filter()
}
