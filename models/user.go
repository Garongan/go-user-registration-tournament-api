package models

type User struct {
	ID          string `gorm:"type:varchar(255);primaryKey" json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	AccountID   string `json:"-"`
}
