#!/usr/bin/env bash
apps_root="/home/apps"
function del_jjgo()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/jjgo/logs/jjgo.log
}

del_jjgo
exit 0