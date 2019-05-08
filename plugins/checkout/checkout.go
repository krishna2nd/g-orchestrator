package main

import (
	"fmt"
	"net/http"
)

type service struct {
	id       string
	sequence int64
}

func (s service) Id() string {
	return "checkout"
}
func (s service) Sequence() int64 {
	return int64(64)
}
func (s service) Format() string {
	return "Format checkout"
}
func (s service) Request() string {
	return "Request checkout"
}
func (s service) Response() string {
	return "Response checkout"
}
func (s service) Url() string {
	return "Url checkout"
}

func (s service) Handle() string {
	return "/checkout"
}

func (s service) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response checkout data")
}

var Service service
