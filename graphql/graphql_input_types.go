package graphql

import "github.com/graphql-go/graphql"

var graphqlKeyInputType = graphql.String

var graphqlKeyArrayInputType = graphql.NewList(graphql.String)

var graphqlValueInputType = graphql.String

var graphqlValueArrayInputType = graphql.NewList(graphql.String)

var graphqlKeyValuePairInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "keyValuePairInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"key": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var graphqlKeyValuePairArrayInputType = graphql.NewList(graphqlKeyValuePairInputType)
