package main

import (
	"net/url"
	"testing"
)

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	if isUsernameAvailable("user1") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}
