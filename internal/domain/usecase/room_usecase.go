package usecase

import (
	"context"
	"time"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/dto"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/mapper"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/repository"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/stringutil"
)

type RoomUsecase interface {
	Create(ctx context.Context, request dto.NewRoomRequest) (dto.RoomResponse, error)
	GetAll(ctx context.Context) ([]dto.RoomResponse, error)
}

type roomUsecase struct {
	transactor repository.Transactor
	roomRepo   repository.RoomRepository
	playerRepo repository.PlayerRepository
}

func NewRoomUsecase(
	transactor repository.Transactor,
	roomRepo repository.RoomRepository,
	playerRepo repository.PlayerRepository,
) RoomUsecase {
	return &roomUsecase{
		transactor,
		roomRepo,
		playerRepo,
	}
}

func (ru *roomUsecase) Create(
	ctx context.Context,
	request dto.NewRoomRequest,
) (dto.RoomResponse, error) {
	var response dto.RoomResponse
	err := ru.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		newHost := entity.Player{
			Name:      request.HostName,
			CreatedAt: time.Now(),
		}

		host, err := ru.playerRepo.Insert(ctx, newHost)
		if err != nil {
			return err
		}

		code, err := stringutil.GenerateRandomString(6)
		if err != nil {
			return err
		}

		newRoom := entity.Room{
			Code:         code,
			HostPlayerID: host.ID,
			Status:       entity.RoomStatusWaiting,
			CreatedAt:    time.Now(),
		}

		createdRoom, err := ru.roomRepo.Insert(ctx, newRoom)
		if err != nil {
			return err
		}

		response = dto.RoomResponse{
			Code:   createdRoom.Code,
			Status: string(createdRoom.Status),
			HostPlayer: dto.PlayerResponse{
				ID:   host.ID,
				Name: host.Name,
			},
		}

		return nil
	})
	return response, err
}

func (ru *roomUsecase) GetAll(ctx context.Context) ([]dto.RoomResponse, error) {
	rooms, err := ru.roomRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(rooms, mapper.RoomToResponse), nil
}
