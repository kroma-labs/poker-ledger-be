package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kroma-labs/poker-ledger-be/internal/adapters/http/middleware"
	"github.com/kroma-labs/poker-ledger-be/internal/provider"
)

func setupRoutes(r *gin.Engine, handlers *provider.HttpHandlers) {
	r.Use(middleware.Error())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/rooms", handlers.Room.HandleCreate())
		}
	}
}
