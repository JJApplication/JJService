#!/usr/bin/env bash
# 保证启动前该程序已经停止 否则会出现端口占用错误
# 分布式部署需要执行两次

function start_blog()
{
  ifstart=$(pgrep app_blog)
  if [ "${ifstart}" == "" ];then {
    #启动操作 否则什么都不干
    cd /home/web/blog && nohup ./app_blog 8001 > blog.log 2>&1 &
    cd /home/web/blog && nohup ./app_blog 8002 > blog.log 2>&1 &
  }
  else
    echo "running"
  fi
}

start_blog
exit 0