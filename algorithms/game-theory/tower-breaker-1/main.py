# Enter your code here. Read input from STDIN. Print output to STDOUT

def reduce(a):
    count = 0
    while a != 1:
        a = a >>1
        count += 1
    return count        


q = int(raw_input())
for i in range(q):
    n = int(raw_input())
    t = [reduce(int(w)) for w in raw_input().strip().split(' ')]
    print(t)
    print(sum(t))
    if sum(t)%2 == 0:
        print(2)
    else:
        print(1)
