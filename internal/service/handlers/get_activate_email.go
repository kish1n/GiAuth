package handlers

import (
	"net/http"

	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetActivateEmail(w http.ResponseWriter, r *http.Request, cfg config.Config) {
	username, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(username)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	code, err := security.GenerateConfirmationCode()
	if err != nil {
		Log(r).WithError(err).Error("Error create secret key")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	user, err := UsersQ(r).FilterByUsername(username).Get()
	if err != nil {
		Log(r).WithError(err).Error("Cannot get user by cookie")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = security.SendConfirmationEmail(user.Email, code, cfg)
	if err != nil {
		Log(r).WithError(err).Error("Error Send code to email")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
