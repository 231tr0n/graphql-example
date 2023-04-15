package graphql

import (
	"github.com/graphql-go/graphql"
	"log"
)

var DatabaseSchema graphql.Schema

func init() {
	var err error
	DatabaseSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
