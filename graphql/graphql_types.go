package graphql

import "github.com/graphql-go/graphql"

var graphqlKeyType = graphql.String

var graphqlKeyArrayType = graphql.NewList(graphql.String)

var graphqlValueType = graphql.String

var graphqlValueArrayType = graphql.NewList(graphql.String)

var graphqlKeyValuePairType = graphql.NewObject(graphql.ObjectConfig{
	Name: "keyValuePair",
	Fields: graphql.Fields{
		"key": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var graphqlKeyValuePairArrayType = graphql.NewList(graphqlKeyValuePairType)
