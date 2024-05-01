package models

const (
	CONTROLLER int = iota
  QUEUE_CONSUMER
  QUEUE_PRODUCER
  REPOSITORY
  SERVICE
)

type WizardAnswers struct {
	Controller    bool
	ModuleName    string
	QueueConsumer bool
	QueueProducer bool
	Repository    bool
	Service       bool
}
