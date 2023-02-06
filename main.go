package main

import (
	"errors"
	"go-http/network"
	"net/http"
	"time"
)

const endpoint string = "http://localhost:8080/api/v1"

type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname,omitempty"`
	Password  string `json:"password,omitempty"`
}

func CreateUser() (*User, error) {
	var (
		err  error
		user User
	)

	statusCode, err := network.NewRequest().
		Method(http.MethodPost).
		URL(endpoint + "/users").
		Headers(&network.Headers{
			"Content-Type":    "application/json",
			"x-access-token":  "<ACCESS_TOKEN>",
			"x-refresh-token": "<REFRESH_TOKEN>",
		}).
		Body(&network.Body{
			"username":  "",
			"firstname": "",
			"lastname":  "",
			"password":  "",
		}).
		Submit().
		Retry(5 * time.Second). // retry in 5 seconds after request failed
		BindJSON(&user)
		// .BindText(&user)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case http.StatusCreated:
		err = nil
	case http.StatusBadRequest:
		err = errors.New("Username already exists")
	case http.StatusUnauthorized:
		err = errors.New("token not found")
	default:
		return errors.New("Something went wrong")
	}

	return &user, err
}

func FetchUsers() (*[]User, error) {
	var users []User

	err := network.NewRequest().
		Method(http.MethodGet).
		URL(endpoint+"/users").
		Header("Content-Type", "application/json").
		Header("x-access-token", "<ACCESS_TOKEN>").
		Header("x-refresh-token", "<REFRESH_TOKEN>").
		Send().
		Bind(&users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func FetchSingleUser(userID string) (*User, error) {
	var user User

	err := network.NewRequest().
		Method(http.MethodGet).
		URL(endpoint+"/users/"+userID).
		Header("Content-Type", "application/json").
		Header("x-access-token", "<ACCESS_TOKEN>").
		Header("x-refresh-token", "<REFRESH_TOKEN>").
		Send().
		Bind(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
