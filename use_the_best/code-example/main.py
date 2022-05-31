#!/bin/python3

a = [1,2,3,4,5,6]
b = [5,5,7,3,5,7,3]
x = {i for i  in a}
y = {i for i in b}
#result = y.union(x)
result = a+b
print(sorted(result))
