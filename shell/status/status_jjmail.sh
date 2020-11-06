#!/usr/bin/env bash

opt=""
function status_jjmail()
{
    opt=$(pgrep celery)
    echo "${opt}"
}

status_jjmail
exit 0