#!/usr/bin/env bash
apps_root="/home/apps"
function backup_jjgo()
{
  zip -q -o -r jjgo.zip "${apps_root}"/jjgo || exit 1
}

backup_jjgo
exit 0