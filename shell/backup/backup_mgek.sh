#!/usr/bin/env bash
apps_root="/home/apps"
function backup_mgek()
{
  zip -q -o -r mgek.zip "${apps_root}"/mgek || exit 1
}

backup_mgek
exit 0