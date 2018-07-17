package service

type ServiceID string

const (
	SlackID ServiceID = "slack"
	TwitterID ServiceID = "twitter"
)


// func GetService(s ServiceID) Service {
//
//
// }

type Service interface {
	Send(token map[string]string, message string)error
	TypeID() string
}
