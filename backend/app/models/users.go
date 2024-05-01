package models

type User struct {
	commonFields
	ID          string  `json:"id" gorm:"primaryKey"`
	Email       string  `json:"email" gorm:"not null;unique"`
	PhoneNumber string  `json:"phoneNumber" gorm:"unique;default:null"`
	FirstName   string  `json:"firstName"`
	MiddleName  string  `json:"middleName"`
	LastName    string  `json:"lastName"`
	Teams       []*Team `gorm:"many2many:user_teams;"`
}
