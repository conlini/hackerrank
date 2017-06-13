

q = int(raw_input().strip())
answers = []
for a0 in xrange(q):
    s = raw_input().strip()
    j = 0
    found = True
    for i in "hackerrank":
#        print("looking for {} from pos {}".format(i,j))
        if j== len(s) and i < len("hackerrank"):
            answers.append("NO")
            found=False
            break
        for pos, x in enumerate(s[j:]):
            if x == i:
                j += pos +1
                break
        else:
            answers.append("NO")
            found=False
            break
    if found:
        answers.append("YES")

for i in answers:
    print(i)
            
