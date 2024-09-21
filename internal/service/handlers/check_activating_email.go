package handlers

import (
	"github.com/go-chi/chi"
	"github.com/kish1n/GiAuth/internal/service/security"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strings"
)

func CheckActivatingEmail(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	if !ok {
		Log(r).Infof(user)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	code := strings.ToLower(chi.URLParam(r, "code"))

	if !security.CheckInEmailList(user, code) {
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	w.WriteHeader(http.StatusFound)
}
