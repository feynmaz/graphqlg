package schemas

import (
	"github.com/feynmaz/graphqlg/gopher"
	"github.com/graphql-go/graphql"
)

var modifyJobArgs = graphql.FieldConfigArgument{
	"employeeid": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"jobid": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"start": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"end": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func generateRootMutation(gs *gopher.GopherService) *graphql.Object {
	mutationFields := graphql.Fields{
		"modifyJob": &graphql.Field{
			Type:        jobType,
			Resolve:     gs.MutateJobs,
			Description: "Modify a job for a gopher",
			Args:        modifyJobArgs,
		},
	}
	mutationConfig := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}

	return graphql.NewObject(mutationConfig)
}
