#!/usr/bin/env bash

if (( $(echo "$2 <= 0" | bc -l) == 1 )) || \
   (( $(echo "$3 <= 0" | bc -l) == 1 )) || \
   (( $(echo "$4 <= 0" | bc -l) == 1 )); then
    echo "false"
    exit 0
fi

if (( $(echo "$2 + $3 <= $4" | bc -l) == 1 )) || \
   (( $(echo "$4 + $3 <= $2" | bc -l) == 1 )) || \
   (( $(echo "$2 + $4 <= $3" | bc -l) == 1 )); then
    echo "false"
    exit 0
fi

if [[ "$1" == "equilateral" ]]; then
    if (( $(echo "$2 == $3" | bc -l) == 1 )) && \
       (( $(echo "$3 == $4" | bc -l) == 1 )) && \
       (( $(echo "$4 == $2" | bc -l) == 1 )); then
        echo "true"
    else
        echo "false"
    fi
elif [[ "$1" == "isosceles" ]]; then
    if (( $(echo "$2 == $3" | bc -l) == 1 )) || \
       (( $(echo "$3 == $4" | bc -l) == 1 )) || \
       (( $(echo "$4 == $2" | bc -l) == 1 )); then
        echo "true"
    else
        echo "false"
    fi
elif [[ "$1" == "scalene" ]]; then
    if (( $(echo "$2 != $3" | bc -l) == 1 )) && \
       (( $(echo "$3 != $4" | bc -l) == 1 )) && \
       (( $(echo "$4 != $2" | bc -l) == 1 )); then
        echo "true"
    else
        echo "false"
    fi
else
    echo "false"
fi