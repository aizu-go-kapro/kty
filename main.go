package main

import (
	"fmt"
	"github.com/aizu-go-kapro/kty/service"
	"github.com/aizu-go-kapro/kty/user"
)

func main() {
	c := user.Conf{
		"slackChannelID": "GBSL4FV9V",
	}
	sc := map[service.ServiceID]user.Conf{
		service.SlackID: c,
	}

	u := user.NewUser("yusuke", sc)

	if err := u.Send(service.SlackID, "タコ助"); err != nil {
		fmt.Println(err)
	}
}
