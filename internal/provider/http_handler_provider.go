package provider

import "github.com/kroma-labs/poker-ledger-be/internal/adapters/http/handler"

type HttpHandlers struct {
	Room *handler.RoomHandler
}

func provideHttpHandlers(usecases *Usecases) *HttpHandlers {
	return &HttpHandlers{
		handler.NewRoomHandler(usecases.Room),
	}
}
