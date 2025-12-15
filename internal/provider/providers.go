package provider

import (
	"database/sql"

	"github.com/kroma-labs/poker-ledger-be/internal/pkg/config"
)

type Providers struct {
	DB *sql.DB
	*Repositories
	*Usecases
	*HttpHandlers
}

func ProvideAll(cfg *config.Config) (*Providers, error) {
	db, err := provideDB(cfg.DBString)
	if err != nil {
		return nil, err
	}

	repos := provideRepositories(db)
	usecases := provideUsecases(repos)

	return &Providers{
		db,
		repos,
		usecases,
		provideHttpHandlers(usecases),
	}, nil
}

func (p *Providers) Shutdown() error {
	return p.DB.Close()
}
