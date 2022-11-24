package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}
