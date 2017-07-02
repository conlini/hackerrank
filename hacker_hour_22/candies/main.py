#!/bin/python

import sys

n = int(raw_input().strip())
T = map(int, raw_input().strip().split(' '))
V = map(int, raw_input().strip().split(' '))
# Write Your Code Here

crates = []
current = {}
for i in range(len(T)):
    if T[i] not in current:
        current[T[i]] = i
    else:
        crateSum = 0
        for k, v in enumerate(current):
            crateSum = crateSum|V[v]
        crates.append(crateSum)
        current = {}

sum = 0
for i in  crates:
    sum += i
print(sum)
