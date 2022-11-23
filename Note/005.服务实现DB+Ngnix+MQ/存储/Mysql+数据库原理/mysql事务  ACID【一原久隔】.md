[TOC]
# 事务的ACID

混合引擎不支持事务

## 原子性（Atomicity）
　　==原子性是指事务包含的所有操作要么全部成功，要么全部失败回滚==

因此事务的操作如果成功就必须要完全应用到数据库，如果操作失败则不能对数据库有任何影响。

## 一致性（Consistency）
　　一致性是指事务必须使数据库从一个一致性状态   变换到   另一个一致性状态，也就是说一个事务执行之前和执行之后都必须处于一致性状态。

　　拿转账来说，假设用户A和用户B两者的钱加起来一共是5000，那么不管A和B之间如何转账，转几次账，事务结束后两个用户的钱相加起来应该还得是5000，这就是事务的一致性。

## 隔离性（Isolation）
　　隔离性是当多个用户并发访问数据库时，比如操作同一张表时，数据库为每一个用户开启的事务，不能被其他事务的操作所干扰，多个并发事务之间要相互隔离。

　　即要达到这么一种效果：对于任意两个并发的事务T1和T2，在事务T1看来，T2要么在T1开始之前就已经结束，要么在T1结束之后才开始，这样每个事务都感觉不到有其他事务在并发地执行。

　　关于事务的隔离性数据库提供了多种隔离级别，稍后会介绍到。

