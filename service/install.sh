#!/usr/bin/env bash
# jjapps install script

end="\033[0m"
red="\033[31m"
green="\033[32m"
yellow="\033[33m"
blue="\033[34m"

export WORK_DIRECTORY="/home/apps"
export ARCH=amd64
export GOOS=linux
export NGINX_CACHE="/home/nginx"
echo "sftp将会在后期被舍弃，改为jjbox专有服务上传文件"

echo -e "${blue}开始安装jjapps全部微服务${end}"
echo "服务包解压中..."
if [ ! -d $WORK_DIRECTORY ];then
  echo "请将服务包置于home路径下"
  echo "卸载旧版服务包"
  rm -rf $WORK_DIRECTORY
fi
unzip /home/apps.zip
echo -e "${green}服务包解压完毕${end}"

cd $WORK_DIRECTORY || echo -e "${red}服务包路径无法进入${end}" && exit 1
echo "开始安装mysite"
cd $WORK_DIRECTORY/mysite && nohup gunicorn -c gun_mysite.py mysite:app > mysite.log 2>&1 &
echo -e "${green}安装完毕${end}"

echo "开始安装mgek"
cd $WORK_DIRECTORY/mgek && nohup gunicorn -c gun_mgek.py mgek_app:app > mgek.log 2>&1 &
echo -e "${green}安装完毕${end}"

echo "开始安装blog"
echo "blog不再使用多实例运行"
cd $WORK_DIRECTORY/blog && nohup ./app_blog 8001 > blog.log 2>&1 &
echo -e "${green}安装完毕${end}"

echo "开始安装jjservice"
cd $WORK_DIRECTORY/jjservice && nohup ./jjservice > /dev/null 2>&1 &
echo -e "${green}安装完毕${end}"

echo "开始安装jjgo"
cd $WORK_DIRECTORY/jjgo && ./jjgo -r daemon
echo -e "${green}安装完毕${end}"

echo "开始安装jjmail"
cd $WORK_DIRECTORY/jjmail && nohup celery -A jjmail_server worker -l info -f jjmail.log --pool=solo > /dev/null 2>&1 &
echo -e "${green}安装完毕${end}"

echo "服务安装完毕 开始检查服务运行状态"

status_mysite=$(ps ax | grep mysite | grep -v grep | grep -v bash)
if [ -n "$status_mysite" ];then
  echo -e "${green}mysite is running...${end}"
else
  echo -e "${yellow}mysite is not running...${end}"
fi

status_mgek=$(ps ax | grep mgek | grep -v grep | grep -v bash)
if [ -n "$status_mgek" ];then
  echo -e "${green}mgek is running...${end}"
else
  echo -e "${yellow}mgek is not running...${end}"
fi

status_blog=$(pgrep app_blog)
if [ -n "$status_blog" ];then
  echo -e "${green}blog is running...${end}"
else
  echo -e "${yellow}blog is not running...${end}"
fi

status_jjservice=$(pgrep jjservice)
if [ -n "$status_jjservice" ];then
  echo -e "${green}jjservice is running...${end}"
else
  echo -e "${yellow}jjservice is not running...${end}"
fi

status_jjgo=$(pgrep jjgo)
if [ -n "$status_jjgo" ];then
  echo -e "${green}jjgo is running...${end}"
else
  echo -e "${yellow}jjgo is not running...${end}"
fi

status_jjmail=$(pgrep celery)
if [ -n "$status_jjmail" ];then
  echo -e "${green}jjmail is running...${end}"
else
  echo -e "${yellow}jjmail is not running...${end}"
fi

echo -e "${blue}Thanks for using jjservice install script!${end}"
exit