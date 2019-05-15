#include <iostream>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include "signal.h"
#include "stdio.h"
#include "stdlib.h"
#include <vector>
#include <string.h>
#include <unistd.h>
using namespace std;

int main(int argc, char **argv) {

    if(argc<3){
        std::cout<< "----------命令说明-----------------------------------------"<<std::endl;
        std::cout<< "--NodeKiller 是用于在node.js服务实例和远程云端之间的桥梁------"<<std::endl;
        std::cout<< "--用法如下:------------------------------------------------"<<std::endl;
        std::cout<< "--NodeKiller  pid  command    ----------------------------"<<std::endl;
        std::cout<< "--command 指令服务器端--------------------------------------"<<std::endl;
        return 0;
    }
    string command;
    command=argv[2];

    //此处逻辑现在内存中开启共享内存，然后信号通知
    //
    int shmid;

    //创建共享内存
    shmid=shmget(0x2235, 64,IPC_CREAT | 0666);
    if(shmid==-1){
        std::cout<<"创建共享内存失败!！"<<std::endl;
        return 0;
    }


    void* p=shmat(shmid,NULL,0);
    if((void*)-1==(void*)p){
        std::cout<<"error"<<endl;
        exit(2);
    }
    char *commandList=NULL;
    commandList=(char*)p;

    strcpy(commandList,command.c_str());
    //通知进程
    kill(atoi(argv[1]),SIGUSR2);
    //休息俩S
    sleep(20);
    shmdt(p);

    //删除内存
    shmctl(shmid,IPC_RMID,NULL);


    return 0;
}