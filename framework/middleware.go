package framework

import "net/http"

type MiddlewareFunc func(*Context, http.Handler)

type Middleware interface {
	GetMWFunc() MiddlewareFunc
}
