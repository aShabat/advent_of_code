#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

line=""
for _ in {0..999}; do
    line="${line}0"
done

lights=()
for _ in {0..999}; do
    lights+=("$line")
done

while read -r command; do
done <"$in"
