[TOC]
# php的理解
PHP是一种适用于web开发的动态语言。具体点说，就是一个用C语言实现包含大量组件的软件框架。更狭义点看，可以把它认为是一个强大的UI框架。
## PHP代码的执行过程: 
![image](https://images2018.cnblogs.com/blog/592892/201804/592892-20180414110927385-1112518846.png)
- Parse。语法分析阶段。
- Compile。编译产出opcode中间码。
- Execute。运行，动态运行进行输出。
    由于PHP是个解释型语言执行的时候先得把程序读进来，然后由Zend引擎编译成opcode。最后Zend虚拟机顺次执行这些opcode（指令）完成操作。因此我们可 以把这个Opcode缓存起来，下次就能避免重新编译了。
### 编译型语言: 
对于C语言,C++, 编译成机器码(二进制)来运行. 
java语言, 把.java编译成 .class,称为bytecode, 由jvm来运行. 
### 解释语言: 
解释器解释执行. 典型的如 linux shell. 
解释器逐行来执行命令.

PHP稍有特殊之处,虽然是一个脚本语言,但不是靠解释器解释. 
==而是 zend 虚拟机,屏蔽了操作系统的区别. ==
php代码编译成opcode, 由zend虚拟机来执行opcode.

![image](https://img-blog.csdn.net/20161203142152200)

PHP变量的底层实现: 
我们解压PHP的源码包, 看到如下的目录 

 ![image](https://img-blog.csdn.net/20161203142245233)
其中, 
最核心的—Zend目录, 这是zend虚拟的实现. 包括栈,数据类型,编译器等,都在这实现. 
最主要的main –PHP的一些内建函数,最主要函数都在这里放着. 
最大的一个目录 ext – PHP的扩展.

PHP的大部分功能,都是以extenstion形式来完成的. 
如果你开发了一个扩展,也放在ext目录下.

# PHP中有8种数据类型,为什么zval->value 联合体中,只有5种? 
答: 
1: NULL,直接 zval->type = IS_NULL,就可以表示,不必设置 value的值. 
2: BOOL型 , zval->type = IS_BOOL, 再设置 zval.value.lval = 1/0; 
3: Resourc型 ,资源型 往往是服务器上打开的一个接口,如果 文件读取接口. 
zval->type = IS_RESOURCE, zval->tyoe.lval = 服务器上打开的接口的编号


# PHP的设计理念及特点
1. 多进程模型：由于PHP是多进程模型，不同请求间互不干涉，这样保证了一个请求挂掉不会对全盘服务造成影响，当然，随着时代发展，PHP也早已支持多线程模型。
1. 弱类型语言：和C/C++、Java、C#等语言不同，PHP是一门弱类型语言。一个变量的类型并不是一开始就确定不变，运行中才会确定并可能发生隐式或显式的类型转换，这种机制的灵活性在web开发中非常方便、高效，具体会在后面PHP变量中详述。
1. 引擎(Zend)+组件(ext)的模式降低内部耦合。
1. 中间层(sapi)隔绝web server和PHP。
1. 语法简单灵活，没有太多规范。缺点导致风格混杂，但再差的程序员也不会写出太离谱危害全局的程序。
# php 四层体系
![image](http://ww2.sinaimg.cn/large/7cc829d3gw1exyap8boh7j20ee0ep0tj.jpg)
## Zend引擎【编译为可执行的opcode】：
2. Zend整体用纯C实现，是PHP的内核部分，它将PHP代码翻译（词法、语法解析等一系列编译过程）为可执行opcode的处理并实现相应的处理方法、实现了基本的数据结构（如hashtable、oo）、内存分配及管理、提供了相应的api方法供外部调用，是一切的核心，所有的外围功能均围绕Zend实现。
## Extensions【内置函数等基础组建扩展】：
围绕着Zend引擎，extensions通过组件式的方式提供各种基础服务，我们常见的各种内置函数（如array系列）、标准库等都是通过extension来实现，用户也可以根据需要实现自己的extension以达到功能扩展、性能优化等目的（如贴吧正在使用的PHP中间层、富文本解析就是extension的典型应用）。
##  Sapi：Server Application Programming Interface
  也就是服务端应用编程接口，Sapi通过一系列钩子函数，使得PHP可以和外围交互数据，这是PHP非常优雅和成功的一个设计，通过sapi成功的将PHP本身和上层应用解耦隔离，PHP可以不再考虑如何针对不同应用进行兼容，而应用本身也可以针对自己的特点实现不同的处理方式。
如前所述，Sapi通过通过一系列的接口，使得外部应用可以和PHP交换数据并可以根据不同应用特点实现特定的处理方法，我们常见的一些sapi有：

### apache2handler【apache】：
这是以apache作为webserver，采用mod_PHP模式运行时候的处理方式，也是现在应用最广泛的一种。
### cgi【ngnix】：
这是webserver和PHP直接的另一种交互方式，也就是大名鼎鼎的fastcgi协议，在最近今年fastcgi+PHP得到越来越多的应用，也是异步webserver所唯一支持的方式。
cli：命令行调用的应用模式
### cli：命令行调用的应用模式 php -a
## 上层应用：
这就是我们平时编写的PHP程序，通过不同的sapi方式得到各种各样的应用模式，如通过webserver实现web应用、在命令行下以脚本方式运行等等

# PHP的opcache
## 什么是opcode缓存？

当解释器完成对脚本代码的分析后，便将它们生成可以直接运行的中间代码，也称为操作码

（Operate Code，opcode）。Opcode cache的目地是避免重复编译，减少CPU和内存开销。

## 为什么要使用Opcode缓存？
本来的php生命周期
![image](https://img-blog.csdn.net/20140120175042625)
加了OpCode缓存后
![image](https://img-blog.csdn.net/20140120175158765)

## 常用插件Optimizer+ eAccelerator、xcache、APC 