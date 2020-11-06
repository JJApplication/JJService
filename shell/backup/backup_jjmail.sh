#!/usr/bin/env bash
apps_root="/home/apps"
function backup_jjmail()
{
  zip -q -o -r jjmail.zip "${apps_root}"/jjmail || exit 1
}

backup_jjmail
exit 0