def computeLps(pat, M):
    answer = [0]*M
    answer[0] = 0
    i = 1
    len = 0
    while i < M:
        if pat[len] == pat[i]:
            len += 1
            answer[i] = len
            i += 1
            print("match i={}, j={}", i, len)
        else:
            if len != 0:
                len = answer[len-1]
            else:
                answer[i]= 0
                i += 1
    return answer


def kvmsearch(pat, txt):
    i, j = 0,0
    lps = computeLps(pat, len(pat))
    print(lps)
    while i < len(txt):

        if txt[i] == pat[j]:
#            print("match - i={}, j={}".format(i,j))
            i += 1
            j +=1

            if j == len(pat):
                # match. reset j
                print("match at index{}".format(i-j))
                j = lps[j-1]
        else:
            if j > 0:
                j = lps[j-1]
            else:
                i += 1
                

txt = raw_input()
pat = raw_input()
#kvmsearch(pat, txt)
print(computeLps(pat, len(pat)))
