#!/usr/bin/env bash
# 保证启动前该程序已经停止 否则会出现端口占用错误

apps_root="/home/apps"

function start_mysite()
{
  ifstart=$(ps ax | grep mysite | grep -v grep | grep -v bash)
  if [ "${ifstart}" == "" ];then {
    #启动操作 否则什么都不干
    cd "${apps_root}"/mysite && nohup gunicorn -c gun_mysite.py mysite:app > mysite.log 2>&1 &
    exit 1
  }
  else
    # echo "running"
    exit 2
  fi

}


start_mysite
exit 0