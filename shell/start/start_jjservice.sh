#!/usr/bin/env bash
# 保证启动前该程序已经停止 否则会出现端口占用错误

apps_root="/home/apps"
function start_jjservice()
{
  ifstart=$(pgrep jjservice)
  if [ "${ifstart}" == "" ];then {
    #启动操作 否则什么都不干
    cd "${apps_root}"/jjservice && nohup ./jjservice > /dev/null 2>&1 &
    exit 1
  }
  else
    # echo "running"
    exit 2
  fi

}


start_jjservice

exit 0