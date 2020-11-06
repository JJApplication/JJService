#!/usr/bin/env bash
apps_root="/home/apps"
function del_mgek()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/mgek/mgek.log
}

del_mgek
exit 0