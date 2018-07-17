package user

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/aizu-go-kapro/kty/service/slack"
	"github.com/aizu-go-kapro/kty/service"
	"github.com/aizu-go-kapro/kty/service/twitter"
)

type User struct {
	Name    string
	Service map[service.ServiceID]Conf
}

func NewUser(name string, sc map[service.ServiceID]Conf)*User{
	return &User{
		Name: name,
		Service: sc,
	}
}

var servises = map[string]service.Service{
	"slack": &slack.Slack{},
	"twitter": &twitter.Twitter{},
}

type Conf map[string]string


func (u *User) Send(sid service.ServiceID, message string) error {
	const errtag = "User.Send failed "
	conf, ok := u.Service[sid]
	if !ok {
		return errors.Wrap(errors.New("not found service"), errtag)
	}

	s := GetService(sid)
	cf := map[string]string(conf)

	if err := s.Send(cf, message); err != nil {
		return errors.Wrap(err, errtag)
	}

	return nil
}

func (u *User) SendAll(message string) error {
	const errtag = "User.SendAll failed "

	for k := range u.Service {
		if err := u.Send(k, message); err != nil {
			msg := fmt.Sprintf("%s service id %d", errtag, k)
			errors.Wrap(err, msg)
		}
	}

	return nil
}

func GetService(sid service.ServiceID) service.Service {
	switch sid {
	case service.SlackID:
		return slack.New()
	case service.TwitterID:
		return twitter.New()
	}
	panic("not found")
}
