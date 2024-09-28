package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AuthByEmail(w http.ResponseWriter, r *http.Request, cfg config.Config) {
	email := chi.URLParam(r, "email")
	code := chi.URLParam(r, "code")

	if !security.CheckInEmailList(email, code) {
		Log(r).Infof("Address %s not in list or code is not valid %s", email, code)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	user, err := UsersQ(r).FilterByEmail(email).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error getting user by email")
		ape.RenderErr(w, problems.InternalError())
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
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	ape.Render(w, SuccessUserAuth(user))
}
