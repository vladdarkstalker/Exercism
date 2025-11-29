#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "false"
    exit 1
fi

strin=$(echo -n "$*" | tr '[:upper:]' '[:lower:]')
letters=$(echo "$strin" | grep -o '[a-z]')
uniq=$(echo "$letters" | sort | uniq | tr -d '\n')

if [ ${#uniq} -eq 26 ]; then
    echo "true"
else
    echo "false"
fi