//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"go-graphql-apollo-federation/users/internal/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserDB *db.Queries
}
