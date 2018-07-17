package main

import (
	"fmt"
	"github.com/aizu-go-kapro/kty/service"
	"github.com/aizu-go-kapro/kty/user"
)

func main() {
	u := user.NewUser("yusuke")

	u.AddService(service.SlackID, "GBSL4FV9V")

	if err := u.Send(service.SlackID, "タコ助"); err != nil {
		fmt.Println(err)
	}
}