## 持久性（Durability）
　　持久性是指一个事务一旦被提交了，那么对数据库中的数据的改变就是永久性的，即便是在数据库系统遇到故障的情况下也不会丢失提交事务的操作。

　　
# 事务的隔离性 ４种隔离级别比较
[https://blog.csdn.net/qq_33290787/article/details/51924963](https://blog.csdn.net/qq_33290787/article/details/51924963)

隔离级别/读数据一致性及允许的并发副作用	| 读数据一致性|脏读|不可重复读|幻读
---|---|---|---|---|
未提交读（Read uncommitted）|                |  是| 是| 是|
已提交度（Read committed）	|语句级	         |  否|	是|	是|
可重复读（Repeatable read）	|事务级       	 |  否|	否|	是|
可序列化（Serializable）	|最高级别，事务级|	否|	否|	否|

==低级别的隔离级一般支持更高的并发处理，并拥有更低的系统开销。==
## 第1级别：Read Uncommitted(读取未提交内容)

(1)所有事务都可以看到其他未提交事务的执行结果
(2)本隔离级别很少用于实际应用，因为它的性能也不比其他级别好多少
(3)该级别引发的问题是——脏读(Dirty Read)：读取到了未提交的数据

## 第2级别：Read Committed(读取提交内容)
(1)这是大多数数据库系统的默认隔离级别（但不是MySQL默认的）
(2)它满足了隔离的简单定义：一个事务只能看见已经提交事务所做的改变
(3)这种隔离级别出现的问题是——不可重复读(Nonrepeatable Read)：不可重复读==意味着我们在同一个事务中执行完全相同的select语句时可能看到不一样的结果==。
     |——>导致这种情况的原因可能有：(1)有一个交叉的事务有新的commit，导致了数据的改变;(2)一个数据库被多个实例操作时,同一事务的其他实例在该实例处理其间可能会有新的commit
A开事务 读表a  B开事务写表a  B提交 A再读一次表a，A事务提交，A两次读的不一致
[https://blog.csdn.net/v123411739/article/details/39298127](https://blog.csdn.net/v123411739/article/details/39298127)
导致这种情况的原因可能有：(1)有一个交叉的事务有新的commit，导致了数据的改变;(2)一个数据库被多个实例操作时,同一事务的其他实例在该实例处理其间可能会有新的commit


## 第3级别：Repeatable Read(可重读) 这是MySQL的默认事务隔离级别
(2)它确保同一事务的多个实例在并发读取数据时，会看到同样的数据行
(3)此级别可能出现的问题——幻读(Phantom Read)：当用户读取某一范围的数据行时，另一个事务又在该范围内插入了新行，当用户再读取该范围的数据行时，会发现有新的“幻影” 行

```
#首先，更改隔离级别
set tx_isolation='repeatable-read';
select @@tx_isolation;
+-----------------+
| @@tx_isolation  |
+-----------------+
| REPEATABLE-READ |
+-----------------+

#事务A：启动一个事务
start transaction;
select * from tx;
+------+------+
| id   | num  |
+------+------+
|    1 |    1 |
|    2 |    2 |
|    3 |    3 |
+------+------+

#事务B：开启一个新事务(那么这两个事务交叉了)
       在事务B中更新数据，并提交
start transaction;
update tx set num=10 where id=1;
select * from tx;
+------+------+
| id   | num  |
+------+------+
|    1 |   10 |
|    2 |    2 |
|    3 |    3 |
+------+------+
commit;

#事务A：这时候即使事务B已经提交了,但A能不能看到数据变化？
select * from tx;
+------+------+
| id   | num  |
+------+------+
|    1 |    1 | --->还是看不到的！(这个级别2不一样，也说明级别3解决了不可重复读问题)
|    2 |    2 |
|    3 |    3 |
+------+------+

#事务A：只有当事务A也提交了，它才能够看到数据变化
commit;
select * from tx;
+------+------+
| id   | num  |
+------+------+
|    1 |   10 |
|    2 |    2 |
|    3 |    3 |
+------+------+
```


## 第4级别：Serializable(可串行化)
(1)这是最高的隔离级别
(2)它通过强制事务排序，使之不可能相互冲突，从而解决幻读问题。简言之,它是在每个读的数据行上加上共享锁。
(3)在这个级别，可能导致大量的超时现象和锁竞争



```
#首先修改隔离界别
set tx_isolation='serializable';
select @@tx_isolation;
+----------------+
| @@tx_isolation |
+----------------+
| SERIALIZABLE   |
+----------------+

#事务A：开启一个新事务
start transaction;

#事务B：在A没有commit之前，这个交叉事务是不能更改数据的
start transaction;
insert tx values('4','4');
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
update tx set num=10 where id=1;
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```
# 脏读 不可重复读 幻读
## 脏读（dirty read）：
 A事务读取B事务尚未提交的更改数据，并在这个数据基础上操作。如果B事务回滚，那么A事务读到的数据根本不是合法的，称为脏读。在oracle中，由于有version控制，不会出现脏读。

## 不可重复读（unrepeatable read）：
A事务读取了B事务已经提交的更改（或删除）数据。比如A事务第一次读取数据，然后B事务更改该数据并提交，A事务再次读取数据，两次读取的数据不一样。

## 幻读（phantom read）：
A事务读取了B事务已经提交的新增数据。注意和不可重复读的区别，这里是新增，不可重复读是更改（或删除）。这两种情况对策是不一样的，对于不可重复读，只需要采取行级锁防止该记录数据被更改或删除，然而对于幻读必须加表级锁，防止在这个表中新增一条数据

==值得一提的是：大多数数据库默认的事务隔离级别是Read committed，比如Sql Server , Oracle  Mysql的默认隔离级别是Repeatable read==

# 事务的实现 redo 和undo
## redo vs undo 比较
1. 事务隔离性由锁实现；
1. 原子性、一致性、持久性由数据库redo log和undo log完成。redo log重做日志，保证事务原子性和持久性，undo log保证事务一致性；
1. redo和undo都可视为一种恢复操作，redo恢复提交事务修改的页操作，undo回滚行记录到某个特定版本；
1. redo是物理日志，记录的是页的物理操作，undo是逻辑日志，根据每行记录进行记录；
1. redo log基本是顺序写，数据库运行时不需要对redo log file进行读取操作，undo log需要随机读写；
1. InnoDB存储引擎启动时，不管上次数据库是否正常关闭，都会尝试进行恢复操作。redo log记录的是物理日志，恢复速度比逻辑日志快；


## redo
![image](88D4D74D0BF445D1B2C411C16F470FA9)
当事务提交时，必须将存储引擎的日志写入磁盘。写数据前，先写日志，也称之为预写日志方式【WAL】
实现事务的持久性。由两部分组成，内存中重做日志缓存redo log buffer,易失；重做日志文件redo log file,持久的；
通过Force Log at Commit机制实现事务的持久性，即事务提交commit时，需将该事务所有日志写入到redo log file进行持久化，待事务commit操作完成才算完成；
每次将redo log buffer写入redo log file,InnoDB调用一次fsync操作；
innodb_flush_log_at_trx_commit参数控制redo log刷新到磁盘的策略；
>         0：事务提交时必须调用一次fsync；

>         1：事务提交时不进行写入redo log操作，该操作仅在master thread中完成，master thread中每1秒进行一次redo log file的fsync操作；

>         2：事务提交时将redo log写入redo log file,但仅写入文件系统的缓存中，不进行fsync；
 redo的例子
- 假设有A、B两个数据，值分别为1,2.
- A.事务开始.
- B.记录A=1到undo log.
- C.修改A=3.
- D.记录A=3到redo log.
- E.记录B=2到undo log.
- F.修改B=4.
- G.记录B=4到redo log.
- H.将redo log写入磁盘。
- I.事务提交
### redo log 和 binlog


1. redo log是在InnoDB层，binlog是MySQL数据库上层产生；
1. binlog是逻辑日志，记录的是对于的SQL语句；redo log是物理日志，记录的是对每个页的修改；
1. binlog在事务commit完成后进行一次写入；redo log在事务中不断的被写入，不是随事务提交顺序写入的

## undo
1. redo log记录事务的行为，可对页进行“重做”操作。回滚操作使用undo。
1. undo存放在数据库内部一个特殊段里，称为undo segment，位于共享表空间内；
1. undo 是逻辑日志，讲数据库逻辑的恢复到原来的样子；
> InnoDB回滚操作：
>         insert ===> delete
>         delete ===> insert
>         update ===> update

undo另一个作用是MVCC，用户读取一行记录时，该记录已经被其他事务占用，当前事务通过undo读取之前的行版本信息；
undo log会产生redo log，undo log也需持久化保护；
- 用Undo Log实现原子性和持久化的事务的简化过程
假设有A、B两个数据，值分别为1,2。
- A.事务开始.
- B.记录A=1到undo log.
- C.修改A=3.
- D.记录B=2到undo log.
- E.修改B=4.
- F.将undo log写到磁盘。
- G.将数据写到磁盘。
- H.事务提交
## redo 和undo 的恢复

# 事务的隐性提交的sql
【ddl语句】
ALTER CREATE DROP
【修改sql架构的】
CREATE USER
# 事务操作的添加
## QPS TPS
  TPS = 【com_commit+com_rollback】/time
  隐性提交不计数
 判断性能
 
 # 分布式事务
 Innodb存储引擎支持XA事务，通过XA事务可以支持分布式事务的实现。分布式事务指的是允许多个独立的事务资源（transac         tional resources）参与一个全局的事务中。事务资源通常是关系型数据库系统，也可以是其它类型的资源。
 
 InnoDB存储引擎的事务隔离级别必须设置成serialiable。
 
 XA事务允许不同数据库之间的分布式事务，如：一台服务器是mysql数据库，一台是oracle的，又有可能还有一台是sqlserver的，只要参与全局事务中的每个节点都支持XA事务。分布式事务可能在银行系统的转帐中比较常见，如一个用户需要从上海转1000元到北京的一个用户账号上面：
 ## 分布式实现
 
 分布式事务是由一个或者多个Resource Managerd，一个事务管理器Transaction Manager以及一个应用程序 Application Program组成。

1. 资源管理器：提供访问事务资源的方法，通常一个数据库就是一个资源管理器。
1. 事务管理器：协调参与全局事务中的各个事务。需要和参与全局事务中的资源管理器进行通信
1. 应用程序：定义事务的边界，指定全局事务中的操作。

在mysql中的分布式事务中，资源管理器就是mysql数据库，事务管理器为连接到mysql服务器的客户端。如下图所示：

![image](https://img-blog.csdn.net/20131026005235687?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbWNoZGJh/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
 
 分布式事务使用两段式提交（two-phase commit）的方式。在第一个阶段，所有参与全局事务的节点都开始准备，告诉事务管理器它们准备好提交了。第二个阶段，事务管理器告诉资源管理器执行rollback或者commit，如果任何一个节点显示不能commit，那么所有的节点就得全部rollback。
 
对于XA事务的支持，是在mysql体系结构的存储引擎层。因此即使不参与外部的XA事务，