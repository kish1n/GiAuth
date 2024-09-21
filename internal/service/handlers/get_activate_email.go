package handlers

import (
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetActivateEmail(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(user)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	secretKey, err := security.GenerateConfirmationCode()
	if err != nil {
		Log(r).WithError(err).Error("Error create secret key")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = security.SendConfirmationEmail(user, secretKey)
	if err != nil {
		Log(r).WithError(err).Error("Error Send code to email")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusFound)
}
