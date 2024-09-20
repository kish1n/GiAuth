package handlers

import (
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/data"
	"github.com/kish1n/GiAuth/internal/service/requests"
	"github.com/kish1n/GiAuth/internal/service/security"
	"github.com/kish1n/GiAuth/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"time"
)

func Authentication(w http.ResponseWriter, r *http.Request, cfg config.Config) {
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

	token, err := security.GenerateJWT(user, cfg)
	if err != nil {
		Log(r).WithError(err).Error("Error generating JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * cfg.JWT().ExpirationTime),
		HttpOnly: true,
		Secure:   true, // true for HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	ape.Render(w, SuccessUserAuth(user))
}

func SuccessUserAuth(user *data.User) resources.SuccessAuthResponse {
	return resources.SuccessAuthResponse{
		Data: resources.SuccessAuth{
			Key: resources.Key{
				ID:   user.Username,
				Type: resources.SUCCESS_AUTH,
			},
			Attributes: resources.SuccessAuthAttributes{
				Email:      user.Email,
				FirstName:  user.FirstName,
				LastName:   user.LastName,
				MiddleName: user.MiddleName,
				Username:   user.Username,
			},
		},
	}
}
