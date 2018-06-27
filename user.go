package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	Name    string
	Service map[ServiceID]Conf
}

type Conf map[string]string

func (u *User) Send(sid ServiceID, message string) error {
	const errtag = "User.Send failed "
	conf, ok := u.Service[sid]
	if !ok {
		return errors.Wrap(errors.New("not found service"), errtag)
	}
	fmt.Println(conf)
	//
	//s := GetService(sid)
	//
	//if err := s.Send(conf, message); err != nil {
	//	return errors.Wrap(err, errtag)
	//}

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
