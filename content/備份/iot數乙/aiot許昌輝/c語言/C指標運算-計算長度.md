[[C語言]]
[[aiot許昌輝]]
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