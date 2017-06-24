#!/bin/python

import sys

def countCutsAndInserts(s):
    # Complete this function
    count = 0
    lenOrig = len(s)
    for i in range(1,lenOrig+1):
        for j in range(0, lenOrig+1):
            if j+i < lenOrig + 1:
                cut = s[j: j+i]
                sub = s[0:j] + s[j+i:]
                if cut != '':
                    for k in range(0, len(sub)+1):
                        insert = sub[0:k] + cut + sub[k:]
                        print(cut, sub, insert, i, j, k)
                        if insert == s:
                            count += 1
    return count
            

#  Return the number of successful ways to cut and insert a substring.
s = raw_input().strip()
result = countCutsAndInserts(s)
print(result)
