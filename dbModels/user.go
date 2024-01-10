package dbModels

import (
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type DbUser struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CreateTime  time.Time
	Email       string
	UserName    string
	Password    string
	DisplayName string
	IsAdmin     bool
}

func (DbUser) TableName() string {
	return "user"
}

func (u *DbUser) ToGQLUser() *model.User {
	return &model.User{
		ID:          u.ID.String(),
		Email:       u.Email,
		CreateTime:  int(u.CreateTime.Unix()),
		UserName:    u.UserName,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}
