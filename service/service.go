package service

type ServiceID string

const (
	SlackID ServiceID = "slack"
	TwitterID ServiceID = "twitter"
)



type Service interface {
	Send(token map[string]string, message string)error
	TypeID() string
	TokenKey()string
}
