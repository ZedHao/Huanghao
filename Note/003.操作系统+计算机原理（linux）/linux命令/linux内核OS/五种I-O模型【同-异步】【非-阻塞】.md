[TOC]
# [同异步，非阻塞]概念的理解： 
首先在网络编程的时候，我们常常见到同步/异步，阻塞/非阻塞四中调用方式：

## 同步和异步主要针对C（client）端 
==同步：== 

所谓的同步，就是在c端发出一个功能调用时，在没有得到结果之前，该调用步返回，也就是说必须一件一件事做，等前一件事完了之后才做后一件事。 
如：普通的B/S模式（同步）：提交请求->等待服务器处理->处理完毕返回，这期间客户端浏览器不能干任何事

==异步：==

与同步相对。当C端一个异步过程调用发出之后，调用者不能立即得到结果，实际处理这个调用的部件在完成后，通过状态，通知和回调来通知调用者。 
如：请求通过事件触发->服务器处理（浏览器仍然可以做其他事情）->处理完毕
## 阻塞和非阻塞主要针对S端（server）
==阻塞==

==阻塞调用是指调用结果返回之前，当前线程会被挂起（线程进入非可执行状态，在这个状态，cpu不会分配时间片，线程暂停运行）函数只有得到结果返回==

阻塞调用和同步调用的区别：对同步来说，很多时候当前线程还是激活的，只是逻辑上没有返回，如，在socket编程中调用recv函数，如果缓冲区没有数据，这个函数就会一直等待，直到有数据返回。而此前当前线程还有可能继续处理各种各样的消息。

阻塞的例子：比如去取A楼一层（假设是内核缓冲区）取快递，但是比不知道什么时候来，你有不能干别的事情，只能死等着但是可以睡觉（进程处于休眠状态)，因为你知道快递把货送来时一定会给比大电话

非阻塞： 
==非阻塞与阻塞概念想对应，指在不能立即得到结果之前，该函数不会阻塞当前线程，而会立即返回。==

非阻塞的例子：还是等快递，如果用轮询的方式，每隔5分钟去A楼一层（内核缓冲区）去看快递来了没，没来，立即返回，如果快递来了，就放到A楼一层，等你去取。

==对象是否处于阻塞模式和函数是不是阻塞调用有很强的相关性，但不是一一对应的。阻塞对象上可以有非阻塞的调用方式，我们可以通过轮询状态，在适当的时候调用阻塞函数，就可以避免阻塞，而对于非阻塞对象，调用函数可以进入阻塞调用==

对于select： 
## 同步异步阻塞非阻塞经典解释
1：同步 
我客户端（C端调用者）一个功能，该功能没有结束前，我死等结果。

2：异步，我（c端调用者）调用一个功能，不知道该功能结果，该功能有结果后通知我，即回调通知
==同步和异步主要针对c端，但是跟s端不是完全没关系，同步和异步必须s端配合才能实现，同步和异步由c端控制，但是s端是否为阻塞还是非阻塞，c端不关心==

3：阻塞，就是调用我（s端被调用者，函数），我（s端被调用者，函数）没有完全接受完数据或者没有得到结果之前，我不会返回。

4：非阻塞，就是调用我（s端被调用者，函数），我（s端被调用者，函数）立即返回，通过select通知调用者

## 同步I/O和异步I/O 
==同步I/O与异步I/O的区别在与数据访问的时候进程是否阻塞==
==阻塞I/O与非阻塞I/O的区别在与：应该程序的调用是否立即返回==

阻塞和非阻塞是指server端的进程访问的数据如果尚未就绪，进程是否需要等待，简单说这相当于函数内部的实现区别，也就是未就绪时时直接返回还是等待就绪。

就同步和异步是指client端访问数据的机制，同步一般指主动请求并等待I/O操作完毕的方式，当数据就绪后再读写额时候必须阻塞，异步则指主动请求数据后便可以继续处理其他任务，随后等待I/O，操作完毕的通知。


# Linux——一切皆文件

==Linux将所有外部设备都看做文件，对文件的读写操作会调用内核提供的系统命令，返回一个file descriptor(fd，文件描述符)。==

Linux也把socket当成文件，称为socketfd(socket描述符).

描述符是一个数字，该数字指向内核中的一个结构体（该结构体包含文件路径，数据区等属性）
# 五种I/O模型【将数据从内核空间拷贝到用户空间】
- 1)阻塞I/O 
- 2)非阻塞I/O 
- 3)I/O复用(select和poll) 
- 4)信号驱动I/O（SIGIO） 
- 5)异步I/O 

==前四中是同步，最后一种是异步。== 
#
1.操作系统是一个运行在硬件之上的软件系统，操作系统控制IO（磁盘，网络（socket）），cpu，内存的绝对控制权限，这些操作都是在内核空间完成的。

2.操作系统对外提供一些接口，其他接口可以调用这些接口控制外设，而用户的这些代码运行在用户空间里面，而这里里面的一些指令不能操作内核控制的一些设备或者内存区间。

