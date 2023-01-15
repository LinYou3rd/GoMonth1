package user

import "github.com/jinzhu/gorm"

type (
	userModel struct {
		gorm.Model
		Name           string `json:"name"`
		Email          string `json:"email"`
		PasswordDigest string `json:"passwordDigest"`
	}

	userService struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
)

func (userModel) TableName() string {
	return "users"
}
