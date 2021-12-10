[toc]
# 垃圾回收机制GC[Garbage Collection]
## 变量结构

```
typedef union _zvalue_value {  
    long lval;                  /* long value */ 
    double dval;                /* double value */ 
    struct {  
        char *val;  
        int len;  
    } str;  
    HashTable *ht;              /* hash table value */ 
    zend_object_value obj;  
} zvalue_value;  
 
struct _zval_struct {  
    /* Variable information 值 */ 
    zvalue_value value;       
/* value */ 
    zend_uint refcount__gc;  //计数器，表示指向这个zval变量容器的变量个数。 
    zend_uchar type;    /* active type 类型 */ 
    zend_uchar is_ref__gc;//是个bool值，用来区分变量是否属于引用集合。
}; 

```
## 非数组存储例子

```
<?php
$a = 1;
xdebug_debug_zval('a');
echo PHP_EOL;
$b = $a;
xdebug_debug_zval('a');
echo PHP_EOL;
 
$c = &$a;
xdebug_debug_zval('a');
echo PHP_EOL;
 
xdebug_debug_zval('b');
echo PHP_EOL;
?>
```
　　运行结果如下：

a:(refcount=1, is_ref=0),int 1

a:(refcount=2, is_ref=0),int 1

a:(refcount=2, is_ref=1),int 1

b:(refcount=1, is_ref=0),int 1
## 数组存储例子
![image](http://php.net/manual/zh/images/12f37b1c6963c1c5c18f30495416a197-simple-array.png)
数组分配了三个zval容器：a   meaning  number
## 垃圾回收算法
### Reference Counting 算法 5.2版本 环状引用内存泄漏

```
<?php
$a = array( 'one' );
$a[] =& $a;
xdebug_debug_zval( 'a' );
?>
```
a 和 1 的zval容器 是一样的。如下：
![image](http://php.net/manual/zh/images/12f37b1c6963c1c5c18f30495416a197-loop-array.png)

在5.2及更早版本的PHP中，没有专门的垃圾回收器GC（Garbage Collection），引擎在判断一个变量空间是否能够被释放的时候是依据这个变量的zval的refcount的值，如果refcount为0，那么变量的空间可以被释放，否则就不释放，这是一种非常简单的GC实现。

现在unset ($a),那么array的refcount减1变为1.现在无任何变量指向这个zval，而且这个zval的计数器为1，不会回收。 成千上万的这种就会内存泄漏

![image](http://php.net/manual/zh/images/12f37b1c6963c1c5c18f30495416a197-leak-array.png)
### Concurrent Cycle Collection in Reference Counted Systems 算法 5.3版本 解决了环状引用,控制内存泄漏的阙值

在php5.3的GC中，针对的垃圾做了如下说明：

1：如果一个zval的refcount增加，那么此zval还在使用，肯定不是垃圾，不会进入缓冲区

2：如果一个zval的refcount减少到0， 那么zval会被立即释放掉，不属于GC要处理的垃圾对象，不会进入缓冲区。

 3：如果一个zval的refcount减少之后大于0，那么此zval还不能被释放，此zval可能成为一个垃圾，将其放入缓冲区。PHP5.3中的GC针对的就是这种zval进行的处理。

开启/关闭垃圾回收机制可以通过修改php配置实现，也可以在程序中使用gc_enable() 和 gc_disable()开启和关闭。

# 用 PHP 进行 HTTP 认证
以用 header() 函数来向客户端浏览器发送“Authentication Required”信息，使其弹出一个用户名／密码输入窗口。当用户输入用户名和密码后，包含有 URL 的 PHP 脚本将会加上预定义变量 PHP_AUTH_USER，PHP_AUTH_PW 和 AUTH_TYPE 被再次调用，这三个变量分别被设定为用户名，密码和认证类型。预定义变量保存在 $_SERVER 数组中。支持“Basic”和“Digest”（自 PHP 5.1.0 起）认证方法。

```
$_SERVER['PHP_AUTH_USER'] = "test";
$_SERVER['PHP_AUTH_PW'] = "123456";
//*********************************
//用户账号
$login = 'test'; 
$pass = '1234567'; 
//*********************************

if(($_SERVER['PHP_AUTH_PW']!= $pass || $_SERVER['PHP_AUTH_USER'] != $login)|| !$_SERVER['PHP_AUTH_USER'])
 { 
    header('WWW-Authenticate: Basic realm="Test auth"'); 
    header('HTTP/1.0 401 Unauthorized'); 
    echo 'Auth failed'; 
    exit; 
} else{
    echo "Auth success";
}
具体业务具体对待.....
```
# cookie和session
 setcookie() 和 setrawcookie() 函数
 
# 处理 XForms 表单
# 文件上传处理 post +put

# 使用远程文件
只要在 php.ini 文件中激活了 allow_url_fopen 选项，就可以在大多数需要用文件名作为参数的函数中使用 HTTP 和 FTP 的 URL 来代替文件名。

# 数据库持久连接
# 安全模式 强级别校验
# PHP 的命令行模式 php -a


# DTrace 动态跟踪 性能调试
