package domain

import "testing"

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("We were not expecting a user with ID 0")
	}

	if err == nil {
		t.Error("We were expecting an error")
	}
}