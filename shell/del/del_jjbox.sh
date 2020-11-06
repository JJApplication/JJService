#!/usr/bin/env bash
apps_root="/home/apps"
function del_jjbox()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/jjbox/jjbox.log
}

del_jjbox
exit 0