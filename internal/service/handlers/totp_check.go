package handlers

import (
	"net/http"

	"github.com/kish1n/GiAuth/internal/service/requests"
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ValidateTOTP(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewValidateTotp(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(username)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	secret, err := UsersQ(r).FilterByUsername(username).Get()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	valid := security.ValidateTOTPCode(secret.SecretKey, req.Data.Attributes.Code)
	if !valid {
		http.Error(w, "Invalid code", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Code valid!"))
}
