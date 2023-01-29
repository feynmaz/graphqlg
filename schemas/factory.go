package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/feynmaz/graphqlg/gopher"
)

var jobType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Job",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"employeeID": &graphql.Field{
			Type: graphql.String,
		},
		"company": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"start": &graphql.Field{
			Type: graphql.String,
		},
		"end": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func generateGopherType(gs *gopher.GopherService) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Gopher",
		// Fields is the field values to declare the structure of the object
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID that is used to identify unique gophers",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the gopher",
			},
			"hired": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "True if the Gopher is employeed",
			},
			"profession": &graphql.Field{
				Type:        graphql.String,
				Description: "The gophers last/current profession",
			},
			"jobs": &graphql.Field{
				Type:        graphql.NewList(jobType),
				Description: "A list of all jobs the gopher had",
				Resolve:     gs.ResolveJobs,
				Args: graphql.FieldConfigArgument{
					"company": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
		}})
}

func GenerateSchema(gs *gopher.GopherService) (*graphql.Schema, error) {
	gopherType := generateGopherType(gs)

	fields := graphql.Fields{
		// We define the Gophers query
		"gophers": &graphql.Field{
			// It will return a list of GopherTypes, a List is an Slice
			Type: graphql.NewList(gopherType),
			// We change the Resolver to use the gopherRepo instead, allowing us to access all Gophers
			Resolve: gs.ResolveGophers,
			// Description explains the field
			Description: "Query all Gophers",
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	rootMutation := generateRootMutation(gs)
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: rootMutation,
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
