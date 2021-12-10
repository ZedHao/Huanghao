# MySQL服务端实用工具：
1. mysqld：SQL的后台程序（MySQL服务器进程）。改程序必须运行之后，客户端才能连接服务器访问数据库。

1. mysqld_safe:服务器启动脚本，在Unix推荐使用mysqld_safe来启动mysqld服务器。mysqld_safe增加了一些安全特性。
1. msq.server:服务器启动脚本，用于使用包含为特定级别的、运行启动服务的脚本的、运行目录的系统。它调用mysqld_safe来启动MySQL服务器。
1. mysqld_multi:服务器启动脚本，可以启动或者停止系统上安装的多个服务器。
1. myisamchk：用来描述、检查、优化和维护MyISAM表的实用工具。


1. mysqlbug：MySQL缺陷报告脚本。他可以用来向MySQL邮件系统发送缺陷报告。
1. mysql_install_db：该脚本用默认权限创建MySQL授权表，通常只有在系统上首次安装MySQL执行一次。
2. 
1. MySQL客户端实用工具：
1. myisampack：压缩MyISAM表以产生一个更小的只读表的一个工具。
1. mysql：交互式输入SQL语句或从文件以批处理模式执行他们的命令行工具。
1. mysqlaccess：检查主机名、用户名和数据库组合的权限脚本。
2. 
1. mysqladmin:执行管理操作的客户程序，例如创建删除数据库，重载授权表，将表刷新到硬盘，以及重新打开日志文件。mysqladmin还可以用来检索版本、进程，以及服务器的状态信息。
1. mysqlbinlog：从二进制日志读取语句的工具。在二进制日志文件中包含执行过的语句，可用来帮助系统从崩溃中恢复。
1. mysqlcheck：检查、修复、分析以及优化表的表维护客户程序
2. 
1. mysqldump：将MySQL数据库转储到一个文件的客户程序。
1. mysqlhotcopy：当服务器运行中，快速备份MyISAM或ISAM表的工具。
1. mysql import：使用LOAD DATA INFILE将文本文件导入相关表的客户程序。
2. 
1. mysqlshow：显示数据库、表、列、索引相关信息的客户程序。
1. perror：显示系统或者MySQL错误代码含义的工具。