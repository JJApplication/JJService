#!/usr/bin/env bash
# 保证启动前该程序已经停止 否则会出现端口占用错误
bash_pid=""
function start_mysite()
{
  ifstart=$(ps ax | grep mysite | grep -v grep)
  if [ "${ifstart}" == "" ];then {
    #启动操作 否则什么都不干
    cd /home/web/mysite && nohup gunicorn -c gun_mysite.py mysite:app &
    bash_pid=$!
    str=$"\n"
    sstr=$(echo -e $str)
    echo "${sstr}"
  }
  fi
    str=$"\n"
    sstr=$(echo -e $str)
    echo "${sstr}"
}

function kill_bash()
{
    str=$"\n"
    sstr=$(echo -e $str)
    echo "${sstr}"
    if [ $bash_pid != "" ];then {
      kill -9 $bash_pid
    }
    fi
}

start_mysite
kill_bash
exit 0