## 阻塞I/O模型【请求等待应答】

在进程空间中调用recvfrom，该系统调用直到数据包到达且被复制到应用进程的缓冲区中或者发生错误时才返回，次期间将会一直等待。

进程从调用recvfrom方法开始到recvfrom方法返回的期间都是被阻塞的。

当调用recv()函数时，系统首先检查是否有准备好的数据，如果数据没有准备好，那么系统就处于等待状态，当数据准备好后，将数据从系统缓冲区复制到用户空间，然后函数返回。在套接应用程序中，当调用recv()函数时，未必用户空间就已经存在数据，那么此时recv()函数处于等待状态 

![image](https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2534342328,3421389422&fm=173&s=E6F1E07E0B426C4F5AD455CA0000E0B2&w=380&h=241&img.JPEG)

## 非阻塞I/O模型(轮询式)【间隔请求应答（true or false）】：

调用recvfrom方法后，如果内核该缓冲区没有数据，就直接返回一个EWOULDBLOCK错误。

我们把一个套接口设置为非阻塞就是告诉内存，当所请求的I/O操作无法完成时，不要惊进程睡眠，而是返回一个错误，河阳I/O函数会不断的测试数据是否准备好，没有准备好，继续测试，直到数据准备好为止。==在测试的过程中会占用大量的CPU时间。 ==

![image](https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2285714997,3468784484&fm=173&s=3BAC702201584DCA5ED555CA0000E0B5&w=640&h=360&img.JPEG)

## I/O复用模型
I/O复用模型会调用select，poll函数，这几个函数也会使进程阻塞，但是和阻塞I/O不同的，这个函数可以同时阻塞多个I/O操作，而且可以同时对多个读操作，多个写操作的I/O函数进行检测，直到有数据可读或可写时，才真正调用I/O操作函数。

### (select/poll)【扫描式】
 进程通过将一定数量的fd传递给select/poll函数。

select函数会被阻塞，select/poll函数侦测多个fd是否处于就绪状态。

select/poll是顺序扫描fd是否就绪的，而且支持的fd数量是有限。
### I/O复用模型(epoll)

Linux提供了一个epoll函数，epoll函数使用事件驱动方式代替顺序扫描，因此性能更高。当有fd就绪时，立即调用回调函数。

![image](https://ss2.baidu.com/6ONYsjip0QIZ8tyhnq/it/u=1111632650,1470429700&fm=173&s=2F2C70221F0A404918DCF1CA0000C0B5&w=640&h=360&img.JPEG)

## 信号驱动I/O模型【两次调用，两次返回 e内核主动通知】
简介：两次调用，两次返回 
首先允许套接口进行信号驱动I/O，并安装一个信号处理函数，进程继续运行并不阻塞。昂数据准备好时，进程会收到一个SIGIO信号，可以在信号处理函数中调用I/O操作函数处理数据。 
先开启套接口的信号驱动I/O功能，sigaction函数安装一个信号处理函数（sigaction函数立即返回，进程继续工作，它是非阻塞的）。

当数据准备就绪时，就为该进程生成一个SIGIO信号，通过信号回调通知应用程序调用recvfrom来读取数据。

先开启套接口的信号驱动I/O功能，sigaction函数安装一个信号处理函数（sigaction函数立即返回，进程继续工作，它是非阻塞的）。

当数据准备就绪时，就为该进程生成一个SIGIO信号，通过信号回调通知应用程序调用recvfrom来读取数据。

![image](https://ss1.baidu.com/6ONXsjip0QIZ8tyhnq/it/u=1373078707,4254225799&fm=173&s=7B2C3062835E79CA58FD55CA0000C0B1&w=640&h=360&img.JPEG)

## 异步I/O

调用系统函数通知内核启动某个操作，并让内核在整个操作完成后（包括将数据从内核复制到用户自己的缓冲区）通知我们。

![image](https://img-blog.csdn.net/20170905135446181?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvZjIwMTY5MTM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


# Linux I/O模型的总结

IO主要涉及两步操作：

==等待内核数据准备完成数据从内核复制到用户空间
I/O模式的改进就是不断减少在这两步所花的时间。==

、

阻塞I/O: 两个阶段阻塞

非阻塞I/O:待内核数据准备完成期间轮询（等同阻塞，也可以多个I/O一起轮询） ,数据从内核复制到用户空间阻塞

I/O复用:等待内核数据准备完成期间可以多个I/O轮询，数据从内核复制到用户空间阻塞

信号驱动I/O：数据从内核复制到用户空间期间线程阻塞

异步I/O：完全不阻塞

Linux信号驱动I/O与异步I/O的区别

信号驱动I/O：内核通知程序何时可以开始一个I/O操作；

异步I/O： 内核通知程序I/O操作何时已完成,也就是说数据准备和数据复制到用户空间都是系统完成的。