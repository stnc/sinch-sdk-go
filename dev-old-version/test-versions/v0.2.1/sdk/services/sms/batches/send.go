package batches


type BatchesInterface interface {
	Send() string
	DryRun() string
}

type GroupsInterface interface {
	Create() string
}

