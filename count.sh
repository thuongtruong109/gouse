#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Please provide at least one path."
  exit 1
fi

total_count=0

# method 1
# for dir in "$@"; do
#   for file in $(find "$dir" -type f -name "*.go" ! -name "*_test.go"); do
#     count=$(grep -E 'func [A-Z]' "$file" | grep -vE 'func New' | wc -l)
#     total_count=$((total_count + count))
#   done
# done

# method 2
for dir in "$@"; do
  files=$(find "$dir" -type f -name "*.go" ! -name "*_test.go")
  if [ -z "$files" ]; then
    continue
  else
    :
  fi

  while IFS= read -r file; do
    count=$(grep -Ec 'func [A-Za-z]' "$file" | grep -vE 'func New')
    total_count=$((total_count + count))
  done <<< "$files"
done

echo "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"85\" height=\"20\" role=\"img\" aria-label=\"$total_count functions\" class=\"mdl-js\">
  <title>$total_count functions</title>
  <linearGradient id=\"s\" x2=\"0\" y2=\"100%\">
    <stop offset=\"0\" stop-color=\"#bbb\" stop-opacity=\".1\"/>
    <stop offset=\"1\" stop-opacity=\".1\"/>
  </linearGradient>
  <clipPath id=\"r\">
    <rect width=\"85\" height=\"20\" rx=\"3\" fill=\"#fff\"/>
  </clipPath>
  <g clip-path=\"url(#r)\">
    <rect width=\"0\" height=\"20\" fill=\"#13708a\"/>
    <rect x=\"0\" width=\"85\" height=\"20\" fill=\"#13708a\"/>
    <rect width=\"85\" height=\"20\" fill=\"url(#s)\"/>
  </g>
  <g fill=\"#fff\" text-anchor=\"middle\" font-family=\"Verdana,Geneva,DejaVu Sans,sans-serif\" text-rendering=\"geometricPrecision\" font-size=\"110\">
    <text aria-hidden=\"true\" x=\"425\" y=\"150\" fill=\"#010101\" fill-opacity=\".3\" transform=\"scale(.1)\" textLength=\"750\">$total_count functions</text>
    <text x=\"425\" y=\"140\" transform=\"scale(.1)\" fill=\"#fff\" textLength=\"750\">$total_count functions</text>
  </g>
</svg>" > public/count.svg
