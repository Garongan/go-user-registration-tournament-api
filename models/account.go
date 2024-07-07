package models

type Account struct {
	ID       string `gorm:"type:varchar(255);primaryKey" json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	UserID   string `json:"-"`
}
