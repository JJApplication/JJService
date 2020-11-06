#!/usr/bin/env bash
# 保证启动前该程序已经停止 否则会出现端口占用错误
bash_pid=""
function start_celery()
{
  ifstart=$(pgrep celery)
  if [ "${ifstart}" == "" ];then {
    #启动操作 否则什么都不干
    cd /home/web/mysite && nohup celery -A mail_server worker -l info -f celery.log --pool=solo &
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

start_celery
kill_bash
exit 0