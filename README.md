# asap
## 功能模块controller
### algorithm
该模块记录了leetcode上刷的题目，有一百多个，可配合算法篇一起查看 
1.[算法总结](https://mp.weixin.qq.com/s/pg94QcxIttHUBlGnHBW4zQ)
### design
该模块记录了设计模式学习内容，可配合如下文章查看
1. [Go设计模式(1)-语法](https://mp.weixin.qq.com/s/ZiobMMJ8HjdEnd5CVbo_bw)
2. [Go设计模式(4)-代码编写](https://mp.weixin.qq.com/s/Iml2GCgIpQ9MU06YwJAogw)
3. [Go设计模式(3)-设计原则](https://mp.weixin.qq.com/s/DCho5dPu-BSjpW-eI7GK9g)
4. [Go设计模式(2)-面向对象分析与设计](https://mp.weixin.qq.com/s/hZOeDStnj8DRs7xRxZ5XsQ)
### grpcclient
该模块记录了grpc的相关配置和使用，可配合如下文章查看
1. [微服务之服务框架和注册中心](https://mp.weixin.qq.com/s/sw5JVKtvYx1Jgsf5KSPXmg)
### limit
该模块记录了常用的限流算法，可配合如下文章查看
1. [限流实现1](https://mp.weixin.qq.com/s/hG6QrPPTHjqEaUVQhjDdBg)
2. [限流实现2](https://mp.weixin.qq.com/s/kvkDfCfFjbessU8UAN1O-g)


## 配置
### Redis
brew install redis
打开 redis.conf 文件，然后按 command + f 进行搜索：#requirepass foobared
修改为：
requirepass 你的密码

/usr/local/bin/redis-server /usr/local/etc/redis.conf
