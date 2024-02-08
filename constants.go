package main

import "errors"

type Method int
type Auth int

var (
	InvalidMethodErr = errors.New("Method is invalid")
	methods          = []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
)

const (
	none Auth = iota
	basic
	bearer
	key
)

func authDisplay(auth Auth) (authName string) {
	authName = ""
	switch auth {
	case none:
		authName = "None"
	case basic:
		authName = "Basic"
	case bearer:
		authName = "Bearer"
	case key:
		authName = "API Key"
	}

	return
}
