[toc]
# 锁产生的背景
 开发多用户、数据库驱动的应用时，最大的一个问题是：==一方面要最大程度地利用数据库的并发访问，另外一方面还要确保每个用户能以一致的方式读取和修改数据== 
 
 因此就有了锁机制。这也是数据库系统区别与文件系统的一个关键特性。
 
## 死锁
死锁是指两个或多个事务在同一资源上互相占用，并请求加锁时，而导致的恶性循环现象。当多个事务以不同顺序试围加锁同一资源时，就会产生死锁。任何时间，多个事务同时加锁-一个资源，-定产生死锁。

例如，设想下列两个事务同时处理stockPrice表:


```
事务1

START TRANSACTION;

UPDATE StockPrice SET close=45.50 WHERE stock id= 4 and date = ‘2002- 05-01’;

UPDATE StockPrice SET close=19.80 WHERE stock id= 3 and date = ‘2002 -05-02’;

COMMIT;

事务2

START TRANSACTION;

UPDATE StockPrice SET high=20.12 WHERE stock id= 3 and date = ‘2002 -05-02’;

UPDATE StockPrice SET high = 47.20 WHERE stock id= 4 and date = ‘2002-05-01’;

COMMIT;
```


如果很不幸凑巧，每个事务在处理过程中，都执行了第一个查询，更新了数据行，也加锁了该数据行。接着，每个事务都去试图更新第二个数据行，却发现该行已被(对方)加锁，然后两个事务都开始互相等待对方完成，陷入无限等待中，除非有外部因素介人，才能解除死锁。
# 查看锁 
INNODN_LOCKS 表

ＩＮＮＯＤＢ＿ＴＲＸ　表

ＩＮＮＯＤＢ＿ＬＯＣＫ＿ＷＡＩＴ　表

ＳＨＯＷ　ＥＮＧＩＮＥ　ＩＮＮＮＯＤＢ　ＳＴＡＴＵＳ　查看当前请求锁信息

https://blog.csdn.net/qq_35240226/article/details/82960507
# 锁的类别
## 行级锁

==InnoDB行锁是通过给索引上的索引项加锁来实现==的，这一点MySQL与Oracle不同，后者是通过在数据块中对相应数据行加锁来实现的。InnoDB这种行锁实现特点意味着：只有==通过索引条件检索数据，InnoDB才使用行级锁==，否则，InnoDB将使用表锁！

InnoDB存储引擎实现了如下两种标准的行级锁：
### 共享锁【S锁】 想安安静静的读 确保读一致性怕被写影响
若sessionA获得某数据表的共享锁权限，那么任何session（包括sessionA）只能对该表进行读取，不能修改该表，sessionA可以继续对该数据表加X锁，其他session可以对该数据表继续加S锁但不能加X锁，直到sessionA释放共享锁权限。加锁方式：set tables tablename read；操作如下： 

这保证了其他事务可以读A，但在T释放A上的S锁之前不能对A做任何修改。
sessionA： 

