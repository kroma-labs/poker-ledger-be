package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/dto"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/usecase"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/util"
)

type RoomHandler struct {
	roomUc usecase.RoomUsecase
}

func NewRoomHandler(roomUc usecase.RoomUsecase) *RoomHandler {
	return &RoomHandler{roomUc}
}

func (rh *RoomHandler) HandleCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := util.BindJSON[dto.NewRoomRequest](ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		response, err := rh.roomUc.Create(ctx, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
