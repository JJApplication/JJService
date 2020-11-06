#!/usr/bin/env bash

pid=(`pgrep jjservice`)
kill "${pid}" || echo "can't find jjservice's pid."
echo "jjservice has stopped."
exit
