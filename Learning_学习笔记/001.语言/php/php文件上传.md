# 一般文件上传的原理
## 一.概述
 Struts/Spring MVC 文件上传实现都是基于Commons-fileupload，但背后的原理，大多数估计没有关注，最近阅读Commons-fileupload源码也发现，只有基础才是最重要的，万变不离其宗，在it领域不然会被漫天的新技术，冲昏了头，不知所措，下面开始。
## 二.HTTP报文
1.要想理解文件上传，先要对HTTP报文有一个基本的了解
2.HTTP报文是简单的格式化数据块，每条报文都包含来自客户端（如浏览器）的请求，或者来自服务器的响应
3.HTTP报文由3部分组成
起始行                  - 对报文的描述，所有的HTTP报文都以一个起始行作为开始，即报文的第一行就是起始行
首部                     - 可以有0或多个首部字段，每个首部 字段都包含一个名字和一个值，首部字段后是\r\n，首部以一个空行\r\n结束
实体的主体部分   - 首部恐婚之后就是可选的报文实体，实体的主体是HTTP报文的负荷，传输的内容就在实体的主体中

![image](https://img-blog.csdn.net/20170904123625516?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGlwaW5nYW5x/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

![image](https://img-blog.csdn.net/20170905214936101?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGlwaW5nYW5x/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

![image](https://img-blog.csdn.net/20170905215154085?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGlwaW5nYW5x/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/Center)

##  TCP
1.TCP为HTTP提供了一条可靠的比特传输管道。从TCP一端填入的字节会从另一端以原有的顺序、正确的传送出来，TCP提供了：

无差错的数据传输
按序传输 - 数据总是按照发送的顺序到达
未分段的数据流 - 可以在任何时刻以任意尺寸将数据发送出去
2.HTTP要传送一条报文时，会以流的形式将报文数据的内容通过一条打开的TCP连接按序传输，TCP收到数据流后，会将数据流砍成被称作段的小数据块，并将段封装在IP分组中，通过因特网传输。所有这些工作都是由TCP/IP软件完成的，HTTP程序员什么也看不见

3.只要建立了TCP连接，客户端与服务器之间的报文交换就不会丢失，不会被破坏，也不会在接收时出现错序了。
# php的文件上传
　　将客户端的文件上传到服务器，再将服务器的临时文件上传到指定目录