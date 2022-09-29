package graph

import "mock-graphql-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DrinkStore map[string]model.Drink
}
