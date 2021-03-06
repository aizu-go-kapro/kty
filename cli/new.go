package main

import (
	"github.com/pkg/errors"
	"fmt"
	"github.com/aizu-go-kapro/kty/user"
	"github.com/aizu-go-kapro/kty/service"
	"bufio"
	"os"
	"strings"
)

/** Send サブコマンド用の実装 **/
type NewUser struct{}

func (n *NewUser) Help() string {
	return `
	Register user information of recipient

	user name
	slack id (channel id)
	twitter id (You do not have to turn on @)
`
}

func (n *NewUser) Run(args []string) int {
	fmt.Print("User Name >")
	u := user.NewUser(StrStdin())

	for _, v := range []service.ServiceID{service.SlackID, service.TwitterID}{
		fmt.Printf("%s ID >", v)
		u.AddService(v, StrStdin())
	}

	if err := writeJson(rootdr, u); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return 1
	}
	fmt.Println("New User: ", u.Name)
	return 0
}

func (n *NewUser) Synopsis() string {
	return "create new user"
}

func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return stringInput
}


func (n *NewUser)ReceiveArgs(args []string) (map[string]string, error) {

	opts := make(map[string]string)

	if !ExitOption("-u",args) {
		return nil, errors.New("user not found")
	}
	opts["-u"] = OptionJudge("-u", args)

	if !ExitOption("-m",args) {
		return nil, errors.New("message not found")
	}
	opts["-m"] = OptionJudge("-m", args)

	if !ExitOption("-s",args) {
		return nil, errors.New("service not found")
	}
	opts["-s"] = OptionJudge("-s", args)

	return opts, nil
}