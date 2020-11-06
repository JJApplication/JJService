/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-06
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

// cli控制器

func main()  {
	flag := os.Args
	if len(flag) <= 1 || len(flag) > 2 {
		fmt.Println("require args like ...[start|restart|stop|init]")
	}else {
		cwd, _ := os.Getwd()
		shPath := path.Join(cwd, "shell/ctl")
		switch flag[1] {
		case "start":
			cmd := exec.Command("bash", shPath+"/start.sh")
			o, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println(string(o[:]))
			}
		case "stop":
			cmd := exec.Command("bash", shPath+"/stop.sh")
			o, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println(string(o[:]))
			}
		case "restart":
			cmd := exec.Command("bash", shPath+"/restart.sh")
			o, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println(string(o[:]))
			}
		case "init":
			cmd := exec.Command("bash", shPath+"/init.sh")
			o, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println(string(o[:]))
			}
		default:
			fmt.Println("check your args...")

		}
	}

}