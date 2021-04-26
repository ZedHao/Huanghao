[toc]
# 数据库安全
## php 连接mysql的三种方式 


### PHP的MySQL扩展  
太古老 高版本不支持
本扩展自 PHP 5.5.0 起已废弃，并在将来会被移除），PHP原生的方式去连接数据库，是面向过程的
```
<?php
$mysql_conf = array(
    'host'    => '127.0.0.1:3306', 
    'db'      => 'test', 
    'db_user' => 'root', 
    'db_pwd'  => 'root', 
    );
$mysql_conn = @mysql_connect($mysql_conf['host'], $mysql_conf['db_user'], $mysql_conf['db_pwd']);
if (!$mysql_conn) {
    die("could not connect to the database:\n" . mysql_error());//诊断连接错误
}
mysql_query("set names 'utf8'");//编码转化
$select_db = mysql_select_db($mysql_conf['db']);
if (!$select_db) {
    die("could not connect to the db:\n" .  mysql_error());
}
$sql = "select * from user;";
$res = mysql_query($sql);
if (!$res) {
    die("could get the res:\n" . mysql_error());
}

while ($row = mysql_fetch_assoc($res)) {
    print_r($row);
}

mysql_close($mysql_conn);
?>
```

### PHP的mysqli扩展
优点：面向对象接口 、prepared语句支持、多语句执行支持、事务支持 、增强的调试能力、嵌入式服务支持 、预处理方式完全解决了sql注入的问题。
缺点：就是只支持mysql数据库。如果你要是不操作其他的数据库

```
<?php
$mysql_conf = array(
    'host'    => '127.0.0.1:3306', 
    'db'      => 'test', 
    'db_user' => 'root', 
    'db_pwd'  => 'joshua317', 
    );

$mysqli = @new mysqli($mysql_conf['host'], $mysql_conf['db_user'], $mysql_conf['db_pwd']);
if ($mysqli->connect_errno) {
    die("could not connect to the database:\n" . $mysqli->connect_error);//诊断连接错误
}
$mysqli->query("set names 'utf8';");//编码转化
$select_db = $mysqli->select_db($mysql_conf['db']);
if (!$select_db) {
    die("could not connect to the db:\n" .  $mysqli->error);
}$sql = "select uid from user where name = 'joshua';";
$res = $mysqli->query($sql);
if (!$res) {
    die("sql error:\n" . $mysqli->error);
}
 while ($row = $res->fetch_assoc()) {
        var_dump($row);
    }

$res->free();
$mysqli->close();
?>
```

### PHP数据对象(PDO) 【功能类似于JDBC、ODBC、DBI】 解决了sql注入的问题

```
<?php
$mysql_conf = array(
    'host'    => '127.0.0.1:3306', 
    'db'      => 'test', 
    'db_user' => 'root', 
    'db_pwd'  => 'joshua317', 
    );
$pdo = new PDO("mysql:host=" . $mysql_conf['host'] . ";dbname=" . $mysql_conf['db'], $mysql_conf['db_user'], $mysql_conf['db_pwd']);//创建一个pdo对象
$pdo->exec("set names 'utf8'");
$sql = "select * from user where name = ?";
$stmt = $pdo->prepare($sql);
$stmt->bindValue(1, 'joshua', PDO::PARAM_STR);
$rs = $stmt->execute();
if ($rs) {
    // PDO::FETCH_ASSOC 关联数组形式
    // PDO::FETCH_NUM 数字索引数组形式
    while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
        var_dump($row);
    }
}

$pdo = null;//关闭连接
?>
```

## 数据库安全采取的手段  
### 预防SQL 注入
#### 在服务器端配置
(1) 打开php的安全模式
但是默认的php.ini是没有打开安全模式的，我们把它打开：

safe_mode = on


(2) 用户组安全

当safe_mode打开时，safe_mode_gid被关闭，那么php脚本能够对文件进行访问，而且相同
safe_mode_gid = off

(3) 安全模式下执行程序主目录

(4) 安全模式下包含文件

如果要在安全模式下包含某些公共文件，那么就修改一下选项：

