package service

import "net/http"

type Service interface {
	Id() string
	Sequence() int64
	Format() string
	Response() string
	Request() string
	Url() string
	Handle() string
	HandlerFunc(http.ResponseWriter, *http.Request)
}
