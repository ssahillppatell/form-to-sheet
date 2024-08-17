package api

import (
	"net/http"

	"github.com/prastavna/form-to-sheet/internals/api/handlers"
	"github.com/prastavna/form-to-sheet/internals/api/middlewares"
)

func Routes() {
	http.Handle("/", middlewares.RateLimiter(middlewares.Cors(http.HandlerFunc(handlers.IndexHandler))))
	http.Handle("/submit", middlewares.CheckMethod(middlewares.RateLimiter(middlewares.Cors(http.HandlerFunc(handlers.SubmitHandler)))))
}
