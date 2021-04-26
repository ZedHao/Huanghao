[TOC]
# php.ini文件
# 内存 
php.ini文件中的memory_limit设置单个php进程可以使用的系统内存最大值
分配原则
1.一共能分给php多少内存？（ngnix，MYSQL，memcache也要耗内存）
2.单个php进程平均耗多少？
3.能负担多少PHP-FPM内存？
# zend OPcache（操作码缓存）
# 文件上传
# 最长执行时间(max_execution_time)
# 处理会话（缓存存会话 可扩展）
# 缓存输出
# 测试百科
分类：
## 单元测试（PHPUnit）搭配Xdebug（代码覆盖率）

## 功能测试

# 开发模式
## 测试驱动开发（TDD）（小步向前）
## 行为驱动开发（BDD）
# 使用Travis CI持续测试（自动化非常赞）

