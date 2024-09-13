package graph

import "go-graphql-apollo-federation/items/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ItemDB *db.Queries
}
