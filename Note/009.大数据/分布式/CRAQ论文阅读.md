#
Object Storage on CRAQ: High-Throughput Chain Replication for Read-Mostly Workloads[针对读多写少的场景下的链式复制]
针对读多写少场景下的高可用链式复制
背景
分布式的三个要素CAP
- 一致性（C）：在分布式系统中的所有数据备份，在同一时刻是否同样的值。
- 高可用（A）：在集群中一部分节点故障后，集群整体是否还能响应客户端的读写请求。（可用性不仅包括读，还有写）
- 分区容忍性（P）：集群中的某些节点在无法联系后，集群整体是否还能继续进行服务。


 一致性协议包括但不限于以下几类：
    1.强一致性：多进程并发读写访问时，要求更新过的数据能被后续的访问都能看到。
    2.弱一致性:如果能容忍后续的部分或者全部访问不到。
    3.最终一致性：如果经过一段时间后要求能访问到更新后的数据，就是最终一致性。
链式复制
基本模型:
副本按照一定顺序关系组织成复制链， 写操作[head]---串行化---读操作[tail]。

      figure 1: 所有的写操作由头结点处理，读操作/写操作的确认由尾结点处理。

example：

节点上的操作和下一个节点同步是FIFO的，因此后续节点上的数据记录肯定是前面节点的前缀。

HEAD:(a=1)(b=2)(c=3)(d=4) 
MIDDLE:(a=1)(b=2)(c=3) 
TAIL:(a=1)(b=2)
头结点拥有四个写操作，其中a=1，b=2已经到达尾节点，由尾节点向客户端返回成功。
c=3，d=4这两条记录尾节点还未返回成功。
这时如果客户端的读发生在尾节点，因此只能读到(a=1, b=2)这两条记录，(c=3, d=4)这两条记录现在读不到。
分析异常情况：
- 头节点宕机 头节点宕机，记录d=4丢失，但是由于对d=4这条记录的确认没有返回给客户端，因此并不影响数据的一致性。
- 尾节点宕机 TAIL节点宕机后，TAIL节点的父节点接替它成为尾节点，由于它上面的记录总是比尾节点上的记录多，而且尾节点上的记录是它父节点记录的前缀。因此客户端看见的数据变多了，并没有对视数据，因此也是强一致的。
- 中间节点宕机 中间节点宕机后，并不影响读的一致性，由上图可知TAIL节点上的数据是HEAD节点数据的前缀，只要父节点将数据同步下来即可。
分析优劣:
- 优势:强一致性[Strong Consistency]，client的成功写入和读取都是依赖tail节点返回
- 劣势:
  - 读取的吞吐率较差[Lower-Throughput]即可用性较差,不符合Availability。
  - 尤其不适用读多写少的场景
即是符合CP的分布式存储

CRAQ[Chain Replication with Apportioned Queries]
CR链式复制的改良版
背景:

those  designed  for  inter-active  use  by  customers,  sacrifice  stronger  consistencyproperties in the desire for greater availability and higher throughput[为了获得更高的可用性和更高的吞吐量而牺牲了强一致性]
基本模型:


FAQ:
从图可见:读请求分散各个节点。一致性如果保证?
一致性保证:



1. 每一个非尾结点可以保存多版本的数据，版本号单调自增。每个版本可能是clean或者dirty，在开始时所有的均为clean.
2. 当头部节点收到写请求，则将自身设置为dirty，传递新版本号给下一个节点。到达尾部节点时，尾部节点设置为clean，依次回传，收到ack之后各自设置为clean。
3. 当节点收到读请求的时候，如果最终版本号为clean则回复该版本号对应数据，否则问询尾部节点上次提交为clean的版本，并回复该版本对应的数据。

类似双向链表
分析优劣:
- 可用性: tail分散了读请求,极大的缓解了读请求,提高可用性
- 一致性:不是强一致性,是最终一致性
故障恢复
 和CR类似 
单链路中心和多链路中心
我们讨论应用程序如何在CRAQ中指定单个链数据中心内以及跨多个数据中心的各种链布局方案。然后，我们描述如何使用协调服务来存储链元数据和组成员身份信息
- Implicit Datacenters & Global Chain Size:(待指定的数据中心)


一致性哈希用于唯一的数据中心标识符
- Explicit Datacenters & Global Chain Size:[指定的数据中心]


- Explicit Datacenter Chain Sizes

分布式协调用的是zookeeper+一致性hash 多链路中心 降低写的负载均衡
其他扩展
Mini-Transactions on CRAQ
Lowering Write Latency with Multicast[基于多播协议更低的写延迟]
https://juejin.im/post/5d5decf6f265da03bd05218a

相较于链式的传播数据
CRAQ可以利用多播协议[41]来提高写入性能，与其在链上向下传播完整的写操作，而增加了与链长成比例的等待时间，不如将其实际值组播到整个链上。然后，只需要在链下传播一条小的元数据消息，以确保所有副本都在尾部之前收到写操作。如果节点由于某种原因未收到多播，则该节点可以在接收到写入提交消息之后并进一步传播提交消息之前从其前节点获取对象。被发送到多播组，而不是沿链向后传播。这样既减少了节点对象在写入后重新进入干净状态所花费的时间，又减少了客户端感知的写入延迟。再次，在多播确认时不需要顺序或可靠性保证，如果一个节点在链未收到确认，它将在下一次读取操作要求查询尾部时重新进入清洁状态。
对比数据





FAQ:
1. CR这种方式不流行的原因是什么？ 有什么问题？
2. 本文的系统相比很久前就提出过的各种chain replication系统，主要改进在什么地方?
3. Item 4 in Section 2.3 says that, if a client read request arrives and the latest version is dirty, the node should ask the tail for the latest committed version. Suppose, instead, that the node replied with its most recent clean version (ignoring any dirty version and not sending a version query to the tail). This change would cause reads to reflect the most recent committed write that the node is aware of. Explain how this could lead to violations of linearizability -- or violations of the paper's goal of strong consistency.
