package entities

import (
	"time"

	"github.com/lib/pq"
)

type Cat struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	AgeInMonth  int            `json:"ageInMonth"`
	ImageUrls   pq.StringArray `json:"imageUrls"`
	Description string         `json:"description"`
	HasMatched  bool           `json:"hasMatched"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"-"`
}

type CreateCat struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type CatSearchFilter struct {
	ID         int
	Limit      int
	Offset     int
	Race       string
	Sex        string
	HasMatched bool
	AgeInMonth string
	Owned      bool
	Search     string
}

type ParamsCreateCat struct {
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   pq.StringArray
	UserId      int
}

type ParamsUpdateCat struct {
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   pq.StringArray
}
