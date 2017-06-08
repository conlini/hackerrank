n,k = raw_input().strip().split(' ')
n,k = [int(n),int(k)]
x = map(int,raw_input().strip().split(' '))


a = sorted(x)
print(a)
pos = 0
bases = []

visited = []
count = 0
base =False



for current,i in enumerate(a):
    if i in visited:
        continue
    else:
        visited.append(i)

        if i - a[pos] > k:
            if not base:

                pos = current -1
                bases.append(a[pos])
                count += 1
                base = True

            else:
                pos = current
                base = False
        print(i, a[pos], i-a[pos], pos, count)
if not base: 
    bases.append(a[-1])
    count += 1

print(count)
print(bases)
        
        
