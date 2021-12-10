[toc]
# context的选项和参数

## 套接字上下文选项 — 套接字上下文选项列表 工作在套接字上的封装协议，像 tcp, http 和 ftp.


```
<?php
// connect to the internet using the '192.168.0.100' IP
$opts = array(
    'socket' => array(
        'bindto' => '192.168.0.100:0',
    ),
);


// connect to the internet using the '192.168.0.100' IP and port '7000'
$opts = array(
    'socket' => array(
        'bindto' => '192.168.0.100:7000',
    ),
);


// connect to the internet using port '7000'
$opts = array(
    'socket' => array(
        'bindto' => '0:7000',
    ),
);


// create the context...
$context = stream_context_create($opts);

// ...and use it to fetch the data
echo file_get_contents('http://www.example.com', false, $context);

?>
```


## HTTP context 选项 — HTTP context 的选项列表 file_get_contents  stream_context_create
## FTP context options — FTP context option listing
## SSL 上下文选项 — SSL 上下文选项清单
## CURL context options — CURL 上下文选项列表
## Phar 上下文（context）选项 — Phar 上下文（context）选项 类似jar 打包文件
## MongoDB context options — MongoDB context option listing
## Context 参数 — Context 参数列表
# 支持的协议
1. file:// — 访问本地文件系统
1. http:// — 访问 HTTP(s) 网址
1. ftp:// — 访问 FTP(s) URLs
1. php:// — 访问各个输入/输出流（I/O streams）
1. zlib:// — 压缩流
1. data:// — 数据（RFC 2397）
1. glob:// — 查找匹配的文件路径模式
1. phar:// — PHP 归档
1. ssh2:// — Secure Shell 2
1. rar:// — RAR
1. ogg:// — 音频流
1. expect:// — 处理交互式的流