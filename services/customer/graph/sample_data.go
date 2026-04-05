package graph

import "github.com/karthikprocoder/federated-graphql-api/services/customer/graph/model"

var customers = map[string]*model.Customer{
	"1": {ID: "1", Name: "Alice", Email: "alice@test.com"},
	"2": {ID: "2", Name: "Bob", Email: "bob@test.com"},
}