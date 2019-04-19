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
	return "orders"
}
func (s service) Sequence() int64 {
	return int64(64)
}
func (s service) Format() string {
	return "Format orders"
}
func (s service) Request() string {
	return "Request orders"
}
func (s service) Response() string {
	return "Response orders"
}
func (s service) Url() string {
	return "Url orders"
}

func (s service) Handle() string {
	return "/orders"
}

func (s service) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response orders")
}

var Service service
