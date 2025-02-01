#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

function vowels {
    local count=
    count=$(grep -oE '[aeiou]' <<<"$1" | grep -c '')
    [[ $count -ge 3 ]] && return 0
    return 1
}

function repeat {
    grep -qE '(.)\1' <<<"$1"
}

function bad_strings {
    for bs in ab cd pq xy; do
        [[ "$1" =~ $bs ]] && return 1
    done
    return 0
}

count=0
while read -r string; do
    if vowels "$string" && repeat "$string" && bad_strings "$string"; then
        ((count++))
    fi
done <"$in"
echo "$count"
