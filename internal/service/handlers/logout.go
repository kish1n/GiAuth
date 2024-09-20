package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	// Отправляем успешный ответ
	ape.Render(w, map[string]string{
		"message": "Successfully logged out",
	})
}
