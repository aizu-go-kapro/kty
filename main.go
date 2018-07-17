package main

import (
	"fmt"
	"github.com/aizu-go-kapro/kty/service"
	"github.com/aizu-go-kapro/kty/user"
)

func main() {

	u := user.NewUser("yusuke")
	u.AddService(service.SlackID, "GBSL4FV9V")
	u.AddService(service.TwitterID, "tastykusa")

	if err := u.Send(service.SlackID, "タコ助"); err != nil {
		fmt.Println(err)
	}

	if err := u.SendAll("ほげ助"); err != nil {
		fmt.Println(err)
	}

}
