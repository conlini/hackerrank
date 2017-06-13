def search(c, n,m, outs):


    for i in range(0, n-1):

        for j in range(i+1, n):

            if c[i] + c[j] == m:
                outs.append([i+1,j+1])
                return


t = input()
outs = []
for i in range(t):
    m = input()
    n = input()
    c = [int(x) for x in raw_input().strip().split(' ')]

    search(c, n,m, outs)

for i in outs:
    print("{} {}".format(i[0], i[1]))
        

