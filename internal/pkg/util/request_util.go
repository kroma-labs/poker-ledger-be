package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/itsLeonB/ginkgo"
)

func BindJSON[t any](ctx *gin.Context) (t, error) {
	return ginkgo.BindRequest[t](ctx, binding.JSON)
}
