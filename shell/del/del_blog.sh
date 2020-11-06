#!/usr/bin/env bash
apps_root="/home/apps"
function del_blog()
{
  echo "Last Clean Date: $(date)" > "${apps_root}"/blog/blog.log
}

del_blog
exit 0