safe_mode_include_dir = D:/usr/www/include/



(5) 控制php脚本能访问的目录

使用open_basedir选项能够控制PHP脚本只能访问指定的目录，这样能够避免PHP脚本访问

不应该访问的文件，一定程度上限制了phpshell的危害，我们一般可以设置为只能访问网站目录：

open_basedir = D:/usr/www

(6) 关闭危险函数


如果你要禁止任何文件和目录的操作，那么可以关闭很多文件操作

disable_functions = chdir,chroot,dir,getcwd,opendir,readdir,scandir,fopen,unlink,delete,copy,mkdir, rmdir,rename,file,file_get_contents,fputs,fwrite,chgrp,chmod,chown

以上只是列了部分不叫常用的文件处理函数，你也可以把上面执行命令函数和这个函数结合，

就能够抵制大部分的phpshell了。

(7) 关闭PHP版本信息在http头中的泄漏

我们为了防止黑客获取服务器中php版本的信息，可以关闭该信息斜路在http头中：

expose_php = Off


(8) 关闭注册全局变量

在PHP中提交的变量，包括使用POST或者GET提交的变量，都将自动注册为全局变量，能够直接访问，

这是对服务器非常不安全的，所以我们不能让它注册为全局变量，就把注册全局变量选项关闭：

register_globals = Off

(9) 打开magic_quotes_gpc来防止SQL注入

SQL注入是非常危险的问题，小则网站后台被入侵，重则整个服务器沦陷，

所以一定要小心。php.ini中有一个设置：

magic_quotes_gpc = Off

这个默认是关闭的，如果它打开后将自动把用户提交对sql的查询进行转换，

比如把 ' 转为 \'等，这对防止sql注射有重大作用。所以我们推荐设置为：

magic_quotes_gpc = On

(10) 错误信息控制+错误日志
####  php代码层面   addslashes() 和 str_replace()  mysql_escape_string()  magic_quotes_gpc 打开

中文防止SQL注入的检查。addslashes的问题在于黑客可以用0xbf27来代替单引
* addslashes() 是强行加\；
* mysql_real_escape_string()  会判断字符集，但是对PHP版本有要求；
* mysql_escape_string不考虑连接的当前字符集
### magic_quotes_gpc选项打开，在这种情况下所有的客户端GET和POST的数据都会自动进行addslashes处理
#### 散列加密  addslashes() == pg_escape_string MD5

