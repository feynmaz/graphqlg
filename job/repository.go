package job

type Repository interface {
	GetJobs(employeeID, companyName string) ([]Job, error)
}
