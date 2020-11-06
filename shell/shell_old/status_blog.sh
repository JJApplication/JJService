#!/usr/bin/env bash

opt=""
function status_blog()
{
    opt=$(pgrep app_blog)
    echo "${opt}"
}

status_blog
exit 0