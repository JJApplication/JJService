#!/usr/bin/env bash
apps_root="/home/apps"
function del_mysite()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/mysite/mysite.log
}

del_mysite
exit 0