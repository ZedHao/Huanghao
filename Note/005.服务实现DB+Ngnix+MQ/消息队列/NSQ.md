[toc]
# 基础概念
NSQ是一个基于Go语言的分布式实时消息平台，它基于MIT开源协议发布，由bitly公司开源出来的一款简单易用的消息中间件。
# 优缺点
## 优点
支持go
快 部署快 
轻量 
可扩展 水平扩展 用户自己负责负载均衡

## 缺点
不支持副本
无序
没鉴权
默认不持久化 
保证消息不丢 
落硬盘 简单
生产者不能动态发现nsq
数据缺少备份
无法实现顺序消费 
# 架构
1. nsqd 是接收、队列和传送消息到客户端的守护进程。
1. nsqlookupd 是管理的拓扑信息，并提供了最终一致发现服务的守护进程。
1. nsqadmin 是一个 Web UI 来实时监控集群(和执行各种管理任务)。
1. utilities ：常见基础功能、数据流处理工具，如 nsq_stat、nsq_tail、nsq_to_file、nsq_to_http、nsq_to_nsq、to_nsq
2.  nsqlookup  法务发现机制 topic 和chanel
 /lookup接口，获取拥有对应topic的nsq列表。注意看代码，里面是遍历了nsqlookup的列表，然后把所有lookup的返回结构，进行合并

因为nsq认为一个producer，就对应一个nsq。

![image](90EE51076F4E4044B31C4E68A32FB905)
# 消息的生命周期
产者往本地的nsqd中发送消息.这个过程会开启一个连接, 并发送一个带有topic和消息体的PUB的命令到nsqd中. 我们假如是发送一个events的topic

events topic 会对消息进行copy,并多路发送到各个channel中, 我们假设有三个channel, 那么这个流程会如下图描述所示:

在channel中的每条消息会被放进队列中, 直到消息被worker所消费掉, 如果队列占用的内存超出限制, 消息会被写进硬盘

nsqd节点会首先向nsqlookd节点广播它的位置信息, 一旦这些信息被nsqlookupd注册上, workers就会发现这些nsqd节点,包括这些节点的events topic
![image](https://lihaoquan.me/media/20160620_nsq.png)

![image](https://lihaoquan.me/media/20160620_nsq-lookups.png)

![image](https://segmentfault.com/img/remote/1460000009194611?w=420&h=281)

![image](https://pic1.zhimg.com/80/v2-973fda23d1ae894ce73054cc404a7560_hd.jpg)

# 服务发现（nsqlookup）
的设计是基于最终一致性，nsqlookup并不协调状态或者应答
# 如何保证 at least once 
重试机制
# 流控 
Nsq流控的方式非常简单，当消费者和nsq建立好连接，准备好接受消息时，会给nsq回一个RDY的响应，同时带上一个rdy_count，代表准备接受消息的数量，于是nsq会给消费者推送消息，每推送一条，对应连接的rdy_count就减1(如果是批量推送，则批量减)，直到连接的rdy_count变成0，则不再继续推送消息。