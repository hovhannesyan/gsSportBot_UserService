package models

type User struct {
	ChatID   int64   `json:"chat_id" gorm:"primaryKey"`
	Username string  `json:"username"`
	Name     *string `json:"name"`
	Phone    *string `json:"phone"`
	IsAdmin  bool    `json:"is_admin"`
}
