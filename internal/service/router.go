package service

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service/handlers"
	"github.com/kish1n/GiAuth/internal/service/middleware"
	"gitlab.com/distributed_lab/ape"
	"net/http"
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
		r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
			handlers.Authentication(w, r, cfg)
		})
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTMiddleware(cfg))
			r.Get("/test", handlers.ProtectedHandler)
			r.Get("/logout", handlers.Logout)
			r.Get("/activate_email", handlers.GetActivateEmail)
			r.Put("/activate_email/{code}", handlers.CheckActivatingEmail)
		})
	})

	cfg.Log().Info("Service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}
