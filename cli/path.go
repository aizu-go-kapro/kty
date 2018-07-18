package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aizu-go-kapro/kty/user"
)

func Exists(drname string) bool {
	_, err := os.Stat(drname)
	return err == nil
}

func createfunc(rootdr string) error {
	if !Exists(rootdr + ".kty") {
		err := os.Mkdir(rootdr+".kty", 0705)
		if err != nil {
			return err
		}
		err = os.Mkdir(rootdr+".kty/user", 0705)
		if err != nil {
			return err
		}

		writeMaster(rootdr, user.MasterUser{})

	}
	// 普通はnilを戻す
	return nil
}

func writeMaster(rootdr string, data user.MasterUser) {
	path := rootdr + ".kty/master.json"
	fout, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	outputJson, err := json.Marshal(&data)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
}

func writeJson(rootdr string, data *user.User) error {
	path := rootdr + ".kty/user/" + data.Name + ".json"
	fout, err := os.Create(path)
	defer fout.Close()

	if err != nil {
		return err
	}

	outputJson, err := json.Marshal(data)
	fout.Write([]byte(outputJson))
	if err != nil {
		return err
	}

	return nil
}

//
//func main() {
//}
