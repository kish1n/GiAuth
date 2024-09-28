package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kish1n/GiAuth/resources"
)

func NewRegistration(r *http.Request) (req resources.UserFormRegRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	return req, validation.Errors{
		"data/id":         validation.Validate(req.Data.ID, validation.Required),
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In(resources.REGISTRATION)),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}.Filter()
}
