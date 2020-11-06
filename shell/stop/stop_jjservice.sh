#!/usr/bin/env bash

function stop_jjservice()
{
  pid=(`pgrep jjservice`)
  # 重定向错误信息输出，自定义错误
  kill "${pid}" > /dev/null 2>&1
  err=$?
  if [ $err == 0 ];then
  {
    echo "successfully stopped"
  }
  else
  {
    echo "stop service failed"
  }
  fi
}

stop_jjservice
exit