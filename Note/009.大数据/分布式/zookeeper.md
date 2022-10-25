# 概念
分布式协调系统
# 问题
1. 更新数据是怎么操作的
2. 原子操作是怎么样的
3. ZooKeeper has two basic ordering guarantees:
Linearizable writes:all requests that update the stateof  ZooKeeper  are  serializable  and  respect  precedence 
zookeeper 保证线性化写入,请问是在分布式场景下如何保证不同实例在有可能出现的网络延迟对其的线性化写入进而保证FIFO,是
全局时钟(global clock) 么

# 背景
在单个实例里,可以通过协调调度保证对同一数据的读写,在分布式场景里,如何保证一致性,容错性,可用性

