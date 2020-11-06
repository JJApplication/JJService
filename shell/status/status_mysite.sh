#!/usr/bin/env bash

opt=""
function status_mysite()
{
    opt=$(ps ax | grep mysite | grep -v grep | grep -v bash)
    echo "${opt}"
}

status_mysite
exit 0