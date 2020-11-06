#!/usr/bin/env bash

nohup ./jjservice > /dev/null 2>&1 &
echo "jjservice is running."
exit