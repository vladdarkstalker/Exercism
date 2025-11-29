#!/usr/bin/env bash

inp="$*"
trimmed_sp="$(echo "$inp" | tr -d '[:space:]')"
trimmed_all="$(echo "$inp" | tr -d '[:space:][:punct:][:digit:]')"

if [[ "$#" -eq 0 || -z "$trimmed_sp" ]]; then
    echo "Fine. Be that way!"
    exit
elif [[ "${trimmed_sp: -1}" == "?" && $trimmed_all =~ ^[A-Z]+$ ]]; then
    echo "Calm down, I know what I'm doing!"
    exit
elif [[ $trimmed_all =~ ^[A-Z]+$ ]]; then
    echo "Whoa, chill out!"
    exit
elif [[ "${trimmed_sp: -1}" == "?" ]]; then
    echo "Sure."
    exit
else
    echo "Whatever."
    exit
fi

