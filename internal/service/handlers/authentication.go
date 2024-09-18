package handlers

import (
	"github.com/kish1n/GiAuth/internal/service/requests"
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAuthentication(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := UsersQ(r).FilterByUsername(req.Data.ID).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error filter by username")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !security.CheckPasswordHash(req.Data.Attributes.Password, user.PasswordHash) {
		Log(r).WithError(err).Error("Password or login incorrect")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	ape.Render(w, SuccessUserAuth(*user))
}
