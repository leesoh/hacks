package main

import "fmt"

type ErrRateLimiting struct{}

func (e ErrRateLimiting) Error() string {
	return "rate limiting encountered. terminating."
}

type ErrUsernameDoesNotExist struct {
	u string
}

func (e ErrUsernameDoesNotExist) Error() string {
	return fmt.Sprintf("username does not exist: %v", e.u)
}
