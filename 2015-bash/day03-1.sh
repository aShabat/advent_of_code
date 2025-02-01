#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

x=0
y=0
houses="0 0"
nl=$'\n'
while read -r -a moves; do
    for move in "${moves[@]}"; do
        case $move in
        '>') ((x++)) ;;
        '<') ((x--)) ;;
        '^') ((y++)) ;;
        'v') ((y--)) ;;
        esac
        houses="$houses${nl}$x $y"
    done
done < <(sed -e "s/./\0 /g" <"$in")
houses=$(echo "$houses" | sort | uniq)
grep -c "^.*$" <<<"$houses"
