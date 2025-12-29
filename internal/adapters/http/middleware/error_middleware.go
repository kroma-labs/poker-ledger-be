package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/logger"
	"github.com/rotisserie/eris"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if err := ctx.Errors.Last(); err != nil {
			logger.Logger.Error(eris.ToString(err, true))
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}
