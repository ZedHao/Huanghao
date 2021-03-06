[TOC]
文件事件（file event）
Redis服务器通过套接字与客户端（或其他Redis服务器）连接，文件事件是服务器对套接字操作的抽象
服务器与客户端（或其他Redis服务器）通信产生文件事件，服务器通过监听并处理这些事件完成一系列网络通信操作

时间事件（time event）
Redis服务器中一些操作（如serverCron函数）需在给定时间点执行，时间事件是服务器对这类定时操作的抽象

# 文件事件

![image](https://img-blog.csdn.net/20170817185916894?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbmluZzAzMjM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

## 套接字
当一个套接字准备好执行连接应答、写入、读取、关闭等操作时，会产生一个文件事件
因一个服务器通常会连接多个套接字，故多个文件事件可能会并发出现

## I/O多路复用程序
通过包装==select、epoll、evport和kqueue这些I/O多路复用函数库==实现
负责监听多个套接字，并向文件事件分派器传送产生了事件的套接字
I/O多路复用程序将产生事件的套接字放到一个队列中，通过该队列，以==有序（sequentially）、同步（synchronously）、每次一个的方式向文件事件分派器传送套接字==
上一个传送的套接字事件处理完毕后（关联的事件处理器执行完毕后），I/O多路复用程序才会向文件事件分派器传送下一个套接字

## 文件事件分派器
接收I/O多路复用程序传来的套接字，根据套接字产生的事件类型，调用相应事件处理器
事件处理器
处理器是一个函数，定义某个事件发生时，服务器应执行的动作

## 事件类型 AE_READABLE和AE_WRITABLE事件
I/O多路复用程序可监听多个套接字的AE_READABLE和AE_WRITABLE事件

- ·套接字变得可读（客户端对套接字执行write或close操作）或有新的可应答套接字（客户端对服务器监听套接字执行connect操作）出现时，套接字产生AE_READABLE事件
- ·套接字变得可写（客户端对套接字执行read操作）时，套接字产生AE_WRITABLE事件
- 注，Redis服务器优先处理同一个套接字的AE_READABLE事件，即，服务器先读套接字，后写套接字

# 事件处理器

## 连接应答处理器
对连接服务器的各个客户端进行应答
Redis初始化时，将监听套接字的AE_READABLE事件与该处理器关联
客户端连接服务器时，产生AE_READABLE事件

## 命令请求处理器
接收客户端传来的命令请求
客户端成功连接服务器后，Redis将该客户端套接字的AE_READABLE事件与该处理器关联
客户端向服务器发送命令请求时，产生AE_READABLE事件
客户端连接服务器过程中，该客户端套接字的AE_READABLE事件始终与该处理器关联

## 命令回复处理器
向客户端返回命令执行结果
服务器有命令执行结果要传送给客户端时，将客户端套接字的AE_WRITABLE事件始终与该处理器关联
客户端准备好接收命令执行结果时，产生AE_WRITABLE事件
命令执行结果传送结束后，关联解除

## 复制处理器
处理主服务器和从服务器复制操作

# 时间事件
## 定时事件
指定时间后执行一次
事件处理器返回AE_NOMORE，该事件达到一次后被删除
## 周期事件
每隔指定时间执行一次
事件处理器返回非AE_NOMORE整数值，根据该值更新when属性，让该事件在一段时间后再次到达
## 时间事件3个属性
1. ·id：全局唯一，递增分配
1. ·when：毫秒精度UNIX时间戳，记录时间事件的到达时间
1. ·timeProc：时间事件处理器，一个函数，时间事件到达时，服务器调用相应处理器处理事件


## 实现原理
服务器==将所有时间事件放在一个无序链表（when属性无序，id有序）
当时间事件执行器运行时遍历整个链表，查找所有已到达的时间事件
然后调用相应事件处理器==

## serverCron函数主要工作（执行定期操作 定期对自身资源和状态 检查和调整）
·更新服务器各类统计信息，如时间、内存占用、数据库占用情况等
·清理数据库过期键值对
·关闭和清理连接失效客户端
·尝试进行AOF或RDB持久化操作
·若服务器是主服务器，则对从服务器进行定期同步
·若服务器处于集群模式，则对集群进行定期同步和连接测试
## 文件事件和时间事件的调度和执行规则
- 1) aeApiPoll函数最大阻塞时间，由到达时间最接近当前时间的时间事件决定。该方法既可避免频繁轮询时间事件（忙等待），也可确保aeApiPoll函数不会阻塞过长时间
- 2) 因文件事件随机出现，若等待并处理完一次文件事件后仍未有任何时间事件到达，则服务器将再次等待并处理文件事件
- 3) 文件事件和时间事件的处理都是同步、有序、原子地执行的，服务器不会中途中断事件处理，也不会对事件进行抢占，故，文件事件和时间事件的处理器都会尽可能的减少程序阻塞时间，并在有需要时主动让出执行权，从而降低造成事件饥饿的可能性
- 4) 因时间事件在文件事件后执行，且事件间不会出现抢占，故时间事件实际处理时间通常会比时间事件设定的到达时间稍晚一些
