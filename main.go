package main

import "fmt"

func main() {
	c := Conf{
		"slackChannelID": "DA8FGSVUH",
	}
	sc := map[ServiceID]Conf{
		SlackID: c,
	}
	u := NewUser("yusuke", sc)
	if err := u.Send(SlackID, "タコ助"); err != nil {
		fmt.Println(err)
	}
}
