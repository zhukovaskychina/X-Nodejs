# X-Node.js 

监控你的node.js项目

#### 描述
node.js基于node-10.15.3开发
为了解决内网无法将监控数据发送给alinode，
从而让这个版本的node.js具备自动打日志，
自动生成http 访问日志，
heapdump,
heapprofiler,
heaptimeline,
cpuprofile,
gc-tracer

从而解决线上运维node.js服务器的这难题。

本node.js项目参考alinode,
根据alinode项目特性，反向研究

#### 还需要做啥

1，最终希望alinode能够开源，解决内网项目无法往外发送数据的问题

2，调整IPC通道通用性问题。尝试使用mkfifo函数，从而让内存地址不固定。

已知ubuntu上ok。
需要增加脚本测试。
3，异常捕捉生成到日志中，现有的日志不正确。

4，希望alinode开源。
#### 软件架构
1,原版node.js
Software architecture description

#### Installation

git clone https://github.com/zhukovaskychina/X-Nodejs.git

./configure
 
make -C out BUILDTYPE=Release -j 8

#### Use
1,在环境变量当中确保
export NODEJS_LOG_DIR=/media/zhukovasky/8868D1D569D1C25C/nodejslogdir
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/env.png)

确保该路径存在；
 
2,正常启动node项目

![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/pm2.png)

3,打开日志文件夹：
D
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/files.png)

4,查看某个文件：
      
      
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/http.png)

5,一个实际的例子：
在线开启--trace_gc --trace_gc_nvp --trace_gc_verbose

![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/pm2-system.png)

执行命令
NodeKiller --gctrace pid
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/NodeKiller.png)
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/img/node-killer.png)

#### Contribution

1. zhukovasky 
