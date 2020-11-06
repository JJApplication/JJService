#!/usr/bin/env bash

opt=""
function status_mgek()
{
    opt=$(ps ax | grep mgek_app | grep -v grep | grep -v bash)
    echo "${opt}"
}

status_mgek
exit 0