#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

function pair {
    grep -qE '(..).*\1' <<<"$1"
}

function single {
    grep -qE '(.).\1' <<<"$1"
}

count=0
while read -r string; do
    pair "$string" && single "$string" && ((count++))
done <"$in"
echo $count
