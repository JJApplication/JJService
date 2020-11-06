#!/usr/bin/env bash

opt=""
function status_mgekfile()
{
    opt=$(lsof -i :8011)
    echo "${opt}"
}

status_mgekfile
exit 0