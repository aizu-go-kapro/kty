package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

/** Send サブコマンド用の実装 **/
type Send struct{}

func (f *Send) Help() string {
	return "app foo"
}

func (f *Send) Run(args []string) int {

	info := ReceiveArgs(args)

	for _, v := range info {
		fmt.Println(v)
	}

	return 0

}

func (f *Send) Synopsis() string {
	return "Print \"Foo!\""
}

func ReceiveArgs(args []string) []string {

	info := make([]string, 0, len(args))

	info = append(info, OptionJudge("-u", args))
	info = append(info, OptionJudge("-m", args))
	info = append(info, OptionJudge("-s", args))

	return info

}

type Options struct {
	option1 string
	args    []string
}

func OptionJudge(opt string, args []string) string {

	var info string

	for i := range args {
		if opt != args[i] {
			continue
		}

		idx := GetIOptIndex(opt, args) + 1

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

func main() {
	// コマンドの名前とバージョンを指定
	c := cli.NewCLI("app", "1.0.0")

	// サブコマンドの引数を指定
	c.Args = os.Args[1:]

	/*ReceiveArgs(os.Args)*/

	// サブコマンド文字列 と コマンド実装の対応付け
	c.Commands = map[string]cli.CommandFactory{
		"send": func() (cli.Command, error) {
			return &Send{}, nil
		},
	}

	// コマンド実行
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

}
