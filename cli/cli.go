package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
	"fmt"
)

var (
	rootdr string
)

func init(){
	rootdr = os.Getenv("HOME") + "/"
	if err := createfunc(rootdr); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func ExitOption(opt string, args []string) bool {
	for _, v := range args {
		if v == opt {
			return true
		}
	}
	return false
}

func main() {
	// コマンドの名前とバージョンを指定
	c := cli.NewCLI("app", "1.0.0")

	// サブコマンドの引数を指定
	c.Args = os.Args[1:]

	// サブコマンド文字列 と コマンド実装の対応付け
	c.Commands = map[string]cli.CommandFactory{
		"send": func() (cli.Command, error) {
			return &Send{}, nil
		},
		"new": func()(cli.Command, error){
			return &NewUser{}, nil
		},
	}

	// コマンド実行
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

}
