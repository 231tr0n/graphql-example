package graphql

import "github.com/graphql-go/graphql"

var graphqlKeyArgument = &graphql.ArgumentConfig{
	Type: graphqlKeyInputType,
}

var graphqlKeyArrayArgument = &graphql.ArgumentConfig{
	Type: graphqlKeyArrayInputType,
}

var graphqlValueArgument = &graphql.ArgumentConfig{
	Type: graphqlValueInputType,
}

var graphqlValueArrayArgument = &graphql.ArgumentConfig{
	Type: graphqlValueArrayInputType,
}

var graphqlKeyValuePairArgument = &graphql.ArgumentConfig{
	Type: graphqlKeyValuePairInputType,
}

var graphqlKeyValuePairArrayArgument = &graphql.ArgumentConfig{
	Type: graphqlKeyValuePairArrayInputType,
}
