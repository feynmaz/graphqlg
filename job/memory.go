package job

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	jobs map[string][]Job
	sync.Mutex
}

func NewMemoryRepository() *InMemoryRepository {
	jobs := make(map[string][]Job)

	jobs["1"] = []Job{
		{
			ID:         "123-123",
			EmployeeID: "1",
			Company:    "Google",
			Title:      "Logo",
			Start:      "2021-01-01",
			End:        "",
		},
	}
	jobs["2"] = []Job{
		{
			ID:         "124-124",
			EmployeeID: "2",
			Company:    "Google",
			Title:      "Janitor",
			Start:      "2021-05-03",
			End:        "",
		}, {
			ID:         "125-125",
			EmployeeID: "2",
			Company:    "Microsoft",
			Title:      "Janitor",
			Start:      "1980-03-04",
			End:        "2021-05-02",
		},
	}
	return &InMemoryRepository{
		jobs: jobs,
	}
}

func (imr *InMemoryRepository) GetJobs(employeeID, companyName string) ([]Job, error) {
	if jobs, ok := imr.jobs[employeeID]; ok {
		filtered := make([]Job, 0)
		for _, job := range jobs {
			if (job.Company == companyName) || companyName == "" {
				filtered = append(filtered, job)
			}
		}
		return filtered, nil
	}
	return nil, errors.New("no such employee exist")
}
