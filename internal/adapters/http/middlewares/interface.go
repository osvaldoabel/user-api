package middlewares

import "net/http"

type HttpMiddleware interface {
	func(http.Handler) http.Handler
}
