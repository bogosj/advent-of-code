#!/bin/sh

cd /tmp
wget -q https://raw.githubusercontent.com/AidanGlickman/Advent-2020/4cc3efc1a5a6a6e7ab0c0bc26b4d5ce60cfa77e0/day20/solution.py
cp /work/input.txt /tmp
python3 solution.py
