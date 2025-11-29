#!/usr/bin/env bash

res=0
size1=${#1}
size2=${#2}
if [ $# -ne 2 ]; then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
fi
if [ $size1 -ne $size2 ]; then
    echo "strands must be of equal length"
    exit 1
fi
for (( arg=0; arg<size1; arg++ )); do
    if [ "${1:$arg:1}" != "${2:$arg:1}" ]; then
        res=$((res+1))
    fi
done
echo $res