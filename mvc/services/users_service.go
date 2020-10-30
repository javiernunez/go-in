package services

import (
	"github.com/javiernunez/go-in/mvc/domain"
	"github.com/javiernunez/go-in/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
