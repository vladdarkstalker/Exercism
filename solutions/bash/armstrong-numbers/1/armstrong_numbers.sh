#!/usr/bin/env bash

arg="$1"
len=${#arg}
arms=0
for (( i=0; i<len; i++ )); do
    el="${arg:$i:1}"
    arms=$(($arms+$el**$len))
done

if [ "$arg" -eq "$arms" ]; then
    echo "true"
else
    echo "false"
fi