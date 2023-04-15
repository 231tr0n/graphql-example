package graphql

import (
	"github.com/231tr0n/graphql-example/database"

	"github.com/graphql-go/graphql"
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutate_database",
	Fields: graphql.Fields{
		"store": &graphql.Field{
			Type: graphqlKeyValuePairType,
			Args: graphql.FieldConfigArgument{
				"entry": graphqlKeyValuePairArgument,
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var entry, ok = params.Args["entry"].(map[string]interface{})
				var entryStruct keyValuePair
				if ok {
					var keyi, ok1 = entry["key"]
					var valuei, ok2 = entry["value"]
					if ok1 && ok2 {
						var key, ok3 = keyi.(string)
						var value, ok4 = valuei.(string)
						if ok3 && ok4 {
							var temp = database.Store(key, value)
							entryStruct = keyValuePair{
								Key:   temp[0],
								Value: temp[1],
							}
						}
					}
				}
				return entryStruct, nil
			},
		},
		"delete": &graphql.Field{
			Type: graphqlKeyValuePairType,
			Args: graphql.FieldConfigArgument{
				"key": graphqlKeyArgument,
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var key, ok = params.Args["key"].(string)
				var entryStruct keyValuePair
				if ok {
					var temp = database.Delete(key)
					if temp[0] != "" {
						entryStruct = keyValuePair{
							Key:   temp[0],
							Value: temp[1],
						}
					}
				}
				return entryStruct, nil
			},
		},
	},
})
