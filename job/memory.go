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

func (imr *InMemoryRepository) GetJob(employeeID, jobID string) (Job, error) {
	if jobs, ok := imr.jobs[employeeID]; ok {
		for _, job := range jobs {

			if job.ID == jobID {
				return job, nil
			}
		}
		return Job{}, errors.New("no such job exist")
	}
	return Job{}, errors.New("no such employee exist")
}

func (imr *InMemoryRepository) Update(j Job) (Job, error) {
	imr.Lock()
	defer imr.Unlock()

	if jobs, ok := imr.jobs[j.EmployeeID]; ok {
		for i, job := range jobs {

			if job.ID == j.ID {
				imr.jobs[j.EmployeeID][i] = j
				return j, nil
			}
		}
	}
	return Job{}, errors.New("no such employee exist")
}
