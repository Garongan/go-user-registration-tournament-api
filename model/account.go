package model

type Account struct {
	ID       string `gorm:"type:varchar(255);primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	User     User   `json:"user"`
}
