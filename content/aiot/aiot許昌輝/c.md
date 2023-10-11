[[aiot許昌輝]]

[[C語言]]
==指標與一維陣列==

```c
#include <stdio.h>

main()

{  int a[]={ 10, 20, 30, 40, 50, 60};

  int *p;

  p=a;

  *p+=2;

  printf("*p=%d\n", *p);

  *p++;

  printf("*p=%d\n", *p);

  p++;

  printf("*p=%d\n", *p);

  ++*p;

  printf("*p=%d\n", *p);

}
```
==指標與二維陣列==
```
A:

int a[]={10,20,30,40,50,60};  為一維陣列 a

如欲取得ａ陣列的第二個元素可使用 a[1] 或 *(a+1) 兩種方式

B:

int a[3] [4] ; 為二維陣列 a
```
![[Pasted image 20231004144800.png]]

```c
#include <stdio.h>

main()

{  int i;

  int a[3][4]={ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12};

  int *p;

  p=a;

  for(i=0; i<12; i++)

  {  printf("p+%d=%x\t*(p+%d)=%d\t", i, p+i, i, *(p+i));

      printf("&a[%d][%d]=%x\n", i/4, i%4, &a[i/4][i%4]);

  }

}
```

```c
#include <stdio.h>

main()

{  int i, j;

  int a[3][4];

  for(i=0; i<3; i++)

      { for(j=0; j<4; j++)

  printf("&a[%d][%d]=%x\t", i, j, &a[i][j]);

  printf("\n");

     }

  for(i=0; i<3;i++)

     { printf("*(a+%d)=%x\n", i,*(a+i));

       for(j=0; j<4; j++)

       printf("*(a+%d)+%d=%x\t", i,j,*(a+i)+j);

      puts("");

    }

}
```
==指標與字串==
```
字串是字元陣列，宣告字串可以使用三種方式﹕

1.  char  str1[10]={ ‘H’ ,‘e’ ,‘l’ ,‘l’,’o’,’!’,’\0’};

2.  char  str2[]=“ C is powerful!“;

3.  char  *strp = “ C is powerful!“;
```
a. 保留4個 Bytes 記憶体設定給指標 strp 使用。

b. 將指標 strp 指向字串常數。
```c
#include <stdio.h>

void main()

{

  char *strp="C is powerful!";

  printf("%s\n",strp);

  printf(strp+5);

}
```
![[Pasted image 20231004144938.png]]
```c
#include <stdio.h>

void main()

{  char *s1="fun";

  char *s2="fun";

  char *s3="fun";

  printf("s1=%x\ts2=%x\ts3=%x\n", s1, s2, s3);

  printf("Input a string:");

  scanf("%s", s1);

  s3="funny";

  printf("s1=%x\ts2=%x\ts3=%x\n", s1, s2, s3);

  printf("s1=%s\n", s1);

  printf("s2=%s\n", s2);

  printf("s3=%s\n", s3);

}
```

```c
#include <stdio.h>

#include <string.h>

void main()

{

  char *str="C is fun!";

  char *p;

  p=str+strlen(str);

  while(--p>=str)

  putchar(*p);

}
```
![[Pasted image 20231004145010.png]]

==指標與函式-反向輸出字串==

```c
#include <stdio.h>

void flip(char*);

void main()

{

  char str[80];

  printf("Input a string:");

  gets(str);

  flip(str);

  printf(str);

}
void flip(char* s)

{

  char d[80];

  int i, len;

  char *p;

  len=strlen(s);

  p=d+len-1;

  for(i=0; i<len; i++)

  *(p-i)=*(s+i);

  d[len]='\0';

  strcpy(s, d);

}
```
![[Pasted image 20231004145041.png]]

==指標與函式-大寫轉小寫==
```c
#include <stdio.h>

void str2l(char*);

void char2l(char*);

void main()

{

  char str[80];

  printf("Input a string:");

  gets(str);

  str2l(str);

  printf(str);

}
void str2l(char* s)

{

  int i, len;

  len=strlen(s);

  for(i=0; i<len; i++)

  char2l(s+i);

}

void char2l(char* c)

{

  if(*c >='A' && *c<='Z')

  *c+=32;

}
```

==回傳指標-反向輸出字串==

