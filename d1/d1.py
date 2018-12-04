# Advent of Code 2018
# Day 1: Chronal Calibration
from itertools import accumulate, cycle

with open('input.txt') as file: 
    # Part 1: 
    frequencies = [int(line) for line in file.read().splitlines()] # list comprehension
    print(sum(frequencies))   
    # Part 2: 
    seen = set()
    for x in accumulate(cycle(frequencies)):
        if x in seen:
            break
        else:
            seen.add(x)
    print(x)