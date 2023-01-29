package gopher

import "github.com/graphql-go/graphql"

type Resolver interface {
	ResolveGophers(p graphql.ResolveParams) (interface{}, error)
	ResolveGopher(p graphql.ResolveParams) (interface{}, error)
	ResolveJobs(p graphql.ResolveParams) (interface{}, error)
}
