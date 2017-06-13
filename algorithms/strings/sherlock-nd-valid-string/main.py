#!/bin/python

import sys
from collections import defaultdict

def isValid(s):
    counter = {}
    for i in s:
        if i in counter:
            counter[i] = counter[i]+1
        else:
            counter[i] = 1
    currentFreq = 0
    numberOfDevs = 0
    print(counter)
    for _,v in counter.items():
        if currentFreq > 0:
            if v != currentFreq:
                if abs(v-currentFreq) > 1 and numberOfDevs >1:
                    return "NO"
                else:
                    numberOfDevs = 1
            elif v > currentFreq:
                currentFreq = v
        else:
            currentFreq = v
    return "YES"

s = raw_input().strip()
result = isValid(s)
print(result)

