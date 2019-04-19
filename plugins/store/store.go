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
	return "stores"
}
func (s service) Sequence() int64 {
	return int64(64)
}
func (s service) Format() string {
	return "Format stores"
}
func (s service) Request() string {
	return "Request stores"
}
func (s service) Response() string {
	return "Response stores"
}
func (s service) Url() string {
	return "Url stores"
}

func (s service) Handle() string {
	return "/stores"
}

func (s service) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response stores")
}

var Service service
