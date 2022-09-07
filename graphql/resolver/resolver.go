package resolver

import (
	"adm-num/ent"
	"adm-num/graphql/generated"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var newSchema *graphql.ExecutableSchema

type Resolver struct {
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {

	if newSchema == nil {
		newSchema = new(graphql.ExecutableSchema)
		*newSchema = generated.NewExecutableSchema(generated.Config{
			Resolvers: &Resolver{
				client: client,
			},
		})
	}
	return *newSchema
}
