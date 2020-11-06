#!/usr/bin/env bash
# 因为静态站点直接部署在nginx 所以直接获取nginx的状态

function status_nginx()
{
  opt=$(pgrep nginx)
  echo $opt
}

status_nginx
exit