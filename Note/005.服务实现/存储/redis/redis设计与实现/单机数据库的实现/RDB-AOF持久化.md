[toc]
因为Redis 是内存数据库，它将自己的数据库状态储存在内存里面，所以如果不想办法将储存在内存中的数据库状态保存到磁盘里面，那么一旦服务器进程退出，服务器中的数据库状态也会消失不见
# RDB【保存的是二进制文件】
为了解决这个问题，Redis提供了RDB持久化功能，这个功能可以将Redis 在内存中的数据库状态保存到磁盘里面，避免数据意外丢失


RDB持久化既可以手动执行，也可以根据服务器配置选项定期执行，该功能可以将某个时间点上的数据库状态保存到一个RDB文件中，如图10-2所示
![image](https://images2015.cnblogs.com/blog/1067264/201612/1067264-20161201171039865-192573692.png)
RDB 持久化功能所生==成的RDB 文件是一个经过压缩的二进制文件==，通过该文件可以还原生成RDB 文件时的数据库状态，如图1 0-3所示

==因为RDB文件是保存在硬盘里面的，所以即使Redis服务器进程退出，甚至运行Redis服务器的计算机停机，但只要RDB文件仍然存在，Redis服务器就可以用它来还原数据库状态==
## RDB文件的创建与载入【save BGSAVE】
有两个Redis命令可以用于生成RDB文件，一个是SAVE，另一个是BGSAVE

```
SAVE命令会阻塞Redis服务器进程，直到RDB文件创建完毕为止，在服务器进程阻塞期间，服务器不能处理任何命令请求
```

```
和SAVE命令直接阻塞服务器进程的做法不同，BGSAVE命令会派生出一个子进程，然后由子进程负责创建RDB 文件，服务器进程(父进程)继续处理命令请求
```
## RDB 的自动载入
和使用SAVE命令或者BGSAVE命令创建RDB文件不同，==RDB 文件的载入工作是在服务启动时自动执行的==，所以Redis并没有专门用于载入RDB文件的命令，只要Redis服务器在启动时检测到RDB文件存在，它就会自动载入RDB文件


==为AOF文件的更新频率通常比RDB 文件的更新频率高，所以如果服务器开启了AOF持久化功能，那么服务器会优先使用AOF文件来还原数据库状态==
只有在AOF持久化功能处于关闭状态时，服务器才会使用RDB 文件来还原数据库状态服务器判断该用哪个文件来还原数据库状态的流程如图10-4 所示

![image](https://images2015.cnblogs.com/blog/1067264/201612/1067264-20161201172319287-676349921.png)
###  服务器在载入RDB 文件期间，会一直处于阻塞状态，直到载入工作完成为止

### 自动间隔性保存

### dirty计数器和lastsave属性

除了saveparams数组之外，服务器状态还维持着一个dirty计数器，以及一个lastsave属性:

dirty计数器记录距离上一次成功执行SAVE命令或者BGSAVE命令之后，服务器对数据库状态(服务器中的所有数据库)进行了多少次修改(包括写入、删除、更新等操作)。
lastsave属性是一个UNIX时间戳，记录了服务器上一次成功执行SAVE命令或者BGSAVE命令的时间。

当服务器成功执行一个数据库修改命令之后，程序就会对dirty计数器进行更新:命令修改了多少次数据库.dirty计数器的值就增加多少。
例如，如果我们为一个字符串键设置值:

# AOF【保存的是命令】

## AOF 配置
为了打开 AOF 持久化的功能，我们只需要将 redis.conf 配置文件中的appendonly配置选项设置为yes即可。涉及 AOF 持久化的几个常用配置如下所示：

![image](https://img-blog.csdn.net/20170718000028777?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGloYW8yMQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![image](https://img-blog.csdn.net/20170718000048891?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGloYW8yMQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## AOF 文件的生成过程具体包括命令追加，文件写入，文件同步三个步骤。 
AOF 文件的生成过程具体包括命令追加，文件写入，文件同步三个步骤。 
Redis 打开 AOF 持久化功能后，Redis 在执行完一个写命令后，都会将执行的写命令追回到 Redis 内部的缓冲区的末尾。这个过程是命令的追加过程。 
接下来，缓冲区的写命令会被写入到 AOF 文件，这一过程是文件写入过程。对于操作系统来说，调用write函数并不会立刻将数据写入到硬盘，为了将数据真正写入硬盘，还需要调用fsync函数，调用fsync函数即是文件同步的过程。只有经过文件同步过程，AOF 文件才在硬盘中真正保存了 Redis 的写命令。appendfsync 配置选项正是用来配置将写命令同步到文件的频率的，各个选项的值的含义如表 1 所示。

==appendonly.aof以 Redis 协议格式 RESP 来保存写命令==
appendonly.aof保存的命令会在 Redis 下次重启时使用来还原 Redis 数据库。

## AOF 重写
由于 Redis 会不断地将被执行的命令记录到 AOF 文件里面，所以随着 Redis 不断运行，AOF 文件的体积会越来越大。另外，如果 AOF 文件的体积很大，那么还原操作所需要的时间也会非常地长。 
为了解决 AOF 文件越来越大的问题，用户可以向 Redis 发送 BGREWRITEAOF 命令，==这个命令会移除 AOF 文件中冗余的命令来重写 AOF 文件==，使 AOF 文件的体积变得尽可能地小。 
BGREWRITEAOF 的工作原理和快照持久化命令 BGSAVE 的工作原理类似，Redis 会创建一个子进程来负责对 AOF 文件进行重写。 

# RDB和AOF对比
## RDB存在哪些优势呢？
1). 一旦采用该方式【紧凑】，那么你的整个Redis数据库将只包含一个文件，这对于文件备份而言是非常完美的。

2). 对于灾难恢复而言，RDB是非常不错的选择

3). 性能最大化。对于Redis的服务进程而言，在开始持久化时，它唯一需要做的只是fork出子进程，之后再由子进程完成这些持久化的工作，这样就可以极大的避免服务进程执行IO操作了。

