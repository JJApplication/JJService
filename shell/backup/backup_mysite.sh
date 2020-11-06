#!/usr/bin/env bash
apps_root="/home/apps"
function backup_mysite()
{
  zip -q -o -r mysite.zip "${apps_root}"/mysite || exit 1
}

backup_mysite
exit 0