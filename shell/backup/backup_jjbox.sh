#!/usr/bin/env bash
apps_root="/home/apps"
function backup_jjbox()
{
  zip -q -o -r jjbox.zip "${apps_root}"/jjbox || exit 1
}

backup_jjbox
exit 0