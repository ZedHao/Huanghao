rocketMQ
[toc]
# 优点
##  持久化 & 多副本
## 非有序消息自动重试（Requeue）
## 死信队列 （Dead Letter Queue）
## 延迟发送
## 写入延迟
## 消息回查
## 同时支持乱序和有序
## 提供二级类型 tag

# 缺点
## 局部不有序
## 不支持广播
## 不支持事务
# 消费特点
## 集群消费（Clustering）和广播消费（Broadcasting）
# 架构
![image](https://youzhixueyuan.com/blog/wp-content/uploads/2019/07/20190731222946_55220.jpg)
## NameServer：
保存 Broker 相关元信息并给 Producer 和 Consumer 查找 Broker 信息  类似 consul 发现
## NameServer：
保存 Broker 相关元信息并给 Producer 和 Consumer 查找 Broker 信息

Broker在启动的时候会去向NameServer注册并且定时发送心跳，Producer在启动的时候会到NameServer上去拉取Topic所属的Broker具体地址，然后向具体的Broker发送消息。具体如下图：
![image](https://youzhixueyuan.com/blog/wp-content/uploads/2019/07/20190731223001_96952.jpg)

# 如何保证顺序
RocketMQ是通过将“相同ID的消息发送到同一个队列，而一个队列的消息只由一个消费者处理“来实现顺序消息。如下图：
![image](https://youzhixueyuan.com/blog/wp-content/uploads/2019/07/20190731223033_30250.jpg)
# rocketMQ  Broker 的集群部署方式有四种
1. 单 Master
1. 多 Master
 单master宕机 恢复前不可消费
1. 多 Master 多 Slave（异步复制）
 缺点是 Master 宕机时在磁盘损坏情况下会丢失极少量消息。
1. 多 Master 多 Slave（同步双写）
 能相对异步复制方式略低，发送消息的延迟会略高