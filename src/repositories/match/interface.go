package matchrepository

import (
	"database/sql"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type sMatchRepository struct {
	DB *sql.DB
}

type MatchRepository interface {
	Create(*entities.CreateMatch) error
	FindMany(*int) ([]*entities.Match, error)
	FindOne(*entities.SearchMatch) (*entities.FindOneMatch, error)
	Approve(*int) error
	// Reject(*int) error
	Delete(*int) error
}

func New(db *sql.DB) MatchRepository {
	return &sMatchRepository{DB: db}
}
