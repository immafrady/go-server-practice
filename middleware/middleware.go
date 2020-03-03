package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	r := &Router{}
	r.mux = make(map[string]http.Handler)
	return r
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

func (r *Router) ListenAndServe(a string, h http.Handler) (err error) {
	for route, handler := range r.mux {
		http.Handle(route, handler)
	}

	err = http.ListenAndServe(a, h)
	return
}
