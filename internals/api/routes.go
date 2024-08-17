package api

import (
	"net/http"

	"gihub.com/prastavna/form-to-sheet/internals/api/handlers"
	"gihub.com/prastavna/form-to-sheet/internals/api/middlewares"
)

func addMiddleware(handler http.HandlerFunc) http.Handler {
	return middlewares.CheckMethod(middlewares.RateLimiter(middlewares.Cors(http.HandlerFunc(handler))))
}

func Routes() {
	http.Handle("/submit", addMiddleware(handlers.SubmitHandler))
}
