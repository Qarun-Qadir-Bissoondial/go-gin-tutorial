package main

import (
	"errors"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	{"user1", "pass1"},
	{"user2", "pass2"},
	{"user3", "pass3"},
}

func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return nil, errors.New("blank username/password provided")
	}

	if !isUsernameAvailable(username) {
		return nil, errors.New("username already taken")
	}

	// TODO: Implement more password validation if the front-end does not implement validation

	newUser := user{username, password}
	userList = append(userList, newUser)
	return &newUser, nil
}

func isUserValid(username, password string) bool {
	return false
}

func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}

	return true
}
