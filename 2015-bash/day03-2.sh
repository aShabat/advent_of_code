#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

x=(0 0)
y=(0 0)
turn=0
houses="0 0"
nl=$'\n'
while read -r -a moves; do
    for move in "${moves[@]}"; do
        ((turn++))
        ((turn %= 2))
        case $move in
        '>') ((x[turn]++)) ;;
        '<') ((x[turn]--)) ;;
        '^') ((y[turn]++)) ;;
        'v') ((y[turn]--)) ;;
        esac
        houses="$houses${nl}${x[$turn]} ${y[$turn]}"
    done
done < <(sed -e "s/./\0 /g" <"$in")
houses=$(echo "$houses" | sort | uniq)
grep "^.*$" -c <<<"$houses"
