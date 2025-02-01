#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

floor=0
while read -r -a numbers; do
    for number in "${numbers[@]}"; do
        case $number in
        ")") ((floor--)) ;;
        "(") ((floor++)) ;;
        esac
    done
done < <(sed -e 's/./\0 /g' <"$in")

echo "$floor"
