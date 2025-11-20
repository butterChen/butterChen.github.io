[[Python]]
[[aiot徐昌旭]]


![[從這裡開啟py_shell.png]]
class
```
class bmi:
    def __init__(self,name, kg, cm):
        self.name = name
        self.kg = kg
        self.m = float(cm/100)
        self.status = self.kg/( self.m**2)
#         self.status = format(self.kg/( self.m**2), '.2f')
        self.sstatus = format(self.kg/( self.m**2), '.2f')
         
    def cal(data):
#         David's BMI is :19.87
        val = data.kg/( data.m**2)
        val = format(val, '.2f')
        print(data.name+"'s BMI is :"+ str(data.sstatus) )
        
#         print(data.name+"'s BMI is :"+ str(data.cal2()) )
        
    def cal2(self):
        return self.kg/( self.m**2)
    
    def cal3(self):
        if self.status <18.5: 
            print("過輕")
        elif 18.5<=self.status <24: 
            print("正常")
        elif 24<=self.status <27: 
            print("過重")
        elif 27<=self.status : 
            print("肥胖")
```
main
```
from b_class import bmi
def main():   #以函式呼叫
    student1 = bmi("David", 68,185)
    student1.cal()
    student1.cal3()
    student2 = bmi("Kevin", 53,172)
    student2.cal()
    student2.cal3()
    print("plz")

main()
```
