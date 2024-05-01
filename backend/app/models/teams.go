package models

import "time"

type Team struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Users       []*User   `gorm:"many2many:user_teams;"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type AccessTokenTeamPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
