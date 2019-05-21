#include<stdio.h>
#include<math.h>
#include<unistd.h>
int func(int p,int n,int lim)
{
  int v;
  
  v=pow(p,n);
  if(v<lim){

      return v;
  }
  else
    {
      printf("%d > %d\n",v,lim);
    }
  return lim;
}

int main()
{int a[2][3]={{1,2,3},{3,4,5}};
  printf("%d %d",func(3,2,10),func(3,3,20));//Could be that C uses multithreading in printf statement
  return 0;
}
