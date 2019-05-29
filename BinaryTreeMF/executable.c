#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
int main(){
  //pid_t pid=fork();
  //if(pid==0)
    {
      execl("/bin/gnome-terminal","gnome-terminal","-e","go run BinaryTree.go",(char*)NULL);
    }
  //else
    //{
      //system("ls");
    //}
            return 0;
//    f=fopen("log.txt","a+");
}
