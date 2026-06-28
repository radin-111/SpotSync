package users

import "gorm.io/gorm"

const (
	UserRoleDriver = "driver"
	UserRoleAdmin  = "admin"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Role     string `json:"role" gorm:"type:varchar(10);default:driver"`
}
