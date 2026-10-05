package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	GlobalMiddleWares []Middleware
}

func NewMiddleware() *Manager {
	return &Manager{
		GlobalMiddleWares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(globalMiddleware ...Middleware) {
	mngr.GlobalMiddleWares = append(mngr.GlobalMiddleWares, globalMiddleware...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	n := next

	for _, globalMiddleware := range mngr.GlobalMiddleWares {
		n = globalMiddleware(n)
	}

	return n
}