4). 相比于AOF机制，如果数据集很大，RDB的启动效率会更高。

## RDB又存在哪些劣势呢？

1). 数据的高可用性低。因为系统一旦在定时持久化之前出现宕机现象，此前没有来得及写入磁盘的数据都将丢失。

2). 由于RDB是通过fork子进程，数据集较大时，可能会导致整个服务器停止服务几百毫秒，甚至是1秒钟。

## AOF的优势有哪些呢？
1). 该机制可以带来更高的数据安全性，及时，迅速不怕宕机

2). 由于该机制对日志文件的写入操作采用的是append模式，因此在写入过程中即使出现宕机现象，也不会破坏日志文件中已经存在的内容。然而如果我们本次操作只是写入了一半数据就出现了系统崩溃问题，不用担心，在Redis下一次启动之前，我们可以通过redis-check-aof工具来帮助我们解决数据一致性的问题。

3). 如果日志过大，Redis可以自动启用rewrite机制。即Redis以append模式不断的将修改数据写入到老的磁盘文件中，同时Redis还会创建一个新的文件用于记录此期间有哪些修改命令被执行。因此在进行rewrite切换时可以更好的保证数据安全性。

4). AOF格式易读
## AOF的劣势有哪些呢？

1). 对于相同数量的数据集而言，AOF文件通常要大于RDB文件。RDB 在恢复大数据集时的速度比 AOF 的恢复速度要快。

2). 根据同步策略的不同，AOF在运行效率上往往会慢于RDB。总之，每秒同步策略的效率是比较高的，同步禁用策略的效率和RDB一样高效。

二者选择的标准，就是看系统是愿意牺牲一些性能，换取更高的缓存一致性（aof），还是愿意写操作频繁的时候，不启用备份来换取更高的性能，