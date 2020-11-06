#!/usr/bin/env bash
apps_root="/home/apps"
function backup_blog()
{
  zip -q -o -r blog.zip "${apps_root}"/blog || exit 1
}

backup_blog
exit 0