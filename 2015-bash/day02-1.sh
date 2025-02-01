#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

paper=0
while read -r dimensions; do
    dimensions=($dimensions)
    IFS=$'\n' dimensions=($(sort -n <<<"${dimensions[*]}"))
    unset IFS

    a=${dimensions[0]}
    b=${dimensions[1]}
    c=${dimensions[2]}
    ((paper += 3 * a * b))
    ((paper += 2 * a * c))
    ((paper += 2 * b * c))
done < <(sed "$in" -e "s/x/\ /g")

echo "$paper"
