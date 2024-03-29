package service

import (
	"github.com/dl-nft-books/doorman/internal/config"
	"github.com/dl-nft-books/doorman/internal/service/handlers"
	"github.com/dl-nft-books/doorman/internal/service/helpers"
	gosdk "github.com/dl-nft-books/go-sdk"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
			//TODO change when admin's contracts added
			helpers.CtxNodeAdmins(gosdk.NewNodeAdminsMock(cfg.AdminsConfig().Admins...)),
		),
	)

	r.Route("/integrations/doorman", func(r chi.Router) {
		r.Get("/validate-token", handlers.ValidateJWT)
		r.Get("/refresh-token", handlers.RefreshJwt)
		r.Get("/token-pair", handlers.GenerateJwtPair)
		r.Get("/check-permission/{owner}", handlers.CheckResourcePermission)
		r.Get("/check-purpose", handlers.CheckPurpose)
	})

	return r
}
