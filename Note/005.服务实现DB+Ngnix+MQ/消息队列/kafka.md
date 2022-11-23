[toc]
# 基本概念
https://www.cnblogs.com/wangzhuxing/p/10051512.html
Kafka是最初由LinkedIn公司开发，是一个分布式、支持分区的（partition）、多副本的（replica），基于zookeeper协调的分布式消息系统。
# 架构 [push pull 类型]

![image](https://typora-images.oss-cn-shenzhen.aliyuncs.com/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3Kafka/20190504173245.png)
![image](03F6956127824CCA9CA955D04136F3C0)
Broker：Kafka集群包含一个或多个服务器，这种服务器被称为broker。

Topic：每条发布到kafka集群的消息都有一个类别，这个类别被称为topic，不同topic的消息分开存储。

Partition：Partition是物理上的概念，每个topic包含一个或多个partition。

Record：生产和消费一条消息，或者记录，每条记录包含：a key, a value, and a timestamp。

Offset：每个record发布到broker后，会分配一个offset，offset在单一partition中是有序递增的。

Producer：负责发布消息到kafka broker。

Consumer：消息消费者，向kafka broker读取消息的客户端。

Consumer Group：管理一组consumer实例，每个consumer属于一个特定的consumer group。


一个典型的kafka集群中包含若干producer，若干broker，若干consumer group；
Kafka有四个核心API，最常用的两种为：
Producer API：发布消息到一个或者多个kafka的topic
Consumer API：订阅一个或者多个kafka topic，并对数据进行处理

## Topic&Partition 【分区策略】
![image](2AD129C18215474DA647282C4F63A69C)
个topic可以认为一个一类消息，每个topic将被分成多个partition，每个partition在存储层面是append log文件。任何发布到此partition的消息都会被追加到log文件的尾部，每条消息在文件中的位置称为offset(偏移量)，offset为一个long型的数字，它唯一标记一条消息。每条消息都被append到partition中，是顺序写磁盘，因此效率非常高(经验证，顺序写磁盘效率比随机写内存还要高，这是Kafka高吞吐率的一个很重要的保证)。
## 文件存储 segment
一个topic 多个partition 一个partition多个segment 
producer不断发送消息，必然会引起partition文件的无限扩张，这样对于消息文件的维护以及已经被消费的消息的清理带来严重的影响
每个partition(目录)相当于一个巨型文件被平均分配到多个大小相等的segment(段)数据文件中

　segment文件由两部分组成，分别为“.index”文件和“.log”文件，分别表示为segment索引文件和数据文件。这两个文件的命令规则为：partition全局的第一个segment从0开始，后续每个segment文件名为上一个segment文件最后一条消息的offset值，
　![image](D51465CE8F0C4E08A5015A8D983DEFFF)

# 优势原因
## 可靠性保证原因【副本策略】
![image](41BC24B2EDE741F4A447B403E882F0D5)

1. Replication: 避免单机故障  parition可以有多个Replication，分布在不同broker 
2. Leader选举：每个partition会在多个replication之间选择一个leader，client读写数据都通过leader partition，其他replication为follower，负责从leader同步数据
3. 数据同步：kafka数据同步介于强同步和弱同步之间，通过两个参数控制: min.insync.replica + 发送确认acks； 即borker要保持最少replica同步，并且保证写入0-1个或者所有同步replica才成功

## batch 发送&高吞吐
![image](8CD8394A366340AC824C508278FDFDC5)
SDK提供了cache，消息由业务方先写入cache，然后当消息达到batch.size或者linger_ms之后，消息会flush到broker
buffer.memroy指定了cache的最大大小，buffer写满会block住client的写入，等待超时异常
kafka高吞吐就是通过batch来实现的，在同样qps的情况下，batch越大，吞吐越高。【是producer以batch的形式push数据，这样会极大的提高broker的性能，但是这样会增加丢失数据的风险。如果需要确保消息的可靠性，必须要将producer.type设置为sync】
# producer
![image](6A4FA295BA9B4CE49E43B5B758D49D77)
## 分区策略 
分区策略： 数据从producer写入partition的分配规则，Producer API提供了两种默认实现，random， roundrobin和key hash

## 同步|| 异步发送 模式
1. 发送并忘记：不care发送结果，调用send接口后，直接返回，kafka内部有重试机制
1. 同步发送：调用send方法，返回future对象，调用get()方法等待，为了减少cache等待时间，可以调整linger_ms为0，或者在调用get之前调用flush
1. 异步发送：send方法指定一个回调函数，在返回相应是回调该函数，回调函数工作在kafka发送线程里面，不要做太繁重的操作，避免影响发送

## ack 发送 可靠性级别
acks参数指定了必须要有多少分区副本收到消息，生产者才确定消息写入是成功的，这个参数是client配置的
1. 0-不需要等待服务器响应，也不保证服务器收到消息

1. 1-只要leader收到消息，生产者就会收到服务器发送成功的消息

1. -1 -所有参与复制的节点（in-sync）全部收到消息时，才会发送成功

## send 失败处理
1. 发送失败：一般发生在滚动升级，机器故障替换或者连接断开，Kafka并不保证100%成功，发送失败根据业务情况，选择丢弃或者重试
1. 重试：消息在发送失败的情况下，kafka内部会重试（默认3次），内部重试失败后，错误返回给业务
 重试多发生在集群滚动升级过程中，client内部重试 + 业务重试即可。
1. 消息重复：broker写入成功，response丢失，client重试，会导致写入数据重复
1. 消息乱序：在消息重复发送的情况下，会导致消息乱序，比如发送了1，2，3，重试1，broker最终数据会是1，2，3，1； 通过设置发送窗口为1，可以避免乱序
1. 幂等：通过producerId+seq_id, 保证同一个producer同一个partition内部支持消息去重。正在开发
1. 事务：待补充

# consumer
Kafka提供了两种Consumer API: low level & high level

某个partition的数据只会被某一个特定的consumer实例所消

每一个consumer或者broker的增加或者减少都会触发consumer rebalance。
## offset 管理
Offset管理
当consumer升级滚动重启，或者挂掉再拉起的时候，我们希望能够从挂起之前的消费位置 ，继续消费数据，不重不丢，那么我们就需要offset管理
每条record有一个唯一的offset，那么只要记录下offset，那么重启后，我们就可以seek到对应的offset继续消费。
offset可以由用户保存，比如spark streaming使用hdfs保存； 也可以提交到kafka保存，这里我们只讨论kakfa保存offset的情况
commit offset并不会改变当前consumer的消费位置

#自动提交 VS 手动提交
用户可以设置定期自动提交offset，用户使用方便，自动提交offset会有一定延迟，如果consumer挂了，已经消费没有及时提交的offset的数据，可能会被重复消费
手动提交：用户控制offset提交时机，可以消费一条提交一次，也可以消费一批提交一次。对于重复消费要求严格的用户

#Offset重置
想跳过数据，从最新位置开始消费，或者想重复消费历史数据，或者想从某个时间点开始消费，Kafka支持重置offset
重试offset需要停止服务再拉起，避免服务和干扰重置操作。
## consumer 顺序性
从多个partition读到数据，不保证数据间的顺序性，kafka只保证在一个partition上数据是有序的，但多个partition，根据你读的顺序会有不同
## low level API 
提供了更为灵活的消费方式，由用户指定每个consumer消费哪几个partition，不同consumer可以消费相同的partition，同时避免了reblance的劣势。
low level API劣势：client维护分配方案，partition扩容是不能动态调整cosumer策略
### 重复消费
目前消费不支持事务，数据消费了，但是没有提交offset，进程重启或者rebalance的时候，会导致重复消费
部分业务通过自定义uuid，结合redis去重，可以减低重复消费

# 高可靠分析
## At most once: 消息可能会丢，但绝不会重复传输
## At least once：消息绝不会丢，但可能会重复传输
## Exactly once：每条消息肯定会被传输一次且仅传输一次
1.  producer 给broker 发消息 失败重试 at least once
1.  consumer 先处理消息再commit crash 通信失败 offset记录失败 at least once
1.  consumer 先commit crash 再处理消息 通信失败 offset记录 at most once

## 引入消息幂等去重
## 