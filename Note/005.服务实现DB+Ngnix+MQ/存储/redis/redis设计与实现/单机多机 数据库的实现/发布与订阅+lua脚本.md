[TOC]
# 基本概念
发布订阅功能是由 PUBLISH SUBSCRIBE PSUBSCRIBE 等命令组成

# 基本数据结构
每个 Redis 服务器进程都维持着一个表示服务器状态的 redis.h/redisServer 结构， 结构的 pubsub_channels 属性是一个字典， 这个字典就用于保存订阅频道的信息：

```
struct redisServer {
    // ...
    dict *pubsub_channels;
    // ...
};
```
 其中，字典的键为正在被订阅的频道， 而字典的值则是一个链表， 链表中保存了所有订阅这个频道的客户端。
 
 当客户端调用 SUBSCRIBE 命令时， 程序就将客户端和要订阅的频道在 pubsub_channels 字典中关联起来。
 
 ![image](https://redisbook.readthedocs.io/en/latest/_images/graphviz-cb250b1be4aaaedc9d5ddde113a80998d7f9c480.svg)
SUBSCRIBE 命令的行为可以用伪代码表示如下：


```
def SUBSCRIBE(client, channels):

    # 遍历所有输入频道
    for channel in channels:

        # 将客户端添加到链表的末尾
        redisServer.pubsub_channels[channel].append(client)
```

        
# 实现原理：


　　　　redisServer中有一个字段dict  *pubsub_channels保存了所有频道和订阅关系。键是频道，值是客户端构成的链表。

　　　　有一个订阅者就会添加到这个字段中。

　　　　退订频道时，从这里面删除。
# 不但频道订阅者会收到 匹配模式也会收到
客户端向频道发送消息， 消息被传递给正在订阅匹配模式的订阅者。
![image](http://1e-gallery.redisbook.com/_images/graphviz-f64be6f3116548c779e2357dae83bec87c04bd50.png)
# pubsub_channels 字典示例。
![image](http://1e-gallery.redisbook.com/_images/graphviz-a81ae64799855f0aa549cece1fd738475dcc4313.png)

# pubsub_patterns 链表的示例。
![image](http://1e-gallery.redisbook.com/_images/graphviz-c3da1c52931024fdcc850654a4f4e061312b1b33.png)

# 要点总结
1. 订阅信息由服务器进程维持的 redisServer.pubsub_channels 字典保存，字典的键为被订阅的频道，字典的值为订阅频道的所有客户端。
1. 当有新消息发送到频道时，程序遍历频道（键）所对应的（值）所有客户端，然后将消息发送到所有订阅频道的客户端上。
1. 订阅模式的信息由服务器进程维持的 redisServer.pubsub_patterns 链表保存，链表的每个节点都保存着一个 pubsubPattern 结构，结构中保存着被订阅的模式，以及订阅该模式的客户端。程序通过遍历链表来查找某个频道是否和某个模式匹配。
1. 当有新消息发送到频道时，除了订阅频道的客户端会收到消息之外，所有订阅了匹配频道的模式的客户端，也同样会收到消息。
1. 退订频道和退订模式分别是订阅频道和订阅模式的反操作。

 
# lua 脚本

初始化 Lua 脚本环境需要一系列步骤，其中最重要的包括：
创建 Lua 环境。
载入 Lua 库，比如字符串库、数学库、表格库，等等。
创建 redis 全局表格，包含各种对 Redis 进行操作的函数，比如 redis.call 和 redis.log ，等等。
创建一个无网络连接的伪客户端，专门用于执行 Lua 脚本中的 Redis 命令。
Reids 通过一系列措施保证被执行的 Lua 脚本无副作用，也没有有害的写随机性：对于同样的输入参数和数据集，总是产生相同的写入命令。
EVAL 命令为输入脚本定义一个 Lua 函数，然后通过执行这个函数来执行脚本。
EVALSHA 通过构建函数名，直接调用 Lua 中已定义的函数，从而执行相应的脚本。