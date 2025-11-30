#!/usr/bin/env bash

res=""
isnum=1
if [ $(($1 % 3)) -eq 0 ]; then
    res+="Pling"
    isnum=0
fi
if [ $(($1 % 5)) -eq 0 ]; then
    res+="Plang"
    isnum=0
fi
if [ $(($1 % 7)) -eq 0 ]; then
    res+="Plong"
    isnum=0
fi
if [ $isnum -eq 1 ]; then
    res+=$1
fi
echo $res