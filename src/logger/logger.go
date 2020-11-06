/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-02
*/
package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

// 初始化日志

type Logger struct {
	lf *os.File // 日志文件流
	logerr error
}

// new logger
var JJLogger Logger
func init() {
	lf, err := getLog()
	JJLogger.lf = lf
	JJLogger.logerr = err
}

func (l Logger) Context(c *gin.Context, err error) {
	// 有时候传入的是err有时候是bool 所以做类型判断
	_, file, line, ok := runtime.Caller(1)
	if l.logerr != nil && ok {
		_, _ = fmt.Fprintf(l.lf, "[ERROR]@FILE {%s} LINE %s API {%s} * ERROR: %v\n",file, line, c.FullPath(), l.logerr.Error())
	}else if err != nil && ok {
		_, _ = fmt.Fprintf(l.lf, "[ERROR]@FILE {%s} LINE %s API {%s} * ERROR: %v\n",file, line, c.FullPath(), err.Error())
	}else {
		_, _ = fmt.Fprintf(l.lf, "[ERROR]@FILE {%s} LINE %s API {%s} * ERROR: %v\n",file, line, c.FullPath(), "")
	}
}

func (l Logger) Print(str string, err error) {
	if l.logerr != nil{
		_, _ = fmt.Fprintf(l.lf, "%s %s\n", str, l.logerr.Error())
	}else if err != nil {
		_, _ = fmt.Fprintf(l.lf, "%s %s\n", str, err.Error())
	}else {
		_, _ = fmt.Fprintf(l.lf, "%s %s\n", str, "")
	}
}

func (l Logger) Console(err error) {
	_, file, line, ok := runtime.Caller(1)
	if l.logerr != nil && ok {
		_, _ = fmt.Printf("[ERROR]@FILE {%s} LINE %s * ERROR: %s\n",file, line, l.logerr.Error())
	}else if err != nil && ok {
		_, _ = fmt.Printf("[ERROR]@FILE {%s} LINE %s * ERROR: %s\n",file, line, err.Error())
	}else {
		_, _ = fmt.Printf("[ERROR]@FILE {%s} LINE %s * ERROR: %s\n",file, line, "")
	}
}
