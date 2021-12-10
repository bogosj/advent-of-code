#!/bin/bash

: ${1?Need a value}

cd ~/src

if [ ! -d "/home/bogosj/src/aocurl" ]; then
  git clone https://github.com/IAmBullsaw/aocurl.git
fi

cd ~/src/aocurl
python3 aocurl.py 2021 $1

PAD=""
if (( $1 < 10 )); then
  PAD="0"
fi

mkdir ~/src/advent-of-code/2021/day$PAD$1
cd ~/src/advent-of-code/2021/day$PAD$1

cp ../../template.go main.go
cp ~/.aocurl/aoc-2021-$1-input.txt input.txt

git add input.txt
git commit -m "Day $1 input"

git add main.go
git commit -m "Day $1 shell"
