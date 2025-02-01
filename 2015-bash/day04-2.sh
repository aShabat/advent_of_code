#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

code=$(cat "$in")
suffix=0
while true; do
    hash=$(echo -n "$code$suffix" | md5sum)
    if [[ ${hash:0:6} == "000000" ]]; then
        echo "$suffix $hash"
        exit 0
    fi
    ((suffix++))
done
