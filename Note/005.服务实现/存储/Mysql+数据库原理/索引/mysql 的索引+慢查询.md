[toc]


# 什么是B树和B+树
B树：有序数组+平衡多叉树； 
B+树：有序数组链表+平衡多叉树；
## B树
B树，概括来说是一个节点可以拥有多于2个子节点的二叉查找树。与自平衡二叉查找树不同，B-树为系统最优化大块数据的读和写操作。B-tree算法减少定位记录时所经历的中间过程，从而加快存取速度。普遍运用在数据库和文件系统

B 树可以看作是对2-3查找树的一种扩展，即他允许每个节点有M-1个子节点。

1. 根节点至少有两个子节点
1. 每个节点有M-1个key，并且以升序排列
1. 位于M-1和M key的子节点的值位于M-1 和M key对应的Value之间
1. 其它节点至少有M/2个子节点

![image](https://images0.cnblogs.com/blog/94031/201403/290047064066682.png)
B树的插入和平衡化操作
![image](https://files.cnblogs.com/yangecnu/btreebuild.gif)
## B+树
B+树是对B树的一种变形树，它与B树的差异在于：

1. 有k个子结点的结点必然有k个关键码；
1. 非叶结点仅具有索引作用，跟记录有关的信息均存放在叶结点中。
1. 树的所有叶结点构成一个有序链表，可以按照关键码排序的次序遍历全部记录。

![image](https://images0.cnblogs.com/blog/94031/201403/290050048129679.png)

B+树的插入和平衡化操作
![image](https://files.cnblogs.com/yangecnu/Bplustreebuild.gif)

## B树和B+树的区别和分析 和应用
B+ 树的优点在于：

1. 由于B+树在内部节点上不包含数据信息，因此在内存页中能够存放更多的key。 数据存放的更加紧密，具有更好的空间局部性。因此访问叶子节点上关联的数据也具有更好的缓存命中率。
1. B+树的叶子结点都是相链的，因此对整棵树的便利只需要一次线性遍历叶子结点即可。而且由于数据顺序排列并且相连，所以便于区间查找和搜索。而B树则需要进行每一层的递归遍历。相邻的元素可能在内存中不相邻，所以缓存命中性没有B+树好。
![image](https://images0.cnblogs.com/blog/94031/201403/290050088914733.png)
这种效率是很高的，对于N=62*1000000000个节点，如果度为1024，则logM/2N <=4，即在620亿个元素中，如果这棵树的度为1024，则只需要小于4次即可定位到该节点，然后再采用二分查找即可找到要找的值。
# B+树为啥适合被用做数据库索引

## 一步步分析为什么B+树适合作为数据库索引的结构
http://blog.codinglabs.org/articles/theory-of-mysql-index.html
https://www.cnblogs.com/aspirant/p/9214485.html

==为什么链表不适合做索引（O(n)）==
==为什么平衡二叉树不适合作为索引==

  索引文件很大，无法一次加载到内存，一次只能加载一个内存页，故需要减少磁盘的I/O操作，==没能充分利用磁盘预读功能==
  
==为什么红黑树也不适合==

  红黑树这种结构，h明显要深的多。由于逻辑上很近的节点（父子）物理上可能很远，无法利用局部性，所以红黑树的I/O渐进复杂度也为O(h)，效率明显比B-Tree差很多。
  
==为什么选择了B树==
  B树就是为了作为索引才被发明出来的的。
  归纳起来： 
  1.B树充分利用了磁盘的预读功能，因为深度小，所以磁盘的I/O读写次数少
  2.虽然B树宽，复杂度确实比BST高，但是因为是在内存中，所以避重就轻
  
##   为什么B+树比B树更适合

做这个优化的目的是为了提高区间访问的性能
B树在提高了磁盘IO性能的同时并没有解决元素遍历的效率低下的问题。正是为了解决这个问题，B+树应运而生。B+树只要遍历叶子节点就可以实现整棵树的遍历。而且在数据库中基于范围的查询是非常频繁的，而B树不支持这样的操作（或者说效率太低）。


# 索引分类
https://blog.csdn.net/timer_gao/article/details/78013826
## 索引表现形式命中规则
 键(Key)

  首先，我们发现在多数情况下，定位操作并不需要匹配整行数据。而是很规律的只匹配某一个
  或几个列的值。 例如，图中第1列就可以用来确定一条记录。这些用来确定一条数据的列，统 
  称为键(Key).

![image](https://img-my.csdn.net/uploads/201210/25/1351131021_9989.jpg)
  根据减少无效数据访问的原则，我们将键的值拿过来存放到独立的块中。并且为每一个键值添
  加一个指针， 指向原来的数据块。如图所示,
  
![image](https://img-my.csdn.net/uploads/201210/25/1351131047_7882.jpg)
![image](https://img-my.csdn.net/uploads/201210/25/1351131308_7570.jpg)
   
   辅助索引
![image](https://img-my.csdn.net/uploads/201210/25/1351131431_5894.jpg)
   
   为了降低难度和增加更新(分裂和合并B+Tree节点)的性能，InnoDB 将 Secondary B+Tree中的指针替换成了主键的键值。如图Fig.11所示：
   
![image](https://img-my.csdn.net/uploads/201210/25/1351131464_1022.jpg)
   
### 联合索引的最左前缀匹配原理

## 聚集索引和非聚集索引
###  聚集索引 
   该索引中键值的逻辑顺序决定了表中相应行的物理顺序。
   ![image](https://images.cnblogs.com/cnblogs_com/aspnethot/Pic11.JPG)
###  非聚集索引
   该索引中索引的逻辑顺序与磁盘上行的物理存储顺序不同。
   
   ![image](https://images.cnblogs.com/cnblogs_com/aspnethot/Pic10.JPG)
   
   
   
## mysql支持的索引
MySQL中，主要有四种类型的索引，分别为：B-Tree索引（B树索引），Hash索引，Fulltext（全文）索引和R-Tree索引
的表总结了何时使用聚集索引或非聚集索引（很重要）：

 

动作描述	使用聚集索引	使用非聚集索引
列经常被分组排序	应	应
返回某范围内的数据	应	不应
一个或极少不同值	不应	不应
小数目的不同值	应	不应
大数目的不同值	不应	应
频繁更新的列	不应	应
外键列	应	应
主键列	应	应
频繁修改索引列	不应	应



### B+树索引

#### Myisam 非聚集索引
MyISAM引擎使用B+Tree作为索引结构，叶节点的data域存放的是数据记录的地址。下图是MyISAM索引的原理图： 典型的B+树

==MyISAM索引文件和数据文件是分离的，索引文件仅保存数据记录的地址==

![image](http://images.cnitblog.com/blog2015/740638/201504/192216401045034.png)

在MyISAM中，主索引和辅助索引（Secondary key）在结构上没有任何区别，只是主索引要求key是唯一的，而辅助索引的key可以重复。如果我们在Col2上建立一个辅助索引，则此索引的结构如下图所示：
![image](http://blog.codinglabs.org/uploads/pictures/theory-of-mysql-index/9.png)

#### innoDb 聚集索引
 ==表数据文件本身就是按B+Tree组织的一个索引结构==这棵树的叶节点data域保存了完整的数据记录
 
==支持 B+树索引  哈希索引==

![image](http://blog.codinglabs.org/uploads/pictures/theory-of-mysql-index/10.png)

因为索引和数据不分离，所以InnoDB要求表必须有主键（MyISAM可以没有），如果没有显式指定，则MySQL系统会自动选择一个可以唯一标识数据记录的列作为主键

第二个与MyISAM索引的不同是InnoDB的辅助索引data域存储相应记录主键的值而不是地址。换句话说，InnoDB的所有辅助索引都引用主键作为data域。例

![image](http://blog.codinglabs.org/uploads/pictures/theory-of-mysql-index/11.png)

==是辅助索引搜索需要检索两遍索引：首先检索辅助索引获得主键，然后用主键到主索引中检索获得记录==
所以 innnoDb索引 不建议过长字段做主键（辅助索引会变大）
也不建议非单调索引（即自增索引），非单调会频繁插入调整，低效

#### myisam 和innodb 基于索引上的区别

自我总结：就索引上来讲，虽然两个都是用的B+树索引 但是还是有区别
   1.Mysisam的索引用的是 索引文件+数据文件分离的，叶结点存的是地址，因而查找其实是查地址，所以新增辅助索引时，且主索引和辅助索引结构无差别（只是主key 要求唯一），
   2.innodb索引和数据在一起，主索引就是主key 叶结点存的是数据，而且辅助索引 叶结点存的是主key，如果通过辅助索引找，是先找到主key 然后根据主key在主索引上找数据。

### Hash索引(Memory/heap 支持 但Innodb自适应)
  简单地说，哈希索引就是采用一定的哈希算法，把键值换算成新的哈希值，检索时不需要类似B+树那样从根节点到叶子节点逐级查找，只需一次哈希算法即可立刻定位到相应的位置，速度非常快。==innoDB存储引擎会根据表的使用情况自动为表生成hash索引，不能人为干预是否在一张表中生成hash索引。==

 ![image](https://images2015.cnblogs.com/blog/99941/201607/99941-20160706162359874-1132773212.jpg)
 
 ==innoDB存储引擎会根据表的使用情况自动为表生成hash索引，不能人为干预是否在一张表中生成hash索引。==
 自适应hash原理
 
 ![image](https://images2017.cnblogs.com/blog/1113510/201708/1113510-20170830183917780-959160821.png)
 
 Innodb存储引擎会监控对表上二级索引的查找，如果发现某二级索引被频繁访问，二级索引成为热数据，建立哈希索引可以带来速度的提升
####  B+树索引和哈希索引的明显区别
    1.等值查询 哈希索引有优势。select * from tblStaff where  id= ;
    2.范围查询，哈希不行，有序的键值哈希之后无需
    3.重复值较多时，也不合适
    
MySQL中，只有HEAP/MEMORY引擎表才能显式支持哈希索引（NDB也支持，但这个不常用），InnoDB引擎的自适应哈希索引（adaptive hash index）不在此列
  
### Fulltext（全文）索引  MyIsam 支持   MYSQL5.6的InnoDB支持全文索引 大数据文本搜索
  https://www.jianshu.com/p/645402711dac
  底层实现原理
  ![image](https://upload-images.jianshu.io/upload_images/9358011-671151d02160a998.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/675/format/webp)
  倒排索引被称为反向索引更为合适
当表上存在全文索引时，就会隐式的建立一个名为FTS_DOC_ID的列，并在其上创建一个唯一索引，用于标识分词出现的记录行。你也可以显式的创建一个名为FTS_DOC_ID的列，
   常规的索引是文档到关键词的映射：文档——>关键词
倒排索引是关键词到文档的映射：关键词——>文档
 
Number	Text	Documents
1	code	(1:6),(4:8)
2	days	(3:2),(6:2)
3	hot	(1:3),(4:4)

是word + ilist的存储结构
Text对应于word，是一个分词。Document存储的是键值对，键为FTS_DOC_ID，值为在文档中的位置，对应于ilist。其中word保存在
Auxiliary Table中，总共有六张，每张表根据word的Latin编码进行分区，下面有介绍


###  R-Tree索引 空间数据索引 
  OpenGIS  
### 前缀索引   不重复的索引值（也称为基数
  select count(distinct left(city,3))/count(*) as sel3,count(distinct left(city,4))/count(*) as sel4,count(distinct left(city,5))/count(*) as sel5, count(distinct left(city,6))/count(*) as sel6 from city_demo;  

# 索引使用策略+慢查询
分析什么样的查询会命中
情况一：全列匹配。

  1.MySQL的查询优化器会自动调整where子句的条件顺序以使用适合的索引 例如 ab是索引
    select* from tblStaff where a= 1 and b= 1;
    select* from tblStaff where b= 1 and 1= 1;
    
情况二：最左前缀匹配。
情况三：查询条件用到了索引中列的精确匹配，但是中间某个条件未提供
。
  2.可以考虑用“IN”来填补这个“坑”从而形成最左前缀：  例如abc是索引 
-   select* from tblStaff where a= 1 and b= 1; 命中 
-   select* from tblStaff where a= 1 and c= 1; 不命中
-   select* from tblStaff where a= 1 and c= 1 and b in(1,3,4);命中

情况四：查询条件没有指定索引第一列。
      select* from tblStaff where  c= 1; 不命中、

情况六：范围查询

  围列可以用到索引（必须是最左前缀），但是范围列后面的列无法用到索引。同时，索
  最多用于一个范围列，因此如果查询条件中有两个范围列则无法全用到索引。
  
情况七：查询条件中含有函数或表达式。 
  无法命中
  

## 联合索引的最左匹配原理
。
  abc是联合索引 
-   select* from tblStaff where a= 1 and b= 1; 命中 
-   select* from tblStaff where a= 1 and c= 1; 不命中
-   select* from tblStaff where a= 1 and c= 1 and b in(1,3,4);命中
 联合索引的结构
 ![image](https://i.imgur.com/vUQ8KOi.png)
 ![image](https://i.imgur.com/4jHKiv7.png)

## 该不该使用索引 （上限值）

1.第一种情况是表记录比较少，例如一两千条甚至只有几百条记录的表，没必要建索引，让查询做全表扫描就好了。至于多少条记录才算多，这个个人有个人的看法，我个人的经验是以2000作为分界线，记录数不超过 2000可以考虑不建索引

2.不建议建索引的情况是索引的选择性较低。所谓索引的选择性（Selectivity），是指不重复的索引值（也叫基数，Cardinality）与表记录数（#T）的比值：


```
Index Selectivity = Cardinality / #T
select (distinct(a))/count(*) as value from tblStaff;
```


==显然选择性的取值范围为(0, 1]，选择性越高的索引价值越大，这是由B+Tree的性质决定的。==

3.前缀索引、
  ADD INDEX `first_name_last_name4` (first_name, last_name(4));
  就是把first_name 和last_name 前4位 加上索引
  于ORDER BY和GROUP BY操作，也不能用于Covering index（即当索引本身包含查询所需全部数据时，不再访问数据文件本身）。
  first_name, left(last_name, 3) 
  
  


