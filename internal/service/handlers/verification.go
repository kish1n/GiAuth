package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	if !ok {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	ape.Render(w, map[string]string{"user": user})
}
