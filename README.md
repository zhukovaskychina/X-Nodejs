# X-Node.js 

监控你的node.js项目

#### 描述
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


确保该路径存在；

2,正常启动node项目


#### Contribution

1. zhukovasky 
