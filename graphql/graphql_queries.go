package graphql

import (
	"github.com/231tr0n/graphql-example/database"

	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "query_database",
	Fields: graphql.Fields{
		"get": &graphql.Field{
			Type: graphqlKeyValuePairType,
			Args: graphql.FieldConfigArgument{
				"key": graphqlKeyArgument,
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var key, ok = params.Args["key"].(string)
				var entryStruct keyValuePair
				if ok {
					var value, ok = database.Get(key)
					if ok {
						entryStruct = keyValuePair{
							Key:   key,
							Value: value,
						}
					}
				}
				return entryStruct, nil
			},
		},
		"list": &graphql.Field{
			Type: graphqlKeyValuePairArrayType,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var entryStructs []keyValuePair
				var entries = database.List()
				for _, entry := range entries {
					entryStructs = append(entryStructs, keyValuePair{
						Key:   entry[0],
						Value: entry[1],
					})
				}
				return entryStructs, nil
			},
		},
	},
})