![image](https://img-blog.csdn.net/20160802165533814)

sessionB：

![image](https://img-blog.csdn.net/20160802165720487)

从上图操作可以看出，sessionA获得ha表共享锁权限以后，sessionA和sessionB都可以访问ha表，但是当sessionA想更改ha表时直接报错（ERROR 1099 (HY000): Table ‘ha’ was locked with a READ lock and can’t be updated ha表有一个共享锁不能被修改），sessionB更改ha表时并没有显示修改成功，而是在等待sessionA释放共享锁权限。

### 排它锁（写锁 或 X LOCK） 想安安静静的写 怕其他读错了
，允许事务删除或者更新一行数据
又称写锁。若事务T对数据对象A加上X锁，事务T可以读A也可以修改A，其他事务既不能读改 也不能再对A加任何锁，直到T释放A上的锁。


事务T

![image](https://img-blog.csdn.net/20160802163900355)

事务T1

![image](https://img-blog.csdn.net/20160802164053699)


从上图操作可以看出sessionA获得ha表排它锁权限以后sessionB执行访问ha表操作以后并没用显示ha数据而是在等待sessionA释放锁权限


###  锁兼容
　　当一个事务已经获得了行r的共享锁，那么另外的事务可以立即获得行r的共享锁，因为读取并没有改变行r的数据，我们称这种情况为锁兼容。

　　但如果有事务想获得行r的排它锁，则它必须等待事务释放行r上的共享锁——这种情况我们成为锁不兼容。

## 表锁

### 表共享锁（Table Read Lock）
==对MyISAM的读操作，不会阻塞其他用户对同一表的读请求，但会阻塞对同一表的写请求；==
### 表独占写锁（Table Write Lock）。
==对MyISAM的写操作，则会阻塞其他用户对同一表的读和写操作；==
MyISAM表的读操作和写操作之间，以及写操作之间是串行的。

==当一个线程获得对一个表的写锁后，只有持有锁线程可以对表进行更新操作。其他线程的读、写操作都会等待，直到锁被释放为止。==

==MyISAM【select自动加锁】==

在执行查询语句（SELECT）前，会自动给涉及的所有表加读锁，在执行更新操作（UPDATE、DELETE、INSERT等）前，会自动给涉及的表加写锁，这个过程并不需要用户干预，因此用户一般不需要直接用LOCK TABLE命令给MyISAM表显式加锁、

MyISAM存储引擎的读和写锁是互斥，读操作是串行的。那么，一个进程请求某个MyISAM表的读锁，同时另一个进程也请求同一表的写锁，MySQL如何处理呢？答案是==写进程先获得锁==。不仅如此，即使读进程先请求先到锁等待队列，写请求后到，写锁也会插到读请求之前！==这是因为MySQL认为写请求一般比读请求重要。这也正是MyISAM表不太适合于有大量更新操作和查询操作应用的原因==
，因为，大量的更新操作会造成查询操作很难获得读锁，从而可能永远阻塞。这种情况有时可能会变得非常糟糕！幸好我们可以通过一些设置来调节MyISAM的调度行为。

），　
　　

### 意向锁
 
为什么没有意向锁的话，表锁和行锁不能共存？
举个粟子（此时假设行锁和表锁能共存）： 事务A锁住表中的一行（写锁）。事务B锁住整个表（写锁）。

但你就会发现一个很明显的问题，事务A既然锁住了某一行，其他事务就不可能修改这一行。这与”事务B锁住整个表就能修改表中的任意一行“形成了冲突。所以，没有意向锁的时候，行锁与表锁共存就会存在问题！

有了意向锁之后，前面例子中的事务A在申请行锁（写锁）之前，数据库会自动先给事务A申请表的意向排他锁。当事务B去申请表的写锁时就会失败，因为表上有意向排他锁之后事务B申请表的写锁时会被阻塞。不需要我们程序员使用代码来申请。

申请意向锁的动作是数据库完成的，就是说，事务A申请一行的行锁的时候，数据库会自动先开始申请表的意向锁，

==因为InnoDB支持的是行级别的锁，所以意向锁其实不会阻塞除全表扫以外的任何请求。==
　==　InnoDB存储引擎支持多粒度锁定，这种锁定允许在行级上的锁和表级上的锁同时存在。==为了支持在不同粒度上进行加锁操作，InnoDB存储引擎支持一种额外的锁方式，我们称之为意向锁。
#### 意向共享锁(IS Lock），事务想要获得一个表中某几行的共享锁。
事务打算给数据行加行共享锁，事务在给一个数据行加共享锁前必须先取得该表的IS锁。

#### 意向排它锁（IX Lock），事务想要获得一个表中某几行的排它锁。
事务打算给数据行加行排他锁，事务在给一个数据行加排他锁前必须先取得该表的IX锁。

## 页锁

## 乐观锁&悲观锁

### 　悲观锁（Pessimistic Concurrency Control，PCC）：
假定会发生并发冲突，屏蔽一切可能违反数据完整性的操作。至于怎么加锁，加锁的范围也没讲。

悲观锁的特点是先获取锁，再进行业务操作，即“悲观”的认为获取锁是非常有可能失败的，因此要先确保获取锁成功再进行业务操作。

==通常所说的“一锁二查三更新”即指的是使用悲观锁==。通常来讲在数据库上的悲观锁需要数据库本身提供支持，即通过常用的select … for update操作来实现悲观锁。

### 乐观锁（Optimistic Concurrency Control，OCC） 典型先后读取版本号

假设不会发生并发冲突，只在提交操作时检查是否违反数据完整性。

乐观锁的特点先进行业务操作，不到万不得已不去拿锁。即“乐观”的认为拿锁多半是会成功的，因此在进行完业务操作需要实际更新数据的最后一步再去拿一下锁就好。
 
#### 乐观锁常见的两种使用方式

==乐观锁一般会使用版本号机制或CAS算法实现。==

#### 1. 版本号机制
一般是在数据表中加上一个数据版本号version字段，表示数据被修改的次数，当数据被修改时，version值会加一。当线程A要更新数据值时，在读取数据的同时也会读取version值，在提交更新时，若刚才读取到的version值为当前数据库中的version值相等时才更新，否则重试更新操作，直到更新成功。
#### 2. CAS算法 即compare and swap（比较与交换）
 是一种有名的无锁算法。无锁编程，即不使用锁的情况下实现多线程之间的变量同步，也就是在没有线程被阻塞的情况下实现变量的同步，所以也叫非阻塞同步（Non-blocking Synchronization）

。这样处理的逻辑是，首先检查某块内存的值是否跟之前我读取时的一样，如不一样则表示期间此内存值已经被别的线程更改过，舍弃本次操作，否则说明期间没有其他线程对此内存值操作，可以把新值设置给此块内存。
反向例子 ABA 问题

# 三级封锁协议

1. 一级封锁协议：事务T中如果对数据R有写操作，必须在这个事务中对R的第一次读操作前对它加X锁，直到事务结束才释放。事务结束包括正常结束（COMMIT）和非正常结束（ROLLBACK）。
1. 二级封锁协议：一级封锁协议加上事务T在读取数据R之前必须先对其加S锁，读完后方可释放S锁。  
1.  三级封锁协议 ：一级封锁协议加上事务T在读取数据R之前必须先对其加S锁，直到事务结束才释放。
　　
　　可见，三级锁操作一个比一个厉害（满足高级锁则一定满足低级锁）。但有个非常致命的地方，一级锁协议就要在第一次读加x锁，直到事务结束。几乎就要在整个事务加写锁了，效率非常低。三级封锁协议只是一个理论上的东西，实际数据库常用另一套方法来解决事务并发问题。

# INNODB的锁相关底层和特性 
Mutex 数据结构实现锁
## MVCC
 开始接触多版本的时候，我的疑问是innodb对每个行要存储多个版本是多么浪费存储空间呀？

然而进一步了解，原来所谓的多版本只是innodb聪明地撒了个谎，多个版本是通过undo日志实现的，这里可以理解为既然==undo日志包括了所有用来恢复历史版本数据的信息==，那么我们只要将“不同版本”指针指向不同时间节点的undo日志即可，这样读取的时候通过对不同时间节点的undo日志进行恢复从而得到不同的版本数据。同时对于undo日志的读取是不需要加锁的，因此这极大地提高了数据库的并发性

### 基本原理
innodb MVCC主要是为Repeatable-Read事务隔离级别做的。在此隔离级别下，A、B客户端所示的数据相互隔离，互相更新不可见

了解innodb的行结构、Read-View的结构对于理解innodb mvcc的实现由重要意义

innodb存储的最基本row中包含一些额外的存储信息 DATA_TRX_ID，DATA_ROLL_PTR，DB_ROW_ID，DELETE BIT

1. 6字节的DATA_TRX_ID 标记了最新更新这条行记录的transaction id，每处理一个事务，其值自动+1
1. 7字节的DATA_ROLL_PTR 指向当前记录项的rollback segment的undo log记录，找之前版本的数据就是通过这个指针
1. 6字节的DB_ROW_ID，当由innodb自动产生聚集索引时，聚集索引包括这个DB_ROW_ID的值，否则聚集索引中不包括这个值.，这个用于索引当中
1. DELETE BIT位用于标识该记录是否被删除，这里的不是真正的删除数据，而是标志出来的删除。真正意义的删除是在commit的时候

![image](https://images2015.cnblogs.com/blog/268981/201512/268981-20151221230743968-739828690.png)

具体的执行过程

==begin->用排他锁锁定该行->记录redo log->记录undo log->修改当前行的值，写事务编号，回滚指针指向undo log中的修改前的行==

上述过程确切地说是描述了UPDATE的事务过程，其实undo log分insert和update undo log，因为insert时，原始的数据并不存在，所以回滚时把insert undo log丢弃即可，而update undo log则必须遵守上述过程

MVCC的实现，通过保存数据在某个时间点的快照来实现的。这意味着一个事务无论运行多长时间，在同一个事务里能够看到数据一致的视图。根据事务开始的时间不同，同时也意味着在同一个时刻不同事务看到的相同表里的数据可能是不同的。


## 一致性的非锁定读操作：
　　是指InnoDB存储引擎通过行多版本并发控制（MVCC）的方式来读取当前执行时间数据库中行的数据。

　　如果读取的行正在执行DELETE、UPDATE操作，这时读取操作不会因此等待行上的锁释放，相反，存储引擎会去读取一个快照数据。

快照数据是指该行之前版本的数据，该实现是通过Undo段来实现。而Undo用来在事务中回滚数据，因而快照数据本身是没有额外的开销。此外，读取快照数据是不必要上锁的，因为没有必要对历史的数据进行修改。
　　所以，对于Read Commited的事务隔离级别而言，其实违反了事务的隔离性。
　　
![image](EBBC8CB9FF064AFCB842DC9B1704ED34)
 

> 　　在Read Comitted事务隔离级别下，对于快照数据，总是读取被锁定行的最新一份快照数据。
> 
> 　　在Repeatable Read事务隔离级别下，对于快照数据，总是读取事务开始时的行数据版本。

问题一：上面Mark A处显然t1已经给记录加了X锁，并且在事务内修改了数据，此时t2看到的数据是什么？
t2此时看到的应该是历史版本的数据，也就是t1修改之前的数据 t3=bb


```
mysql> select * from mytest where t2='bb';
+------+------+------+------+
| t1   | t2   | t3   | t4   |
+------+------+------+------+
| a    | bb   | bb   | ccc  |
+------+------+------+------+
1 row in set (0.00 sec)
```


问题二：上面Mark B处事务t1已经提交此时t2看到的数据是什么？

   此时t2看到的是什么跟此时数据库的隔离级别有关系
   1.隔离级别是 RR【可重复读】 即事务t2在事务前后应该一致，还是老数据 基于mvcc【违反了RC】
   2.隔离级别是 RC【读可提交】READ COMMITTED，则t2在Mark B处看到的应该是新数据。

### 总结

再总结一下，一致性非锁定读讲的是一条记录被加了X锁其他事务仍然可以读而不被阻塞，是通过innodb的行多版本实现的，行多版本并不是实际存储多个版本记录而是通过undo实现。一致性锁定读讲的是我可以通过SELECT语句显式地给一条记录加X锁从而保证特定应用场景下的数据一致性。。
## 自增长锁 【auto_increment】
   为了插入性能，锁不是事务完成才释放，而是自增长sql擦汗如即释放

# InnoDB存储引擎 锁的算法

- 1.Record Lock：单个行记录上的锁，锁定的对象是索引，而不是数据。
   Record Lock总是会锁住索引记录，如果InnoDB存储引擎建立的时候没有设置任何一个索引，这时InnoDB存储引擎会使用隐式的主键来进行锁定。
- 2.Gap Lock：间隙锁，锁定一个范围的索引，但不包含记录本身
- 3.Next-Key Lock： Gap Lock + Record Lock，锁定一个范围的索引，并且锁定记录本身。

　==　InnoDB的默认事务隔离级别是READ REPEATABLE，采用Next-Key Lock算法，解决了不可重复读（幻读）问题。==

　　在Next-Key Lock 算法下，不仅仅是锁住扫描到的索引，而且还锁住这些索引覆盖的范围（gap）。因此对于这个范围内的插入都是不允许的。
## 详解　 Gap 锁和 Next-key Lock
![image](A04867E8ECE84E138F49AFAA91371375)

间隙锁（Gap Lock）一般是针对非唯一索引而言的，test表中的v1（非唯一索引）字段值可以划分的区间为：

（-∞,1）

（1,3）

（3,4）

（4,5）

（5,7）

（7,9）

（9, +∞）

假如要更新v1=7的数据行，那么此时会在索引idx_v1对应的值，也就是v1的值上加间隙锁，锁定的区间是（5,7）和（7,9）。==同时找到v1=7的数据行的主键索引和非唯一索引，对key加上锁。==

# 锁升级
　　指将当前锁的粒度降低，比如1000个行锁升级为一个页锁，或者将页锁升级为表锁。

　　InnoDB不存在锁升级的问题。

　　其根据每个事务访问的每个页对锁进行管理，采用的是位图的方式。

　　不管事务锁住页中的一个记录还是多个记录，其开销是一样的。





