#!/usr/bin/env bash

function system()
{
  cpuload=(`cat /proc/loadavg | awk '{print $3}'`)
  memory=(`awk '/MemTotal/{total=$2}/MemFree/{free=$2}END{print (total-free)/1024}' /proc/meminfo`)
  network=(`awk 'BEGIN{ORS=" "}/eth0/{print $2}/eth1/{print $2,$10}' /proc/net/dev`)
  echo "${cpuload}|${memory}|${network}"
}

system
exit 0