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
	Update(*int, *entities.ParamsUpdateCat) (*entities.CreateCat, error)
	Delete(*int) error
}

func New(db *sql.DB) CatRepository {
	return &sCatRepository{DB: db}
}
