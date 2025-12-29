package provider

import "github.com/kroma-labs/poker-ledger-be/internal/adapters/http/handler"

type HTTPHandlers struct {
	Room *handler.RoomHandler
}

func provideHTTPHandlers(usecases *Usecases) *HTTPHandlers {
	return &HTTPHandlers{
		handler.NewRoomHandler(usecases.Room),
	}
}
