package graph

import "github.com/karthikprocoder/federated-graphql-api/services/product/graph/model"

var products = map[string]*model.Product{
	"1": {ID: "1", Name: "iPhone", Price: 999},
	"2": {ID: "2", Name: "Laptop", Price: 1500},
}
