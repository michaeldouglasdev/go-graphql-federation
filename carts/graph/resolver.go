package graph

import "go-graphql-apollo-federation/carts/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CartDB *db.Queries
}
