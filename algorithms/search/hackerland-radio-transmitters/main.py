n,k = raw_input().strip().split(' ')
n,k = [int(n),int(k)]
x = map(int,raw_input().strip().split(' '))


a = sorted(x)
#print(a)
pos = 0
i = 0
count = 0
while i < n:
    next = a[i] + k
 #   print(next)
    while a[i] <= next:
        i += 1
        if i > n-1:
            break
    i -= 1
  #  print(i)
    count += 1
    right = a[i] + k
   # print(right)
    if right > a[-1]:
        break
    while a[i] <= right:
        i += 1
        if i >n-1:
            break
   # print(i)
   # print("\n")
    
print(count)

        
        
