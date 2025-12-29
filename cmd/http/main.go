package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kroma-labs/poker-ledger-be/internal/adapters/http"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/config"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/logger"
	"github.com/rotisserie/eris"
)

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		logger.Logger.Fatal(eris.ToString(err, true))
	}

	srv, err := http.Setup(cfg)
	if err != nil {
		logger.Logger.Fatal(eris.ToString(err, true))
	}

	srv.ServeGracefully()
}
