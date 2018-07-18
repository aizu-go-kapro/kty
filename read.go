package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aizu-go-kapro/kty/user"
)

func main() {
	rootdr := os.Getenv("HOME") + "/"
	u := user.User{}
	path := rootdr + ".kty/user/" + u.Name + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var mu user.JsonUser

	if err := json.Unmarshal(data, &mu); err != nil {
		fmt.Println(err)
	}

	fmt.Println(mu.SlackChannelID)
	fmt.Println(mu.TwitterID)

}
