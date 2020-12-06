#!/bin/bash

: ${1?Need a value}

cd ~/src/aocurl
python3 aocurl.py 2020 $1

mkdir ~/src/advent-of-code/2020/day$1
cd ~/src/advent-of-code/2020/day$1

cp ../../template.go main.go
cp ~/.aocurl/aoc-2020-$1-input.txt input.txt

git add input.txt
git commit -m "Day $1 input"

git add main.go
git commit -m "Day $1 shell"
