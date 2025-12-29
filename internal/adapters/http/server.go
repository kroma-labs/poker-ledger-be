package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/ginkgo"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/config"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/logger"
	"github.com/kroma-labs/poker-ledger-be/internal/provider"
)

func Setup(configs *config.Config) (*ginkgo.HttpServer, error) {
	providers, err := provider.ProvideAll(configs)
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	setupRoutes(r, providers.HTTPHandlers)

	srv := &http.Server{
		Addr:              ":" + configs.Port,
		Handler:           r,
		ReadTimeout:       configs.Timeout,
		ReadHeaderTimeout: configs.Timeout,
		WriteTimeout:      configs.Timeout,
		IdleTimeout:       configs.Timeout,
	}

	return ginkgo.NewHttpServer(srv, configs.Timeout, logger.Logger, providers.Shutdown), nil
}
