# http的历史
![image](http://mmbiz.qpic.cn/mmbiz_png/cmOLumrNib1cfBOtIMQ6JfSibJdd6QkQribgQuEeJaevuN9LRgQ9WR85hRiaVISeia7SDz1aU9hAAgO33XFaJ3FhmhQ/0?wx_fmt=png)

1. HTTP1.0需要使用keep-alive参数来告知服务器端要建立一个长连接，而HTTP1.1默认支持长连接
1. 节约带宽：
 HTTP1.1支持只发送header信息（不带任何body信息），如果服务器认为客户端有权限请求服务器，则返回100，否则返回401。客户端如果接受到100，才把body发送到服务器。当服务器返回401的时候，节约了带宽。
1. HOST域：
HTTP1.0没有HOST域
1. 多路复用：
 HTTP2.0使用多路复用的技术，做到同一个连接并发处理多个请求，并且并发请求的数量比HTTP1.1大了好几个数量级
1. 数据压缩：
HTTP1.1不支持header数据的压缩，HTTP2.0使用HPACK算法对header的数据进行压缩，这样数据体积小了，在网络上传输就会更快。
1. 服务器推送：
在HTTP2.0中服务器可以对客户端的一个请求发送多个响应，并且可以缓存。

