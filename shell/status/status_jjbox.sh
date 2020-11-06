#!/usr/bin/env bash

opt=""
function status_box()
{
    opt=$(pgrep jjbox)
    echo "${opt}"
}

status_box
exit 0