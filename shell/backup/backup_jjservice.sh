#!/usr/bin/env bash
apps_root="/home/apps"
function backup_jjservice()
{
  zip -q -o -r jjservice.zip "${apps_root}"/jjservice || exit 1
}

backup_jjservice
exit 0