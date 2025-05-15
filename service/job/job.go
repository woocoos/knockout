package job

type Job interface {
	JobName() string
	JobFunc()
	CronSpec() string
}
