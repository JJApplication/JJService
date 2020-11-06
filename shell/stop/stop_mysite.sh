#!/usr/bin/env bash

function stop_mysite()
{
  pids=(`ps ax | grep mysite | grep -v grep | grep -v bash | awk '{print $1}'`)
  array=${pids[@]}
  arraypid=($array)
  firstpid=${arraypid[0]}
  # 重定向错误信息输出，自定义错误
  kill "${firstpid}" > /dev/null 2>&1
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

stop_mysite
exit