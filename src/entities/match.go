package entities

import (
	"time"

	"github.com/lib/pq"
)

type SearchMatch struct {
	MatchCatId int `json:"matchCatId" validate:"required"`
}

type CreateMatch struct {
	MatchCatId int    `json:"matchCatId" validate:"required"`
	UserCatId  int    `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}

type IssuedBy struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type MatchCatDetail struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	Description string         `json:"description"`
	AgeInMonth  int            `json:"ageInMonth"`
	ImageUrls   pq.StringArray `json:"imageUrls"`
	HasMatched  bool           `json:"hasMatched"`
	CreatedAt   time.Time      `json:"createdAt"`
}

type UserCatDetail struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Race        string         `json:"race"`
	Sex         string         `json:"sex"`
	Description string         `json:"description"`
	AgeInMonth  int            `json:"ageInMonth"`
	ImageUrls   pq.StringArray `json:"imageUrls"`
	HasMatched  bool           `json:"hasMatched"`
	CreatedAt   time.Time      `json:"createdAt"`
}

type Match struct {
	ID             string         `json:"id"`
	IssuedBy       IssuedBy       `json:"issuedBy"`
	MatchCatDetail MatchCatDetail `json:"matchCatDetail"`
	UserCatDetail  UserCatDetail  `json:"userCatDetail"`
	Message        string         `json:"message"`
	CreatedAt      time.Time      `json:"createdAt"`
}

type FindOneMatch struct {
	ID         string    `json:"id"`
	UserCatId  string    `json:"user_cat_id"`
	MatchCatId string    `json:"match_cat_id"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"-"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
