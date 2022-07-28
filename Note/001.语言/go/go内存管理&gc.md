https://juejin.cn/post/6844903795739082760#heading-5
https://segmentfault.com/a/1190000020338427
https://www.infoq.cn/article/IEhRLwmmIM7-11RYaLHR
# heap 堆和 stack 栈 的差别
1. heap 自下向上 越来越大 stack 自下向上 越来越小
2. heap是OS自动分配 但是因为链表容易产生碎片 释放要花cpu stack c++ 需要程序员自己free del
![img.png](img.png)
# TMalloc
Golang运行时的内存分配算法主要源自 Google 为 C 语言开发的TCMalloc算法，全称Thread-Caching Malloc。核心思想就是把内存分为多级管理，从而降低锁的粒度。它将可用的堆内存采用二级分配的方式进行管理：每个线程都会自行维护一个独立的内存池，进行内存分配时优先从该内存池中分配，当内存池不足时才会向全局内存池申请，以避免不同线程对全局内存池的频繁竞争。
![img_2.png](img_2.png)
# 快速分配内存的三个层次
1. 引入虚拟内存后，让内存的并发访问问题的粒度从多进程级别，降低到多线程级别。
2. 为线程预分配缓存需要进行1次系统调用，后续线程申请小内存时，从缓存分配，都是在用户态执行，没有系统调用，缩短了内存总体的分配和释放时间，这是快速分配内存的第二个层次。
3. 多个线程同时申请小内存时，从各自的缓存分配，访问的是不同的地址空间，无需加锁，把内存并发访问的粒度进一步降低了，这是快速分配内存的第三个层次。

# go 的内存管理 说的其实是虚拟内存
1. OS的线程切换需要上下文切换 CPU 比较耗时
2. golang的goroutine 是语言层面 切换耗时短
3. 参考了 TCMalloc【Thread Cache Malloc】 Thread Process Center
4. 分配内存块，回收内存块和组织内存块。
![img_1.png](img_1.png)
# gc 垃圾回收
