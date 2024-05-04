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
	// FindMany()
	// Approve()
	// Reject()
	// Delete()
}

func New(db *sql.DB) MatchRepository {
	return &sMatchRepository{DB: db}
}
