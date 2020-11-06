#!/usr/bin/env bash

opt=""
function status_jjgo()
{
    opt=$(pgrep jjgo)
    echo "${opt}"
}

status_jjgo
exit 0