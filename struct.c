#include<stdio.h>

struct str
{
  int a,b;
  
};

typedef struct str str;


int main()
{
  
  str a={0,1};
  printf("%d%d",a.a,a.b);
  return 0;
}
