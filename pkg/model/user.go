package model

import (
	"time"

	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/shared/constant"
	"gorm.io/gorm"
)

type User struct {
	ID        string            `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	Role      constant.UserRole `json:"role"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = util.GenerateUUID()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
