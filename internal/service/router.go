package service

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service/handlers"
	"github.com/kish1n/GiAuth/internal/service/middleware"
	"gitlab.com/distributed_lab/ape"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			handlers.CtxLog(cfg.Log()),
		),
		handlers.DBCloneMiddleware(cfg.DB()),
	)

	r.Route("/integrations/GiAuth", func(r chi.Router) {
		r.Post("/reg", handlers.Registration)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				handlers.Authentication(w, r, cfg)
			})
			r.Post("/{email}/{code}", func(w http.ResponseWriter, r *http.Request) {
				handlers.AuthByEmail(w, r, cfg)
			})
		})
		r.Route("/public", func(r chi.Router) {
			r.Use(middleware.JWTMiddleware(cfg))
			r.Get("/test", handlers.ProtectedHandler)
			r.Get("/logout", handlers.Logout)
			r.Route("/email", func(r chi.Router) {
				r.Post("/", func(w http.ResponseWriter, r *http.Request) {
					handlers.GetActivateEmail(w, r, cfg)
				})
				r.Patch("/{code}", handlers.CheckActivatingEmail)
			})
			r.Route("/google_auth", func(r chi.Router) {
				r.Get("/get", handlers.GenerateTOTP)
				r.Get("/validate", handlers.ValidateTOTP)
			})
		})
	})

	cfg.Log().Info("Service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}
