package domain

import (
	"fmt"
	"github.com/javiernunez/go-in/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User {
		123: &User{Id:123, FirstName: "Javi", LastName: "Nuñez", Email: "javi@javiernunez.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %v not found", userId),
		Status: http.StatusNotFound,
		Code: "not_found",
	}
}
