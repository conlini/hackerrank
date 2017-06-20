#!/bin/python

import sys

def minimumChocolateMoves(n, X):
    invalidEven, invalidOdd = 0,0
    oddsGt1 = 0
    evenGt2 = 0
    maxOdd, maxEven = 0,0
    for i in range(n):
        if i%2 == 0:
            if X[i]%2 == 0:
                if X[i] > 2:
                    evenGt2 += 1
                if X[i]>maxEven:
                    maxEven = X[i]
            else:
                invalidEven += 1
        elif i%2 == 1:
            if X[i]%2 == 1:
                if X[i] >1:
                    oddsGt1 += 1
                if X[i] > maxOdd:
                    maxOdd = X[i]
            else:
                    invalidOdd += 1
    invalid = invalidOdd + invalidEven
    if invalid == 0:
        return 0
    elif invalid%2 == 1:
        return -1
    else:
        if invalidOdd == invalidEven:
            return invalid/2
        count = 0
        if invalidOdd%2 == 0:
            if maxEven >= invalidEven +2:
                count += invalidOdd/2
            else:
                return -1
        if ivalidEven%2 == 0:
            if maxOdd >= invalidOdd +2:
                count += invalidEven/2
            else:
                return -1
        return count


    # Complete this function

#  Return the minimum number of chocolates that need to be moved, or -1 if it's impossible.
q = int(raw_input().strip())
for a0 in xrange(q):
    n = int(raw_input().strip())
    X = map(int, raw_input().strip().split(' '))
    result = minimumChocolateMoves(n, X)
    print(result)
