#!/usr/bin/python2.7/python
# coding=utf-8
'''
n个苹果，m个框，要求将苹果分配到框里，客户来买任意数量的苹果都可以不必重新分配
'''
def anyAllocation(n, m):

    if m <= 0 or n <= 0:
        return []

    num_ready = {}
    result = []
    sum_now = 0

    for i in range(1,n+1):
        if( num_ready.get(i) ): #已经满足顾客需求，无需分配
            continue

        for k,v in num_ready.items(): #不满足顾客需求，新分配到框里
            num_ready[k+i] = True

        num_ready[i] = True
        sum_now = sum_now + i

        if sum_now > n:
            i = i - ( sum_now - n )

        result.append(i)

        if len( result ) == m and sum_now < n:
            print "m太小"
            return []

        if sum_now >= n:
            return result

print anyAllocation(77, 5)

