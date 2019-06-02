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
#include <dirent.h>
#include <libgen.h>
#include <fcntl.h>
#include <mutex>
#include <boost/lambda/lambda.hpp>
#include <iostream>
#include <iterator>
#include <algorithm>
#include <boost/program_options.hpp>
#include <boost/program_options/errors.hpp>


using namespace std;
using namespace boost;

using namespace boost::program_options;
std::mutex mtx;




int main(int argc, char **argv) {
    mtx.lock();


    try {
        options_description desc("平安科技版权所有   \n"
                                 "Usage: NodeKiller [-options]  [args...]\n"
                                 " e.g NodeKiller --help  \n"
                                 "     \n");
        desc.add_options()
                ("help", "帮助命令")
                ("cpuprofile", value<int>(), "NodeKiller --cpuprofile PID;向指定的Node.js发送cpuprofile 信号，它将产生3分钟那该进程的CPU分配详细信息,以'Node-进程-时间戳.cpuprofile'于服务器指定位置 \n")
                ("heapprofile",value<int>(),"NodeKiller --heapprofile PID;向指定的Node.js发送heapprofile 信号，它将产生3分钟那该进程的堆内存分配详细信息,以'Node-进程-时间戳.heapprofile'于服务器指定位置 \n")
                ("heapdump",value<int>(),"NodeKiller --heapdump PID;向指定的Node.js发送heapdump 信号，它将产生这一刹那该进程的堆内存分配详细信息,以'Node-进程-时间戳.heapsnapshot'于服务器指定位置 \n")
                ("timeline",value<int>(),"NodeKiller --timeline PID;向指定的Node.js发送timeline 信号，它将产生3分钟之内该进程的堆内存分配详细信息,以'Node-进程-时间戳.heaptimeline'于服务器指定位置 \n")
                ("gctrace",value<int>(),"NodeKiller --gctrace PID;向指定的Node.js发送gctrace 信号，它将产生3分钟之内该进程的GC详细信息,以'Node-进程-gc-tracer-时间戳.log'于指定位置 \n")
                ("forceGC",value<int>(),"NodeKiller --forceGC PID;向指定的Node.js进程发送强制GC信号。")
                ;

        variables_map vm;
        store(parse_command_line(argc, argv, desc), vm);
        notify(vm);

        if (vm.count("help")) {
            cout << desc << "\n";
            return 1;
        }
        std::string command="";
        int pid=-1;
        if (vm.count("cpuprofile")) {

            pid=vm["cpuprofile"].as<int>();
            command="cpuprofile";

        } else if(vm.count("heapprofile")) {
            pid=vm["heapprofile"].as<int>();
            command="heapprofile";

        }else if(vm.count("heapdump")){
            pid=vm["heapdump"].as<int>();
            command="heapdump";
        }else if(vm.count("timeline")){
            pid=vm["timeline"].as<int>();
            command="timeline";
        }else if(vm.count("gctrace")){
            pid=vm["gctrace"].as<int>();
            command="gctrace";
        }else if(vm.count("forceGC")){
            pid=vm["forceGC"].as<int>();
            command="forceGC";
        }
        else{
            cout<< "没有指定命令, 退出进程；" <<endl;
            exit(1);
        }

        if(pid==-1){
            exit(1);
        }
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
        kill(pid,SIGUSR2);
        //休息俩S
        sleep(20);
        shmdt(p);

        //删除内存
        shmctl(shmid,IPC_RMID,NULL);

        mtx.unlock();
    }catch (std::exception& e){
        std::cout<<e.what()<<std::endl;
    }



    return 0;

}