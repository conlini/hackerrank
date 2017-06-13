def mapify(s):

    counts = {}
    for i in range(len(s)):

        if i < len(s) -1:
            if s[i] == s[i+1] and s[i] in counts:
                # if two chars are back to back, dont consider them
                del counts[s[i]]
            else:
                if s[i]  in counts:
                    counts[s[i]] = counts[s[i]] +1
                else:
                    counts[s[i]] = 1
        else:
            if s[i]  in counts:
                counts[s[i]] = counts[s[i]] +1
            else:
                counts[s[i]] = 1
    sorted_count = sorted(counts.items(), key=lambda x:x[1], reverse=True)
    print(sorted_count[0][1] + sorted_count[1][1])



s_len = int(raw_input().strip())
s = raw_input().strip()
mapify(s)
