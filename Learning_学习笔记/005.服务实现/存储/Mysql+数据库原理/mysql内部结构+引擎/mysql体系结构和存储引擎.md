# 一、mysql体系结构和存储引擎
[TOC]
## 1.1 定义数据库和实例
数据库 和数据库实例（==数据库是文件的集合==，数据库实例==是数据管理软件，用户的增删改查，操作数据==。应用程序通过数据库实例操作数据库）
## 1.2 MYSQL体系结构


![image](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1529939280026&di=640194ea19d2b235db300c73b15bf39e&imgtype=jpg&src=http%3A%2F%2Fimg4.imgtn.bdimg.com%2Fit%2Fu%3D4044845634%2C3630039161%26fm%3D214%26gp%3D0.jpg)


https://www.cnblogs.com/zhoubaojian/articles/7866292.html
MySQL的查询优化器会自动调整where子句的条件顺序以使用适合的索引
Mysql是由SQL接口，解析器，优化器，缓存，存储引擎组成的（SQL Interface、Parser、Optimizer、Caches&Buffers、Pluggable Storage Engines）

（1） Connectors【接口】指的是不同语言中与SQL的交互

（2）Management Serveices & Utilities：系统管理和控制工具，例如备份恢复、Mysql复制、集群等 

（3）Connection Pool:连接池：管理缓冲用户连接、用户名、密码、权限校验、线程处理等需要缓存的需求

（4）SQL Interface: SQL接口：接受用户的SQL命令，并且返回用户需要查询的结果。比如select from就是调用SQL Interface 

（5）Parser: 解析器，SQL命令传递到解析器的时候会被解析器验证和解析。解析器是由Lex和YACC实现的，是一个很长的脚本

（6）Optimizer: 查询优化器，SQL语句在查询之前会使用查询优化器对查询进行优化。他使用的是“选取-投影-联接”策略进行查询。

（7） Cache和Buffer（高速缓存区）： 查询缓存，如果查询缓存有命中的查询结果，查询语句就可以直接去查询缓存中取数据。 
通过LRU算法将数据的冷端溢出，未来得及时刷新到磁盘的数据页，叫脏页。 
这个缓存机制是由一系列小缓存组成的。比如表缓存，记录缓存，key缓存，权限缓存等 

（8）Engine ：存储引擎。存储引擎是MySql中具体的与文件打交道的子系统。也是Mysql最具有特色的一个地方。 
Mysql的存储引擎是插件式的。它根据MySql AB公司提供的文件访问层的一个抽象接口来定制一种文件访问机制（这种访问机制就叫存储引擎） 
现在有很多种存储引擎，各个存储引擎的优势各不一样，最常用的MyISAM,InnoDB,BDB 
默认下MySql是使用MyISAM引擎，它查询速度快，有较好的索引优化和数据压缩技术。但是它不支持事务。 
![image](6D5881C5E20F4E4085DF4B8F51C2F054)


# mysql var路径下的.ibd .frm
.frm: 存储表结构
.myd: 存储数据，MYData的缩写
.myi: 存储索引，MYIndex的缩写

```
ibd是MySQL数据文件、索引文件，无法直接读取。
frm是表结构文件，可以直接打开
```

# mysql引擎的分类(MyISAM,InnoDB,Memory,BDB,Merge,Federated,BlackHole,CSV)

(1)==MyISAM==： 其特点是不支持事务、表锁和全文 索 引 

(2)==InnoDB==：事务型数据库的首选引擎，支持ACID事务，支持行级锁定 
使用多版本并发控制（MVCC)来获得髙并发性，并且实现了SQL标准的 4种隔离级別，默认为REPEATABLE级别。同时使用一种被称为next-key locking的策略来 避免幻读（phantom)现象的产生.除此之外，InnoDB储存引擎还提供了插入缓冲（insert b u ffer)、二 次 写 （double write)、 自 适 应 哈 希 索 引 （adaptive hash index) 、预 读 

InnoDB存储引擎支持事务，主要面向在线事务处理（OLTP) 方面的应用。其特点是 行锁设计、支持外键，并支持类似于Oracle的非锁定读，

(3)BDB：源自Berkeley DB，事务型数据库的另一种选择，支持COMMIT和ROLLBACK等其他事务特性

(4)Memory ：
Memory存储引擎（之前称为HEAP存储引擎）将表中的数据存放在内存中，如果数据 库重启或发生崩溃，表中的数据都将消失。它非常适合用于存储临时数据的临时表，以及
数据仓库中的纬度表。它默认使用哈希索引，而不是我们熟悉的B+树索引。
(5)Merge：将一定数量的MyISAM表联合而成一个整体，在超大规模数据存储时很有用

(6)Archive ：非常适合存储大量的独立的，作为历史记录的数据。因为它们不经常被读取。Archive拥有高效的插入速度，但其对查询的支持相对较差

(7)Federated：将不同的Mysql服务器联合起来，逻辑上组成一个完整的数据库。非常适合分布式应用 Cluster/NDB ：高冗余的存储引擎，用多台数据机器联合提供服务以提高整体性能和安全性。适合数据量大，安全和性能要求高的应用

(8)CSV：逻辑上由逗号分割数据的存储引擎。它会在数据库子目录里为每个数据表创建一个.CSV文件。这是一种普通文本文件，每个数据行占用一个文本行。CSV存储引擎不支持索引。

(9)BlackHole：黑洞引擎，写入的任何数据都会消失，一般用于记录binlog做复制的中继

## 1.3 Mysql表存储引擎
### 1.3.1 InnoDB存储引擎
### 1.3.2 MyISAM存储引擎 
### 1.3.3 NDB存储引擎
### 1.3.4 Memeroy存储引擎
### 1.3.5 Archive存储引擎
### 1.3.6 Federated存储引擎
### 1.3.7 Maria存储引擎
### 1.3.8 其他存储引擎
## 1.4各种存储引擎之间比较
## 1.5链接MYSQL
### 1.5.1 TCP/IP
### 1.5.2 命名管道和共享内存
### 1.5.3 Unix域套接字
