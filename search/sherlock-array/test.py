#!/bin/python

import sys

def solve(a):
    # Complete this function
    if len(a) == 1:
        return "YES"
    if len(a) > 2:
        total = sum(a)  
        left = 0
        right = total
        print(total)
        for i in range(1, len(a)):
        
            left += a[i-1]
            right -= a[i] + left
            print(a[i], left, right)
            if left == right:
                return "YES"
    return "NO"

T = int(raw_input().strip())
for a0 in xrange(T):
    n = int(raw_input().strip())
    a = map(int, raw_input().strip().split(' '))
    result = solve(a)
    print(result)