```c
#include <stdio.h>

char* flip(char*);

void main()

{

  char str[80];

  char *afterflip;

  printf("Input a string:");

  gets(str);

  afterflip=flip(str);

  printf("str=\"%s\"\n", str);

  printf("afterflip=\"%s\"", afterflip);

}
char* flip(char* s)

{

  static char d[80];

  int i, len;

  char *p;

  len=strlen(s);

  p=d+len-1;

  for(i=0; i<len; i++)

  *(p-i)=*(s+i);

  d[len]='\0';

  return d;

}
```
==指標運算==
```c
void main()

{  char str[80];

  char *afterflip;

  printf("Input a string:");

  gets(str);

  printf("%d\n", slen(str));

}

int slen(char* s)

{

  char* p=s-1;

  while(*++p);

  return p-s;

}
```

```c
void main()

{  char str[80];

  char *afterflip;

  printf("Input a string:");

  gets(str);

  printf("%d\n", slen(str));

}

int slen(char* s)

{  char* p=s-1;

  while(*++p);

  return p-s;

}
```

```c
#include <stdio.h>

void main()

{

  char *p;

  char *q=0;

  p=(char*)0x22fe40;

  printf("p=%d\n", *p);

  printf("q=%s", q);

}
```

==深入指標-指標陣列==

![[Pasted image 20231004145321.png]]

```c
#include <stdio.h>

void main()

{  int i, j;

  char n[6][15] = { "Michael", "Bill", "Carlson",

              "Stephanie", "Joe", "Mary"};

  char *m[6] = {"Michael", "Bill", "Carlson",

              "Stephanie", "Joe", "Mary"};

  for(i=0; i<5; i++)

  {

    printf("n[%d]-n[%d]=%d\t", i+1, i, n[i+1]-n[i]);

    printf("m[%d]-m[%d]=%d\t", i+1, i, m[i+1]-m[i]);

    printf("%s\n", m[i]);

  }

}
```

```
執行結果

n[1]-n[0]=15      m[1]-m[0]=8        Michael

n[2]-n[1]=15      m[2]-m[1]=5        Bill

n[3]-n[2]=15      m[3]-m[2]=8        Carlson

n[4]-n[3]=15      m[4]-m[4]=10       Stephanie

n[5]-n[4]=15      m[5]-m[4]=4        Joe
```

==深入指標==
```c

#define SIZE 6

void main()

{

  char *m[SIZE] = {"Michael", "Bill", "Carlson",

  "Stephanie", "Joe", "Mary"};

  int i, j;

  char *temp;

  char *ori;

  ori=m[0];
  for(i=0; i<SIZE-1; i++)

  for(j=0; j<SIZE-i-1 ; j++)

  if( strcmp(*(m+j), *(m+j+1))>0 )

      {  temp=m[j];

  m[j]=m[j+1];

  m[j+1]=temp;}

  puts(">>----------------After Sorting:");

  for(i=0; i<SIZE; i++)

        puts(m[i]);

  puts(">>----------------Strings in memory:");

  for(i=0; i<40; i++)

        if(*(ori+i)=='\0')

  putchar(' ');

        else

  putchar(*(ori+i));

}
```

```
執行結果:

>>..after sorting

Bill

Carlson

Joe

Mary

Michael

Stephanie

>>..String in memory

Michael Bill Carlson

Stephanie Joe Mary
```
![[Pasted image 20231004152049.png]]

==指向指標的指標==
20231004
[[2023-10-04]]

![[Pasted image 20231004152120.png]]

```c
#include <stdio.h>

void main()

{

  int i, *p, **pp;

  i=100;

  p=&i;

  pp=&p;

  printf("i  address:%x\t value:%d\n", &i, i);

  printf("p  address:%x\t value:%x\n", &p, p);

  printf("pp address:%x\t value:%x\n", &pp, pp);

  printf("*&pp =====>%x, **&pp =====>%x, ***&pp =====>%d\n",*&pp,**&pp,***&pp);

  printf("*pp ======>%x, **pp ======>%d\n", *pp, **pp);

  }}
  
```

```
執行結果:

i      address:22fe4c   Value:100

p     address:22fe40   Value:22fe4c

pp   address:22fe38   Value:22fe40

*pp=22fe38

**pp=100
```
==這裡很特別 內層+1是移動一列==
另外一個想法是 
```
m[0] 轉換成 **(m+0)
m[1] 轉換成 **(m+1)
m[0][0] 轉換成 *(*(m+0)+0)
m[1][0] 轉換成 *(*(m+1)+0)
```
[[2023-10-04]]
```c
*(*(m+1)+0)=B  
**(m+1)=B
```


