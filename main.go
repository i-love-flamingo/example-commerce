package main

import (
	projectGraphql "example-commerce/graphql"
	"flamingo.me/dingo"
	"flamingo.me/flamingo-commerce/v3/cart"
	"flamingo.me/flamingo-commerce/v3/category"
	"flamingo.me/flamingo-commerce/v3/checkout"
	"flamingo.me/flamingo-commerce/v3/product"
	"flamingo.me/flamingo/v3"
	"flamingo.me/graphql"
)

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

func main() {
	flamingo.App([]dingo.Module{
		new(product.Module),
		new(cart.Module),
		new(category.Module),
		new(checkout.Module),
		new(graphql.Module),
		new(projectGraphql.Module),
	})
}
