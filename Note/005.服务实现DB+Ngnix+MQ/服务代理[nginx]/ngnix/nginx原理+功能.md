[TOC]
https://www.jianshu.com/p/6215e5d24553
http://www.imooc.com/article/19278
# Nginx是什么
![image](https://images2015.cnblogs.com/blog/398358/201602/398358-20160202133724350-1807373891.jpg)
## 怎么理解反向代理
![image](https://upload-images.jianshu.io/upload_images/6807865-2cede76e2384c39f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/349/format/webp) 
![image](https://upload-images.jianshu.io/upload_images/6807865-90603b54f3e3e521.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/354/format/webp)
## nginx 结构分类
Nginx的模块从结构上分为核心模块、基础模块和第三方模块：

### 核心模块：HTTP模块、EVENT模块和MAIL模块
### 基础模块：HTTP Access模块、HTTP FastCGI模块、HTTP Proxy模块和HTTP Rewrite模块，
### 第三方模块：HTTP Upstream Request Hash模块、Notice模块和HTTP Access Key模块。

## Nginx的模块从功能上分为如下四类：

Core(核心模块)：构建nginx基础服务、管理其他模块。

Handlers（处理器模块）：此类模块直接处理请求，并进行输出内容和修改headers信息等操作。

Filters （过滤器模块）：此类模块主要对其他处理器模块输出的内容进行修改操作，最后由Nginx输出。

Proxies （代理类模块）：此类模块是Nginx的HTTP Upstream之类的模块，这些模块主要与后端一些服务比如FastCGI等进行交互，实现服务代理和负载均衡等功能。

## nginx工作流程
1. 用户通过域名发出访问Web服务器的请求，该域名被DNS服务器解析为反向代理服务器的IP地址；
1. 反向代理服务器接受用户的请求；
1. 反向代理服务器在本地缓存中查找请求的内容，找到后直接把内容发送给用户；
1. 如果本地缓存里没有用户所请求的信息内容，反向代理服务器会代替用户向源服务器请求同样的信息内容，并把信息内容发给用户，如果信息内容是缓存的还会把它保存到缓存中。

# ngnix CGI
## CGI
为了解决不同的语言解释器(如php、python解释器)与webserver的通信，于是出现了cgi协议。只要你按照cgi协议去编写程序，就能实现语言解释器与webwerver的通信。如php-cgi程序。
## FAST-CGI
有了cgi协议，解决了php解释器与webserver通信的问题，webserver终于可以处理动态语言了。但是，webserver每收到一个请求，都会去fork一个cgi进程，请求结束再kill掉这个进程。这样有10000个请求，就需要fork、kill php-cgi进程10000次。有没有发现很浪费资源

==于是，出现了cgi的改良版本，fast-cgi。fast-cgi每次处理完请求后，不会kill掉这个进程，而是保留这个进程，使这个进程可以一次处理多个请求。这样每次就不用重新fork一个进程了，大大提高了效率。==

## Php-fpm是什么
php-fpm即php-Fastcgi Process Manager. 
php-fpm是 FastCGI 的实现，并提供了进程管理的功能。 
进程包含 master 进程和 worker 进程两种进程。 
master 进程只有一个，负责监听端口，接收来自 Web Server 的请求，而 worker 进程则一般有多个(具体数量根据实际需要配置)，每个进程内部都嵌入了一个 PHP 解释器，是 PHP 代码真正执行的地方。
## php-cgi 和 php-fpm的区别
php-cgi是早期php官方出品的fastcgi管理器，不支持平滑重启，改了php.ini就要kill掉原来的php-cgi再重新启动才能生效；不支持动态worker调度，只能一开始指定要起几个worker。

php-fpm是从5.3.3才加入的fastcgi进程管理器，加入了动态调度功能，可以根据请求来访压力变化动态增减worker进程数量；支持reload指令，让worker进程在完成当前请求后重启，并应用php.ini新配置。

## nginx 运行原理
Nginx+Php-fpm运行原理详解
![image](https://upload-images.jianshu.io/upload_images/6807865-aefcf693d5c431f0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/687/format/webp)

```
                        www.example.com
                                |
                                |
                              Nginx
                                |
                                |
                        路由到www.example.com/index.php
                                |
                                |
                        加载nginx的fast-cgi模块
                                |
                                |
                        fast-cgi监听127.0.0.1:9000地址
                                |
                                |
                        www.example.com/index.php请求到达127.0.0.1:9000
                                |
                                |
                        php-fpm 监听127.0.0.1:9000
                                |
                                |
                        php-fpm 接收到请求，启用worker进程处理请求
                                |
                                |
                        php-fpm 处理完请求，返回给nginx
                                |
                                |
                        nginx将结果通过http返回给浏览器
```

[link](https://blog.csdn.net/u013474436/article/details/52972699)
https://www.cnblogs.com/augus007/articles/8085185.html

![image](http://img.blog.csdn.net/20151111145052945)

php-fpm只是一个php-fastcgi的管理器，为php提供管理服务。

这个CGI-APP就是PHP程序。而FastCGI的上游是Nginx，他们之间有一个通信载体，即图中的socket。在我们上文图3的配置文件中，fastcgi_pass所配置的内容，便是告诉Nginx你接收到用户请求以后，你该往哪里转发，在我们图3中是转发到本机的一个socket文件，这里fastcgi_pass也常配置为一个http接口地址（这个可以在php-fpm.conf中配置）。而上图5中的Pre-fork，则对应着我们PHP-FPM的启动，也就是在我们启动PHP-FPM时便会根据用户配置启动诸多FastCGI触发器（FastCGI Wrapper）。

从图6中可以看到，fastcgi_param所声明的内容，将会被传递给“FastCGI server”，那这里指的就是fastcgi_pass所指向的server，也就是我们Nginx+PHP模式下的PHP-FPM所管理的FastCGI进程，或者说是那个socket文件载体。这时，有的同学会问：“为什么PHP-FPM管理的那些FastCGI进程要关心这些参数呢？”，好问题，我们一起想想我们做PHP应用开发时候有没有用到 $_SERVER 这个全局变量，它里面包含了很多服务器的信息，比如包含了用户的IP地址。同学们不想想我们的PHP身处socket文件之后，为什么能得到远端用户的IP呢？聪明的同学应该注意到图4中的一个fastcgi_param配置 REMOTE_ADDR ，这不正是我们在PHP中用 $_SERVER[‘REMOTE_ADDR’] 取到的用户IP么。的确，Nginx这个模块里fastcgi_param参数，就是考虑后端程序有时需要获取Webserver外部的变量以及服务器情况，那么ngx_http_fastcgi_module就帮我们做了这件事。真的是太感谢它啦！
## 如何处理请求
　　首先，Nginx在启动时，会解析配置文件，得到需要监听的端口与IP地址，然后在Nginx的master进程里面，先初始化好这个监控的socket(创建socket，设置addrreuse等选项，绑定到指定的IP地址端口，再listen)，然后再fork(一个现有进程可以调用fork函数创建一个新进程。由fork创建的新进程被称为子进程 )出多个子进程出来，然后子进程会竞争accept新的连接。 
　　此时，客户端就可以向Nginx发起连接了。当客户端与Nginx进行三次握手，与Nginx建立好一个连接后，某一个子进程会accept成功，得到这个建立好的连接的socket，然后创建Nginx对连接的封装，即ngx_connection_t结构体。 
　　接着，设置读写事件处理函数并添加读写事件来与客户端进行数据的交换。最后，Nginx或客户端来主动关掉连接，到此，一个连接就寿终正寝了。

# ngnix 特点
1.       跨平台、配置简单
1.        ==非阻塞、高并发连接==：处理2-3万并发连接数，官方监测能支持5万并发
1.         内存消耗小：开启10个nginx才占150M内存，Nginx采取了分阶段资源分配技术
1.         nginx处理静态文件好,耗费内存少
1.         ==内置的健康检查功能==：如果有一个服务器宕机，会做一个健康检查，再发送的请求就不会发送到宕机的服务器了。重新将请求提交到其他的节点上。
1.         节省宽带：支持GZIP压缩，可以添加浏览器本地缓存
1.         稳定性高：宕机的概率非常小
1.      ==master/worker结构：一个master进程，生成一个或者多个worker进程==
1.         接收用户请求是异步的：浏览器将请求发送到nginx服务器，它先将用户请求全部接收下来，再一次性发送给后端web服务器，极大减轻了web服务器的压力
1.         一边接收web服务器的返回数据，一边发送给浏览器客户端
1.         网络依赖性比较低，只要ping通就可以负载均衡
1.         可以有多台nginx服务器
1.         事件驱动：==通信机制采用epoll模型==
2.         解决Ajax跨域问题；
3.         节约了有限的IP地址资源
4.         减少WEB服务器压力，提高响应速度
5.         作为真实服务器的缓冲，解决瞬间负载量大的问题；


# nginx 基本功能

## 1、Http代理，反向代理：

Nginx在做反向代理时，提供性能稳定，并且能够提供配置灵活的转发功能。Nginx可以根据不同的正则匹配，采取不同的转发策略，比如图片文件结尾的走文件服务器，动态页面走web服务器，只要你正则写的没问题，又有相对应的服务器解决方案，你就可以随心所欲的玩。并且Nginx对返回结果进行错误页跳转，异常判断等。如果被分发的服务器存在异常，他可以将请求重新转发给另外一台服务器，然后自动去除异常服务器。

## 2、负载均衡算法 upstream 
nginx 的 upstream目前支持 4 种方式的分配 
### 1)、轮询（默认） 
      每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器down掉，能自动剔除。 
### 2)、weight 
      指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。 
### 2)、ip_hash 
      每个请求按访问ip的hash结果分配，这样每个访客固定访问一个后端服务器，可以解决session的问题。  
     
对客户端请求的ip进行hash操作，然后根据hash结果将同一个客户端ip的请求分发给同一台服务器进行处理，可以解决session不共享的问题。

![image](https://images2015.cnblogs.com/blog/398358/201602/398358-20160201162405944-676557632.jpg)
### 3)、fair（第三方） 
      按后端服务器的响应时间来分配请求，响应时间短的优先分配。  
### 4)、url_hash（第三方）
--------------------- 
作者：Watson2016 
来源：CSDN 
原文：https://blog.csdn.net/watson2016/article/details/77938678/ 
版权声明：本文为博主原创文章，转载请附上博文链接！

Nginx提供的负载均衡策略有2种：内置策略和扩展策略。内置策略为轮询，加权轮询，Ip hash。扩展策略，就天马行空，只有你想不到的没有他做不到的啦，你可以参照所有的负载均衡算法，给他一一找出来做下实现。

![image](https://images2015.cnblogs.com/blog/398358/201602/398358-20160202133753382-1863657242.jpg)


## 4、web缓存

Nginx可以对不同的文件做不同的缓存处理，配置灵活，并且支持FastCGI_Cache，主要用于对FastCGI的动态程序进行缓存。配合着第三方的ngx_cache_purge，对制定的URL缓存内容可以的进行增删管理


# 为什么不使用多线程 Ngnix和Apache？

Apache: 创建多个进程或线程，而每个进程或线程都会为其分配cpu和内存（线程要比进程小的多，所以worker支持比perfork高的并发），并发过大会榨干服务器资源。

Nginx: 采用单线程来异步非阻塞处理请求（管理员可以配置Nginx主进程的工作进程的数量）(epoll)，不会为每个请求分配cpu和内存资源，节省了大量资源，同时也减少了大量的CPU的上下文切换。所以才使得Nginx支持更高的并发。

# Nginx是如何处理一个请求的呢？
    首先，nginx在启动时，会解析配置文件，得到需要监听的端口与ip地址，然后在nginx的master进程里面
    先初始化好这个监控的socket，再进行listen
    然后再fork出多个子进程出来,  子进程会竞争accept新的连接。
此时，客户端就可以向nginx发起连接了。当客户端与nginx进行三次握手，与nginx建立好一个连接后

    此时，某一个子进程会accept成功，然后创建nginx对连接的封装，即ngx_connection_t结构体
    接着，根据事件调用相应的事件处理模块，如http模块与客户端进行数据的交换。

最后，nginx或客户端来主动关掉连接，到此，一个连接就寿终正寝了
