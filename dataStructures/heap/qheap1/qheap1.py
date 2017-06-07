
def merge(l, r):
    merged = []
    li, ri = 0, 0
    total = len(l) + len(r)
#    
#    
    for i in range(total):
        if li < len(l) and ri< len(r):
            if l[li] < r[ri]:
                merged.append(l[li])
                li = li+1
            elif r[ri] <= l[li]:
               
                merged.append( r[ri])
                ri = ri +1
        if li < len(l):
            merged.append(l[li])
            li = li+1
        if ri < len(r):
            merged.append(r[ri])
            ri = ri +1
    return merged
                

def sort(items):
    if len(items) <= 1:
        return items
    else:
        mid = len(items)/2
        
        l = sort(items[:mid])
        r = sort(items[mid:])
        return merge(l,r)
    

def insert(val, items):
    items.append(val)
    return sort(items)

def remove (val, items):
    search = items
    mid = len(search)/2
    
    while search[mid] != val:
        if val > search[mid]:
            search = search[mid:]
        elif val < search[mid]:
            search = search[:mid]
        mid = len(search)/2
        
    # we found mid. Shift all elemets to the right left by 1
    for i in range(len(items)-mid -1):
        items[mid + i] = items[mid+i+1]
    return items[:-1]

def display(items):
    return items[0]

def main():
    # we start off with taking an input
    # this is the number of queries
    # now we can define a handler for 1 2 3
    # for insert, place the element at the end of the items and merge sort
    # for delete, use binary search and evict it. followed by shift all cells to the right left
    # for 3, just return items[0]

    data = []
    count = input()
    printstack = []
    for i in range(count):
        inputs  = raw_input().split(' ')
        if inputs[0] == "1":
            data=insert(int(inputs[1]), data)
        elif inputs[0] == "2":
            data=remove(int(inputs[1]), data)
        elif inputs[0] == "3":
            printstack.append(display(data))

    for i in printstack:
        print(i)

if __name__ == "__main__":
    main()
