package job

type Repository interface {
	GetJobs(employeeID string) ([]Job, error)
}
