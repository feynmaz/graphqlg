package gopher

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/feynmaz/graphqlg/job"
)

type GopherService struct {
	gophers Repository
	jobs    job.Repository
}

func NewService(repo Repository, jobrepo job.Repository) GopherService {
	return GopherService{
		gophers: repo,
		jobs:    jobrepo,
	}
}

func (gs *GopherService) ResolveGophers(p graphql.ResolveParams) (interface{}, error) {
	gophers, err := gs.gophers.GetGophers()
	if err != nil {
		return nil, fmt.Errorf("failed to get gophers: %w", err)
	}
	return gophers, nil
}

func (gs *GopherService) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {
	g, ok := p.Source.(Gopher)

	if !ok {
		return nil, errors.New("source was not a Gopher")
	}

	jobs, err := gs.jobs.GetJobs(g.ID)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
