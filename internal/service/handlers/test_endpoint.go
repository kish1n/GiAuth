package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	Log(r).Infof("%s", r.Context().Value("user"))
	if !ok {
		Log(r).Infof(user)
		ape.RenderErr(w, problems.Unauthorized())
		return
	}
	ape.Render(w, map[string]string{"user": user})
}
