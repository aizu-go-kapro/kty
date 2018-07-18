package main

import (
	"fmt"
	"os"
	"github.com/pkg/errors"
	"io/ioutil"
	"encoding/json"
	"github.com/aizu-go-kapro/kty/user"
	"github.com/aizu-go-kapro/kty/service"
)

/** Send サブコマンド用の実装 **/
type Send struct{}

func (s *Send) Help() string {
	return "app foo"
}

func (s *Send) Run(args []string) int {

	info, err := s.ReceiveArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	u, err := readUser(info["-u"])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	if err := u.Send(service.ServiceID(info["-s"]), info["-m"]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	return 0
}

func (s *Send) Synopsis() string {
	return "Print \"Foo!\""
}

func (s *Send)ReceiveArgs(args []string) (map[string]string, error) {

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
func readUser(u string) (*user.User, error){
	path := rootdr + ".kty/user/" + u + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	 mu := user.NewUser(u)

	if err := json.Unmarshal(data, mu); err != nil {
		return nil, err
	}

	return mu, err
}


func OptionJudge(opt string, args []string) string {

	var info string

	for i := range args {
		if opt != args[i] {
			continue
		}

		idx := GetIOptIndex(opt, args) + 1
		if idx == -1 {
			return ""
		}

		if opt == "-s" {
			return args[idx]
		}

		for _, v := range args[idx:] {
			if v[0] == '-' {
				break
			}
			info += v + " "
		}

	}
	return info[:len(info)-1]
}

func GetIOptIndex(opt string, args []string) int {
	for i, v := range args {
		if v == opt {
			return i
		}
	}
	return -1
}