```
<?php

// 存储密码散列  
$query  = sprintf("INSERT INTO users(name,pwd) VALUES('%s','%s');",
            pg_escape_string($username), md5($password));
$result = pg_query($connection, $query);

// 发送请求来验证用户密码
$query = sprintf("SELECT 1 FROM users WHERE name='%s' AND pwd='%s';",
            pg_escape_string($username), md5($password));
$result = pg_query($connection, $query);

if (pg_num_rows($result) > 0) {
    echo 'Welcome, $username!';
} else {
    echo 'Authentication failed for $username.';
}

?>
```
### 预防xxs(跨站脚本攻击(Cross Site Scripting)) 
#### 概念
![image](https://images2015.cnblogs.com/blog/1114374/201707/1114374-20170704220225331-622725464.png)
XSS是一种经常出现在web应用中的计算机安全漏洞，它允许恶意web用户将代码植入到提供给其它用户使用的页面中。比如这些代码包括HTML代码和客户端脚本
 HTML 注入，攻击者的输入没有经过后台的过滤直接进入到数据库，最终显示给来访的用户。
#### 分类
##### 反射型【非持久型】 如弹窗
　　反射型XSS，也叫非持久型XSS，是指发生请求时，XSS代码出现在请求URL中，作为参数提交到服务器，服务器解析并响应。响应结果中包含XSS代码，最后浏览器解析并执行。从概念上可以看出，反射型XSS代码是首先出现在URL中的，然后需要服务端解析，最后需要浏览器解析之后XSS代码才能够攻击。
##### 存储型【持久型】 如盗身份　
####  第三方跳转

#### 解决方案 
php防止XSS跨站脚本攻击的方法:是针对非法的HTML代码包括单双引号等，使用htmlspecialchars()函数 。
　　1、入参字符过滤
　　2、出参进行编码
　　3、入参长度限制
　　4、设置cookie httponly为true

### 预防csrf(（Cross-site request forgery）跨站请求伪造)
#### 概念
![image](https://images2015.cnblogs.com/blog/1114374/201707/1114374-20170704221329987-250861756.png)
登录受信任网站A，并在本地生成Cookie。
在不登出A的情况下，访问危险网站B。
看到这里，你也许会说：“如果我不满足以上两个条件中的一个，我就不会受到CSRF的攻击”。是的，确实如此，但你不能保证以下情况不会发生：
1.你不能保证你登录了一个网站后，不再打开一个tab页面并访问另外的网站。
2.你不能保证你关闭浏览器了后，你本地的Cookie立刻过期，你上次的会话已经结束。（事实上，关闭浏览器不能结束一个会话，但大多数人都会错误的认为关闭浏览器就等于退出登录/结束会话了......）
3.上图中所谓的攻击网站，可能是一个存在其他漏洞的可信任的经常被人访问的网站。

#### 解决方案
客户端页面增加伪随机数。
(1).Cookie Hashing(所有表单都包含同一个伪随机值)：
    这可能是最简单的解决方案了，因为攻击者不能获得第三方的Cookie
    
(2).验证码
　　这个方案的思路是：每次的用户提交都需要用户在表单中填写一个图片上的随机字符串

(3).One-Time Tokens(不同的表单包含一个不同的伪随机值)

(4).在 HTTP 头中自定义属性并验证

(5).在请求地址中添加 token 并验证

(6).验证 HTTP Referer 字段
### 预防Ddos攻击



# 以 CGI 模式安装 可能受到的攻击 解决办法
## 情形一：只运行公开的文件  --enable-force-cgi-redirect
## 情形二：使用 --enable-force-cgi-redirect 选项
## 情形三：设置 doc_root 或 user_dir
## 情形四：PHP 解释器放在 web 目录以外

# 文件系统安全 Null 字符问题
想要删除自己主目录中的一个文件。假设此情形是通过 web 界面来管理文件系统，因此 Apache 用户有权删除用户目录下的文件。

```
<?php
// 从用户目录中删除指定的文件
$username = $_POST['user_submitted_name'];
$userfile = $_POST['user_submitted_filename'];
$homedir = "/home/$username";
unlink ("$homedir/$userfile");
echo "The file has been deleted!";
?>
```

# 错误报告（利用log日志干坏事 php.ini）
1. 第一个是彻底地检查所有函数，并尝试弥补大多数错误。
1. 第二个是对在线系统彻底关闭错误报告。
1. 第三个是使用 PHP 自定义的错误处理函数创建自己的错误处理机制。根据不同的安全策略，三种方法可能都适用。

一个能提前阻止这个问题发生的方法就是利用 error_reporting() 来帮助使代码更安全并发现变量使用的危险之处。在发布程序之前，先打开 E_ALL 测试代码，可以帮你很快找到变量使用不当的地方。一旦准备正式发布，就应该把 error_reporting() 的参数设为 0 来彻底关闭错误报告或者把 php.ini 中的 display_errors 设为 off 来关闭所有的错误显示以将代码隔绝于探测。当然，如果要迟一些再这样做，就不要忘记打开 ini 文件内的 log_errors 选项，并通过 error_log 指定用于记录错误信息的文件。

# 用户提交的数据 校验（防SQL注入和XSS攻击等） addslashes() 
PHP 对所有的 GET、POST 和 COOKIE 数据自动运行 addslashes()。所以您不应对已转义过的字符串使用 addslashes()，因为这样会导致双层转义。遇到这种情况时可以使用函数 get_magic_quotes_gpc() 进行检测。

# 隐藏 PHP
在 php.ini 文件里设置 expose_php = off ，可以减少他们能获得的有用信息。

另一个策略就是让 web 服务器用 PHP 解析不同扩展名。无论是通过 .htaccess 文件还是 Apache 的配置文件，都可以设置能误导攻击者的文件扩展名：