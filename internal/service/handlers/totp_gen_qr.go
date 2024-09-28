package handlers

import (
	"net/http"

	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GenerateTOTP(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(username)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	usr, err := UsersQ(r).FilterByUsername(username).Get()
	if err != nil {
		Log(r).WithError(err).Infof("Error get secretkey")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	uri := security.GenerateTOTPQRCodeURL(usr.SecretKey, username)
	qrCode, err := security.GenerateQRCode(uri)
	if err != nil {
		Log(r).WithError(err).Error("Error generating QR code")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(qrCode)
}
