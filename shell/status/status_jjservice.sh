#!/usr/bin/env bash
# return PIDs
opt=""
function status_jjservice()
{
    opt=$(pgrep jjservice)
    echo "${opt}"
}

status_jjservice
exit 0