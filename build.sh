#!/usr/bin/env bash
# 构建脚本

function build_jjservice()
{
  echo "开始编译jjservice"
  env=$(go env)
  echo "$env"
  echo "start go build -ldflags '-s -w'"
  go build -ldflags "-s -w" ./src/jjservice.go
  if [ ! -f "./jjservice" ];then
  {
    echo "编译失败"
  }
  else
    echo "jjservice编译成功"
  fi
}

function build_jjctl()
{
  echo "开始编译jjctl"
  go build -ldflags "-s -w" ./jjctl.go
  if [ ! -f "./jjctl" ];then
  {
    echo "编译失败"
  }
  else
    echo "jjctl编译成功"
  fi
}

function build_package()
{
  echo "开始生成服务包"
  date=$(date "+%Y%m%d")
  path="jjservice$date"
  mkdir "$path"
  cp -R ./shell "$path"
  cp -R ./conf "$path"
  cp -R ./staticfile "$path"

  chmod +x ./jjservice
  chmod +x ./jjctl
  mv ./jjctl "$path"
  mv ./jjservice "$path"
  touch jjservice.log
  touch jjservice.pid
  mv ./jjservice.log "$path"
  mv ./jjservice.pid "$path"

  zip_name="$path.zip"
  zip_path="$path/"
  zip -r "$zip_name" "$zip_path"

  if [ ! -f "./$zip_name" ];then
  {
    echo "生成服务包失败"
  }
  else
    echo "生成服务包成功"
  fi

  if [ -d "$zip_path" ];then
    rm -rf "$zip_path"
    echo "缓存清理完毕"
  fi
}

build_jjservice
build_jjctl
build_package
exit