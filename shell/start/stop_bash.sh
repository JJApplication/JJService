#!/usr/bin/env bash
# 关闭由nohup开启的进程

running_sh=$1 # 第一个参数即脚本名称
# 记得排除自己
shell_self="stop_bash.sh"
pid=$(ps ax | grep "${running_sh}" | grep -v grep | grep -v $shell_self | awk '{p=p$1 " "}END{print p}')
# 对于多进程的情况需要转换为字符串
if [ "$pid" != "" ];then
  kill "$pid" || exit
fi

exit