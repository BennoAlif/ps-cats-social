package entities

import "time"

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
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserCatDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Match struct {
	ID             string         `json:"id"`
	IssuedBy       IssuedBy       `json:"issuedBy"`
	MatchCatDetail MatchCatDetail `json:"matchCatDetail"`
	UserCatDetail  UserCatDetail  `json:"userCatDetail"`
	Message        string         `json:"message"`
	CreatedAt      time.Time      `json:"createdAt"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
