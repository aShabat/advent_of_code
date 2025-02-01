#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

floor=0
position=0
while read -r -a numbers; do
    for number in "${numbers[@]}"; do
        case $number in
        ")") ((floor--)) ;;
        "(") ((floor++)) ;;
        esac
        ((position++))
        if [[ $floor -lt 0 ]]; then
            echo "$position"
            exit 0
        fi
    done
done < <(sed -e 's/./\0 /g' <"$in")

echo "$floor"
