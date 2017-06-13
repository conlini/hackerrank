def lonely(arr):
    one, two = 0,0
    common = 0
    for i in arr:
        print("number {}".format(i))
        two = two | (one & i)
        print( one, two)
        one = one ^ i
        print( one)
        common = ~(one & two)
        print(common, one, two)
        one &= common
        print(common, one)
        two &= common
        print(common, two)
#        print(i, one, two, common)
    return one


print(lonely([5,3,5,2,5,2,2]))
