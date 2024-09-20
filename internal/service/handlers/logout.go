package handlers

import (
	"github.com/kish1n/GiAuth/resources"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(user)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	currentUser, err := UsersQ(r).FilterByUsername(user).Get()
	if err != nil {
		Log(r).WithError(err).Error("Error filter by username")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if currentUser == nil {
		Log(r).Errorf("User with this username: %s, already dosent exit", user)
		ape.RenderErr(w, problems.Conflict())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	ape.Render(w, resources.SuccesLogout{
		Key: resources.Key{
			ID: currentUser.Username,
		},
		Attributes: resources.SuccesLogoutAttributes{
			Email:      currentUser.Email,
			FirstName:  currentUser.FirstName,
			LastName:   currentUser.LastName,
			MiddleName: currentUser.MiddleName,
			Username:   currentUser.Username,
		},
	})
}
