package main

type ServiceID string

const (
	SlackID ServiceID = "slack"
)


// func GetService(s ServiceID) Service {
//
//
// }

type Service interface {
	Send(token map[string]string, message string)error
	TypeID() string
}
