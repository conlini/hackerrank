def longestIncreasingSubSequence(arr):
    lis = [1]*len(arr)
    max = 0
    for i in range(1, len(arr)):
        for j in range(0, i):
            #print("comparing {} with {} for lis position {}. with lis {} and {}".format(arr[i], arr[j], i, lis[i], lis[j]))
            if arr[i] > arr[j] and lis[i] < lis[j]+1:
                lis[i]= lis[j] + 1
                if lis[i] > max:
                    max = lis[i]
              #      print("new max {}".format(max))
    print(lis)
    return max


arr = [int(x) for x in raw_input().strip().split()]
print(longestIncreasingSubSequence(arr))
                   
    
