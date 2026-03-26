package domain

import "errors"

var (
	ErrFetchingUrl = errors.New("error fething url: ")
	ErrRedingBody  = errors.New("error reding body from url: ")
	ErrUnmarshJson = errors.New("error unmarshall json to struct: ")
	ErrIncorrectId = errors.New("error id equal to 0")
)
