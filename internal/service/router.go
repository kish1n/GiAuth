package service

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/kish1n/GiAuth/internal/config"
	"github.com/kish1n/GiAuth/internal/service/handlers"
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
	)
	r.Route("/integrations/GiAuth", func(r chi.Router) {
		r.Post("/reg", handlers.Registration)
	})

	cfg.Log().Info("Service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}
