package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aizu-go-kapro/kty/user"
)

func main() {
	rootdr := os.Getenv("HOME") + "/"
	u := user.User{}
	u.Name = "daisuke"
	path := rootdr + ".kty/user/" + u.Name + ".json"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}
