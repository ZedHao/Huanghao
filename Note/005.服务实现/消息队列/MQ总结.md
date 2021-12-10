[toc]
# MQ常用衡量点
## 消息不漏发 
QoS：Quality of Service，服务质量 消息不漏
1. 最多一次（At most once）
1. 至少一次（At least once）
1. 仅一次（ Exactly once）

## 可扩展 
## 存储可靠性 【内存||磁盘】
  备份可靠性 复制与备份
### kafka ISR机制
### rocketMQ 主从
## 顺序消费
## 消息回放
## 一致性 
## 持久化存储 
## 常见的消息协议
### JMS（Java Message Service）
### AMQP（Advanced Message Queuing Protocol
### MQTT（Message Queuing Telemetry Transport）
## 常见的发送推拉模式
pull  有延迟 空pull
push 延时低 但是不好流控
push + 内存存储[NSQ 容易塞满]，
or pull + 磁盘存储[kafka]

# MQ 比较

MQ | 单topic多消费 | 发现 | 集群 | 存储 | pull&push |持久化 | 有序 | 投递语义 | 扩展 | 一致性 | 特点|缺点
---|--- | ---|--- | ---|--- | ---|--- | ---|--- | ---|---|---|
nsq|topic-chanel|nsqlookup|nsqd/生产者一一对应|内存转磁盘|push|有但低|无序|至少一次消费|||快轻| 无鉴权，无副本策略
kafka|group概念|zookeeper||磁盘|pull长轮寻|备份ISR|局部有序|三种都支持||自行定制|高吞吐，可靠性高，|延迟高
roketMQ|，||||pull|主从|始终局部有序||||回溯，二级tag，无序乱序都支持，支持重试，filter|
# kafka
## 是什么
Kafka是最初由LinkedIn公司开发，是一个分布式、支持分区的（partition）、多副本的（replica），基于zookeeper协调的分布式消息系统。
## 优势
1. 高吞吐量、低延迟【基于broker的cache 会影响准确率】 吞吐高 分布式
1. 可扩展性：Kafka集群支持热扩展【topic/partition】 动态分布topic算法
1. 持久性、可靠性：消息被持久化到本地磁盘，并且支持数据备份防止数据丢失，影响反应  
1. 高并发：支持数千个客户端同时读写
1. 容错性：允许集群中节点失败（若副本数量为n,则允许n-1个节点失败）
2. 保证消息不丢失 

## 劣势
- 1.容易重复消费 at least once 
- 2.不支持事务     
- 3.全局有序做不到【Kafka某一个固定的Partition内部的消息是保证有序的，如果一个Topic有多个Partition，partition之间的消息送达不保证有序。】
- Kafka单机超过64个队列/分区，Load会发生明显的飙高现象，队列越多，load越高，发送消息响应时间变长
-使用短轮询方式，实时性取决于轮询间隔时间；
-消费失败不支持重试；
-支持消息顺序，但是一台代理宕机后，就会产生消息乱序；
# roketMQ
## 优点
###  持久化 & 多副本
### 非有序消息自动重试（Requeue）
### 死信队列 （Dead Letter Queue）
### 延迟发送
### 写入延迟
### 消息回查
### 同时支持乱序和有序
### 提供二级类型 tag

## 缺点
### 局部不有序
### 不支持广播
### 不支持事务

# NSQ
## 优点
支持go
快 部署快 
轻量 
安全传输层协议（TLS）
与数据格式无关的消息结构

## 缺点
不支持副本
无序
没鉴权
默认不持久化



