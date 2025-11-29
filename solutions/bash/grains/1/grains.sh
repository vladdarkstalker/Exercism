#!/usr/bin/env bash

num="$1"

if [[ $num == "total" ]]; then
    echo "2^64-1" | bc
    exit 0
fi

if [[ $num -le 0 || $num -gt 64 ]]; then
    echo "Error: invalid input"
    exit 1
fi

i=1
res=1
while [ "$i" -lt $num ]; do
    res=$(echo "$res*2" | bc)
    i=$((i+1))
    done
echo "$res"