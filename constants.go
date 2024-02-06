package main

import "errors"

type Method int

const (
	get Method = iota
	head
	post
	put
	patch
	delete
	options
)

var (
	InvalidMethodErr = errors.New("Method is invalid")
)

func methodDisplay(method Method) (methodName string, err error) {
	err = nil
	switch method {
	case get:
		methodName = "GET"
	case head:
		methodName = "HEAD"
	case post:
		methodName = "POST"
	case put:
		methodName = "PUT"
	case patch:
		methodName = "PATCH"
	case delete:
		methodName = "DELETE"
	case options:
		methodName = "OPTIONS"
	}

	return
}
