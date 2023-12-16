#!/bin/bash

: ${1?Need a value}

cd ~/aocurl
python3 aocurl.py 2022 $1

PAD=""
if (( $1 < 10 )); then
  PAD="0"
fi

mkdir /workspaces/advent-of-code/2022/day$PAD$1
cd /workspaces/advent-of-code/2022/day$PAD$1

cp ../../template.go main.go
cp ~/.aocurl/aoc-2022-$1-input.txt input.txt

git add input.txt
git commit -m "Day $1 input"

git add main.go
git commit -m "Day $1 shell"
