package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/feynmaz/graphqlg/gopher"
	"github.com/feynmaz/graphqlg/job"
	"github.com/feynmaz/graphqlg/schemas"
)

func main() {
	gopherService := gopher.NewService(
		gopher.NewMemoryRepository(),
		job.NewMemoryRepository(),
	)
	schema, err := schemas.GenerateSchema(&gopherService)
	if err != nil {
		panic(err)
	}

	StartServer(schema)
}

func StartServer(schema *graphql.Schema) {
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
