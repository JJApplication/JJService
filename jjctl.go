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
)

// cli控制器

func main()  {
	flag := os.Args
	if len(flag) <= 1 || len(flag) > 2 {
		fmt.Println("require args like ...[start|restart|stop|status|init]")
		fmt.Println("[start]   -> 启动jjservice服务")
		fmt.Println("[stop]    -> 停止jjservice服务")
		fmt.Println("[restart] -> 重启jjservice服务")
		fmt.Println("[status]  -> 查看jjservice服务")
		fmt.Println("[init]    -> 初始化服务脚本")
	}else {
		switch flag[1] {
		case "start":
			cmd0 := exec.Command("bash", "-c", "pgrep jjservice")
			o, _ := cmd0.Output()
			if len(o) <= 0 {
				// 无pid即服务未启动
				cmd := exec.Command("bash", "-c", "cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err.Error())
				}else {
					fmt.Println("启动成功")
				}
				// get pid
				// 获取bash的进程pid 关闭由nohup开启的进程
				cmd1 := exec.Command("bash", "-c", "ps ax | grep 'bash -c cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &' | grep -v grep | awk '{print $1}'")
				o1, err1 := cmd1.Output()
				if err1 == nil {
					pid := string(o1[:])
					cmd2 := exec.Command("bash", "-c", "kill " + pid)
					_ = cmd2.Run()
				}

			}else {
				fmt.Println("服务已在运行")
			}

		case "stop":
			cmd0 := exec.Command("bash", "-c", "pgrep jjservice")
			o, _ := cmd0.Output()
			if len(o) <= 0 {
				fmt.Println("服务已停止")
			}else {
				pid := string(o[:])
				cmd := exec.Command("bash", "-c", "kill "+pid)
				err := cmd.Run()
				if err != nil {
					fmt.Println(err.Error())
				}else {
					fmt.Println("服务已停止")
				}
			}
		case "restart":
			cmd0 := exec.Command("bash", "-c", "pgrep jjservice")
			o, _ := cmd0.Output()
			if len(o) > 0 {
				// 先停止再启动
				fmt.Println("服务正在运行")
				pid := string(o[:])
				cmd1 := exec.Command("bash", "-c", "kill " + pid)
				err1 := cmd1.Run()
				if err1 != nil {
					fmt.Println(err1.Error())
				}else {
					fmt.Println("服务已停止")
				}

				cmd2 := exec.Command("bash", "-c", "cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &")
				err2 := cmd2.Run()
				if err2 != nil {
					fmt.Println(err2.Error())
				}else {
					fmt.Println("服务启动成功")
				}
				// get pid
				cmd3 := exec.Command("bash", "-c", "ps ax | grep 'bash -c cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &' | grep -v grep | awk '{print $1}'")
				o3, err3 := cmd3.Output()
				if err3 == nil {
					pid := string(o3[:])
					cmd4 := exec.Command("bash", "-c", "kill " + pid)
					_ = cmd4.Run()
				}
			}else {
				cmd3 := exec.Command("bash", "-c", "cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &")
				err3 := cmd3.Run()
				if err3 != nil {
					fmt.Println(err3.Error())
				}else {
					fmt.Println("服务启动成功")
				}
				// get pid
				cmd5 := exec.Command("bash", "-c", "ps ax | grep 'bash -c cd /home/apps/jjservice && nohup ./jjservice > /dev/null 2>&1 &' | grep -v grep | awk '{print $1}'")
				o5, err5 := cmd5.Output()
				if err5 == nil {
					pid := string(o5[:])
					cmd6 := exec.Command("bash", "-c", "kill " + pid)
					_ = cmd6.Run()
				}
			}

		case "init":
			cmd := exec.Command("bash", "-c", "chmod -R 655 /home/apps/jjservice/shell")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("脚本已提权")
			}

		case "status":
			cmd0 := exec.Command("bash", "-c", "pgrep jjservice")
			o, _ := cmd0.Output()
			pid := string(o[:])
			if len(o) <= 0 {
				fmt.Println("JJService is not running")
			}else {
				fmt.Println("JJService is running on PID: " + pid)
			}

		default:
			fmt.Println("check your args...")
		}
	}

}