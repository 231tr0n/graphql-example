package routes

import (
	"github.com/231tr0n/graphql-example/handlers"
	"net/http"
)

var Mux http.ServeMux

func init() {
	Mux.Handle("/", handlers.Graphql)
}
