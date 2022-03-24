package main

import "errors"

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []user{
	{"user1", "pass1"},
	{"user2", "pass2"},
	{"user3", "pass3"},
}

func registerNewUser(username string, password string) (*user, error) {
	return nil, errors.New("implement this function!")
}

func isUsernameAvailable(username string) bool {
	return false
}
