#!/usr/bin/env bash

function stop_blog()
{
  pids=(`pgrep app_blog | awk '{print $1}'`)
  array=${pids[@]}
  arraypid=($array)
  # 重定向错误信息输出，自定义错误
  kill "${arraypid[0]}" > /dev/null 2>&1
  err1=$?
  kill "${arraypid[1]}" > /dev/null 2>&1
  err2=$?
  if [ $err1 == 0 ] || [ $err2 == 0 ];then
  {
    echo "successfully stopped"
  }
  else
  {
    echo "stop service failed"
  }
  fi
}

stop_blog
exit