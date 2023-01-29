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

	company := ""
	if value, ok := p.Args["company"]; ok {
		company, ok = value.(string)
		if !ok {
			return nil, errors.New("id has to be a string")
		}
	}

	jobs, err := gs.jobs.GetJobs(g.ID, company)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (gs *GopherService) MutateJobs(p graphql.ResolveParams) (interface{}, error) {
	employeeID, err := grabStringArgument("employeeid", p.Args, true)
	if err != nil {
		return nil, err
	}

	jobID, err := grabStringArgument("jobid", p.Args, true)
	if err != nil {
		return nil, err
	}

	start, err := grabStringArgument("start", p.Args, false)
	if err != nil {
		return nil, err
	}

	end, err := grabStringArgument("end", p.Args, false)
	if err != nil {
		return nil, err
	}

	// Get the job
	job, err := gs.jobs.GetJob(employeeID, jobID)
	if err != nil {
		return nil, err
	}
	if start != "" {
		job.Start = start
	}
	if end != "" {
		job.End = end
	}
	return gs.jobs.Update(job)
}

func grabStringArgument(k string, args map[string]interface{}, required bool) (string, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}
