#!/usr/bin/env bash
apps_root="/home/apps"
function del_jjservice()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/jjservice/jjservice.log
}

del_jjservice
exit 0