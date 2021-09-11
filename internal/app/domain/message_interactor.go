package domain

type MessageInteractor interface {
	Message() string
	Send(message string) error
}
