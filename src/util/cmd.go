/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-25
*/

/*
脚本执行库默认使用内置cmd执行shell命令
在此情况失效时使用python执行，多了一步python解释器操作
 */

package util

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

// 执行shell脚本
func RunShell(cmd string) (output string, e error) {
	// if failed return error
	p, _ := os.Getwd()
	shellPath := path.Join(p, cmd)
	shell := exec.Command("bash", "-c", shellPath)
	optByte, err := shell.Output()
	opt := strings.TrimSpace(string(optByte[:]))
	// 输出是字节流最好转换为string字符串
	if err != nil {
		return opt, err
	}
	return opt, nil
}

func RunShellAsync(cmd string) (e error) {
	p, _ := os.Getwd()
	shellPath := path.Join(p, cmd)
	shell := exec.Command("bash", "-c", shellPath)
	err := shell.Run()

	// 异步执行没有返回值
	if err != nil {
		return err
	}
	// 关闭进程

	return
}