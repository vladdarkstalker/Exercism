#!/usr/bin/env bash

instr=${1,,}
siz=${#instr}
flag=0

for ((i=0; i<siz; i++)); do
    [[ ${instr:i:1} == " " || ${instr:i:1} == "-" ]] && continue
    for ((j=$((i+1)); j<siz; j++)); do
        [[ ${instr:j:1} == " " || ${instr:j:1} == "-" ]] && continue
        if [[ ${instr:j:1} != ${instr:i:1} ]]; then
            continue
        else
            echo "false"
            exit 0
        fi
    done
done

echo "true"
