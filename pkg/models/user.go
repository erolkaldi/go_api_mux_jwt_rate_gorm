package models

type User struct {
	ID             int    `gorm:"primary_key" json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	EmailConfirmed bool   `json:"email_confirmed"`
}
