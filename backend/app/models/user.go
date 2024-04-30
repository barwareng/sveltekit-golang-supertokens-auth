package models

type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"index;not null;unique"`
	// Other fields
}
