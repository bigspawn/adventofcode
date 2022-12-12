#!/usr/bin/env bash

last_day="$(ls | sort -n | tail -1)"
next_day="$(printf "%02d" $((10#$last_day + 1)))"
next_day_1=$next_day/1
next_day_2=$next_day/2
mkdir -p "$next_day_1" "$next_day_2"
touch "$next_day/input.txt" "$next_day/input_small.txt"
cp main.go "$next_day_1/main.go"
cp main.go "$next_day_2/main.go"
