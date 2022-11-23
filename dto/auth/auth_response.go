package authdto

type LoginResponse struct {
	ID    int    `json:"id" `
	Name  string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Email string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Token string `gorm:"type: varchar(255)" json:"token" validate:"required"`
}

type CheckAuthResponse struct {
	Id    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(50)"  json:"role"`
}