```c
atoi()
[

### atoi()

](http://tw.gitbook.net/c_standard_library/c_function_atoi.html)
```
![[Pasted image 20231004152203.png]]

```c
#include <stdio.h>

void main()

{  int i, j;

  char *m[6] = {"Michael", "Bill", "Carlson",

       "Stephanie", "Joe", "Mary"};

  char c;

  for(i=0; i<2; i++)

  {     for(j=0; j<15; j++)

    {   c=*(*(m+i)+j);

         if(c=='\0')

              break;

         else

     printf("*(*(m+%d)+%d)=%c\t", i, j, c);

    }

  printf("\n");

  }

}
```

```
執行結果:

*(*(m+0)+0)=M   *(*(m+0)+1)=i     *(*(m+0)+2)=c   *(*(m+0)+3)=h

*(*(m+0)+4)=a    *(*(m+0)+5)=e    *(*(m+0)+6)=l

*(*(m+1)+0)=B    *(*(m+1)+1)=i     *(*(m+1)+2)=l   *(*(m+1)+3)=l
```

==程式參數==
```c
void main(int argc, char *argcv[])
```


![[Pasted image 20231004161807.png]]

```c
#include <stdio.h>

void main(int argc, char *argv[])

{

  if(argc<2 || argc>3)

  {

  printf("Using:hi [-option] YOURNAME\n");

  printf("option:1~3");

  }

  else

  if(*argv[1]!='-')

  printf("Good day, %s!", argv[1]);
  else  {

  switch(*(argv[1]+1)){

  case '1':

  printf("Good morning, ");  break;

  case '2':

  printf("Good afternoon, ");

  break;

  case '3':

  printf("Good night, ");

  break;

  default:

  printf("Good day, ");

  }

  printf("%s!", argv[2]);

  }

}
```
```c
//-2的2
//可在 `main` 函數中加上 `argc` 與 `argv` 兩個參數：
//argc 為參數數量
*(argv[1]+1)

```

測試結果 [[2023-10-11]]
![[Pasted image 20231011085328.png]]
==函數的位址==

C 的基本型別變數和陣列都有位址，因此都可以在 int、char、 float、 array上使用指標。

C 的函數也有位址，也可用指標來間接定位，或使用指標指明位址呼叫。

函數名稱加上() ， C 會將它譯成函數，函數名稱本身 C 會將它譯成函數位址， 如下列：

   int func();    func() 為函數， func 為函數位址。

==函數的指標==

```c
#include <stdio.h>

void main()

{

  int func();

  printf("func's adress:%x\n", func);

  printf("func's return value:%u\n", func());

}

func()

{

  return 123;

}
```
```
執行結果:

func’s address:XXX

func’s return value:123
```

==指向函數的指標==
```c
int fp;  fp 為某個整數值。

int (*fp)(); fp 為指向某個整數函數的位址。

int *fp(); fp()為一個函數，表示會傳回一個指向整數的指標
```
```c
#include <stdio.h>

void main()

{

  int func();

  int (*fp)();

  fp=func;

  printf("func's return value:%d\n", func());

  printf("Using functional pointer:%d", (*fp)());

}

func()

{

  return 123;

}
```
```
執行結果:

func’s return value:123

Using functional pointer:123
```

```c
void main()

{  char *func(int);

  char *(*fp)(int);

  fp=func;

  printf("%s\n", func(1));

  printf("%s", (*fp)(0));

}

char *func(int i)

{  if(i)

       return "C is fun!";

  else

      return "Have a good time!";

}
```
```
執行結果:

C is fun!

Have a good time
```

```c
#include <stdio.h>

void main(int argc, char *argv[])

{

  int add(int, int);

  int multiply(int, int);

  int subtract(int, int);

  int bad(int, int);

  int (*fp)(int, int);

  int i=1;
  while(i<argc)

  {

  if(!strcmp("+", argv[i]))

  fp=add;

  else if(!strcmp("*", argv[i]))

  fp=multiply;

  else if(!strcmp("-", argv[i]))

  fp=subtract;

  else

  fp=bad;

  (*fp) (atoi(argv[i+1]), atoi(argv[i+2]));

  i+=3;

  }

}
add(int a, int b)

{  printf("%d\n", a+b);

}

multiply(int a, int b)

{  printf("%d\n", a*b);

}

subtract(int a, int b)

{  printf("%d\n", a-b);

}

bad(int a, int b)

{  printf("No this option!\n");

}
```