package job

type Repository interface {
	GetJobs(employeeID, companyName string) ([]Job, error)
	GetJob(employeeID, jobID string) (Job, error)
	Update(Job) (Job, error)
}
