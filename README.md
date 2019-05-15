 X-Node.js 

监控你的node.js项目

#### 描述

[![Greenkeeper badge](https://badges.greenkeeper.io/zhukovaskychina/X-Nodejs.svg)](https://greenkeeper.io/)

node.js基于node-10.15.3开发
为了解决内网无法将监控数据发送给alinode，
从而让这个版本的node.js具备自动打日志，
自动生成http 访问日志
heapdump
heapprofiler

gc-tracer


从而解决线上运维node.js服务器的这难题。

本node.js项目参考alinode,
根据alinode项目特性，反向研究

#### 还需要做啥
1，cpuProfiler

2，timeline

3，report

4，更改NodeKiller

5，最终希望alinode能够开源，解决内网项目无法往外发送数据的问题


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
![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/env.png)

确保该路径存在；
 
2,正常启动node项目

![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/pm2.png)

3,打开日志文件夹：

![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/files.png)

4,查看某个文件：

![Image text](https://github.com/zhukovaskychina/X-Nodejs/blob/master/http.png)
#### Contribution

1. zhukovasky 
