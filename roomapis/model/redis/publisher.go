package redis

type PublisherService interface {
	PublishMessage(chanel string, message string) error
}
