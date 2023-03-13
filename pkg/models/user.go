package models

type User struct {
	Id       int64   `json:"id" gorm:"primaryKey"`
	Username string  `json:"username"`
	Name     *string `json:"name"`
	Phone    *string `json:"phone"`
	IsAdmin  bool    `json:"is_admin"`
}
