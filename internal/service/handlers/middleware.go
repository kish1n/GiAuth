package handlers

import (
	"context"
	"net/http"

	"github.com/kish1n/GiAuth/internal/data/pg"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func DBCloneMiddleware(db *pgdb.DB) func(http.Handler) http.Handler {
	type ctxExtender func(context.Context) context.Context

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clone := db.Clone()
			ctx := r.Context()

			extenders := []ctxExtender{
				CtxUsersQ(pg.NewUsers(clone)),
			}

			for _, extender := range extenders {
				ctx = extender(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
