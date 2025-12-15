package mapper

import (
	"github.com/kroma-labs/poker-ledger-be/internal/domain/dto"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
)

func RoomToResponse(room entity.Room) dto.RoomResponse {
	return dto.RoomResponse{
		Code:   room.Code,
		Status: string(room.Status),
	}
}
