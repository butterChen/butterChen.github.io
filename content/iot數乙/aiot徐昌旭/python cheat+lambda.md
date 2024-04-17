
[[aiot蔡佳喻]]
[[aiot徐昌旭]]

a = [int(s) for s in input().split()]  
a = [[int(j) for j in input().split()] for i in range(n)]  
a = [['.'] * n for i in range(n)]  
map 與 lambda組合  
用list轉換  
------------------------------------------------  
------------------------------------------------------------------------------  
import math  
score = [30, 38, 89, 22, 43]  
print([math.ceil(i**0.5*10) for i in score])  
print(list(map(lambda x:math.ceil(x**0.5*10),score)))  
------------------------------------------------------------------------------  
  
import math  
score = [30, 38, 89, 22, 43]  
c = list(map(lambda x: math.ceil(math.sqrt(x)*10),score))  
print(c)  
------------------------------------------------------------------------------  
positional arguments  
keyword arguments