package catrepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sCatRepository struct {
	DB *sql.DB
}

type CatRepository interface {
	Create(*entities.ParamsCreateCat) (*entities.CreateCat, error)
	FindMany(*entities.CatSearchFilter) ([]*entities.Cat, error)
	// FindOne() (*entities.Cat, error)
	// Update() (*entities.Cat, error)
	// Delete() (*entities.Cat, error)
}

func New(db *sql.DB) CatRepository {
	return &sCatRepository{DB: db}
}
