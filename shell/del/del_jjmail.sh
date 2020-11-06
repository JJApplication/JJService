#!/usr/bin/env bash
apps_root="/home/apps"
function del_jjmail()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/jjmail/jjmail.log
}

del_jjmail
exit 0