#!/bin/bash
in="${1:-${0%-[0-9].*}.input}"
[[ -e "$in" ]] || exit 1

ribbon=0
while read -r dimensions; do
    dimensions="${dimensions//\ /$'\n'}"
    dimensions=$(sort -n <<<"$dimensions")
    echo "$dimensions" >/tmp/dimensions
    dimensions=()
    while read -r dim; do
        dimensions+=("$dim")
    done <"/tmp/dimensions"

    a=${dimensions[0]}
    b=${dimensions[1]}
    c=${dimensions[2]}
    ((ribbon += a + a + b + b + a * b * c))
done < <(sed "$in" -e "s/x/\ /g")

echo "$ribbon"
