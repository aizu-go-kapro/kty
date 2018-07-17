package user

import (
	"fmt"

	"github.com/aizu-go-kapro/kty/service"
	"github.com/aizu-go-kapro/kty/service/slack"
	"github.com/aizu-go-kapro/kty/service/twitter"
	"github.com/pkg/errors"
)

type User struct {
	Name    string
	Service map[service.ServiceID]Conf
}

type MasterUser struct {
	SlackTeamName string `json:"slack_team_name"`
	TwitterID     string `json:"twitter_id"`
}

func NewUser(name string) *User {
	return &User{
		Name:    name,
		Service: map[service.ServiceID]Conf{},
	}
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

var servises = map[service.ServiceID]service.Service{
	service.SlackID:   &slack.Slack{},
	service.TwitterID: &twitter.Twitter{},
}

//TODO 今回はどちらも必要なものが一つなので、stringを入れるようにする
func (u *User) AddService(sid service.ServiceID, s string) {

	c := Conf{
		servises[sid].TokenKey(): s,
	}

	u.Service[sid] = c
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
