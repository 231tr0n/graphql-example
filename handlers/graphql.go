package handlers

import (
	"github.com/231tr0n/graphql-example/graphql"
	"github.com/graphql-go/handler"
)

var Graphql *handler.Handler = handler.New(&handler.Config{
	Schema:     &graphql.DatabaseSchema,
	GraphiQL:   true,
	Playground: false,
	Pretty:     true,
})
