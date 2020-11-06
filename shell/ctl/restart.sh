#!/usr/bin/env bash

pid=(`pgrep jjservice`)
kill "${pid}" || echo "can't find jjservice's pid."
nohup ./jjservice > /dev/null 2>&1 &
echo "jjservice has restarted."
exit