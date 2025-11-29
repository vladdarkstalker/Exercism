#!/usr/bin/env bash

res=""
input=$(echo "${1//-/ }" | tr -cd '[:alpha:] [:space:]')
for word in $input; do
res+=${word:0:1}
done
res=${res^^}
echo $res