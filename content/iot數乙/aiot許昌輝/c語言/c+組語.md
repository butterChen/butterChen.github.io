[[C語言]]
http://ccckmit.wikidot.com/as:inlinec

https://hackmd.io/@sysprog/c-bitwise

,,,
,,,
'''
```
main()
{
	int i;
	int a[]={ 10, 20, 30, 40, 50, 60};
	int *p, *q;

	p=a;
	q=a+2;
	for(i=0; i<6; i++)
	{
		printf("p[%d]=%d\t\t", i, p[i]);
		printf("q[%d]=%d\n", i-2, q[i-2]);
	}
	 int foo = 10, bar = 15;
	 asm(                           
    "addl %%ebx,%%eax"           
    :"=a"(foo)                   
    :"a"(foo), "b"(bar)          
  );                
  printf("foo=%d\n", foo);
}
```
