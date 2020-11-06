/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-06
*/
package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func RecordPID() error {
	pid := os.Getpid()
	cwd, _ := os.Getwd()
	file := path.Join(cwd, "jjservice.pid")
	err := ioutil.WriteFile(file, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
