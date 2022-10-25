<!-- ToC start -->
# Table of Contents

1. [计算机早期](#1计算机早期)
1. [电子计算机](#2电子计算机)
1. [晶体管怎么运算的](#3晶体管怎么运算的)
1. [二进制](#二进制)
1. [ALU(算数逻辑单元 arithmetic logic unit) 【如何计算的】](#alu算数逻辑单元-arithmetic-logic-unit-【如何计算的】)
1. [寄存器和内存 【如何存储的】](#寄存器和内存-【如何存储的】)
<!-- ToC end -->



# 1.计算机早期
![](../pic/Lark20201119142708.png)
![](../pic/555.png)
- 算盘
- 莱布尼兹 步进计算器 [减法反向运动，乘除是多个加减]
- 差分机-分析机 查尔斯·巴贝奇 最早的 计算雏形
- 打孔卡片制表机 【一串纸上打孔，已婚的话，通过小孔，给已婚齿轮加一】
# ***********************硬件相关****************************

# 2.电子计算机
- 机械继电器 【机械臂有质量无法快速开关/磨损】
![](../pic/Lark20201119200005.png)
- 热电子管
中间是加热产生负电子，下面只有正电通过发光，负电不可以，【二极管思想】
![](../pic/Lark20201119200617.png)
- 三极管/真空管
> 红色持续-电，+是控制线[+允许电流流过，-和红色负电排斥，禁止流过]
会烧坏
![](../pic/Lark20201119200912.png)
- 晶体管
![](../pic/b225b9a0bfbb451dbe4de0b8df5b1ff4_th.jpeg)
![](../pic/e17eea6db58a4c1b8e263283dafafc23_th.png)
# 3.晶体管怎么运算的
- 电流流过True 1
- 电流断开False 0
- 最开始不是二进制 有五进制 四进制等，因为对01处理的布尔代数非常成熟，而是波形里 0 1更好区分
- 布尔代数里 AND OR NOT XOR
- NOT
下面接地
![](../pic/Lark20201119203936.png)
![img.png](img.png)
- AND【串联】
![](../pic/Lark20201119204207.png)
- OR【并联】
![](../pic/Lark20201119204327.png)
- XOR
  XOR 主要用来判断两个值是否不同。
![](../pic/Lark20201119204549.png)
# 二进制
一个 0 1 就是一个bit 1byte【字节】 = 8bit【比特】
- 表示正负数
- 表示浮点数 1符号+8指数+23位数 32位
- 表示字母 ascii 不适合汉语日语 unicode 16位
# ALU(算数逻辑单元 arithmetic logic unit) 【如何计算的】
1. arithmetic 单元
half adder 半加器 
XOR 主要用来判断两个值是否不同。
![](../pic/Lark20201119210839.png)
full adder 全加器
![](../pic/Lark20201119211130.png)
八位加法器
![](../pic/Lark20201119211258.png)
超前进位加法器 计组知识 
https://zhuanlan.zhihu.com/p/101332501
flag 小于 相等 溢出
![](../pic/Lark20201119212039.png) 
# 寄存器和内存 【如何存储的】
RAM【random access memory】

- AND OR 锁存器 latch memory
分别能存0 和 1
![](../pic/Lark20201119213125.png) 
看存1 还是 0
![](../pic/Lark20201119212817.png) 
- 门锁 latch gate
![](../pic/Lark20201119213349.png) 
- 8个锁存器就是一个寄存器 8位宽寄存器
- 16X16 门锁矩阵寄存器
    1. 对于256的存储 只需要 1[数据线]+1[写入]+1[读取]+16+16 = 35 条线即可
    2. 16X16 即为内存地址 memory address 因为16个最多4位就搞定， 8位表示行列
![](../pic/Lark20201119215233.png) 
![](../pic/Lark20201119215340.png) 
- 多路复用器
![](../pic/Lark20201119220126.png) 
- 八位地址协议 存256位信息
![](../pic/Lark20201119220251.png) 
- 存储字扩展/位扩展 
字扩展是为了存下更多个数的数字
位扩展是为了存下更长位数的数字
![](../pic/Lark20201119222200.png) 
https://blog.csdn.net/m0_37345402/article/details/89818103
我们有256个地址 一个地址能读写一个八位值
![](../pic/Lark20201119222227.png)
# CPU[centeral proccess unit]
- 指令表
前4位是opcode 操作码 后四位是地址或者寄存器
![](../pic/3a7a7977.png)

- 指令地址寄存器【Instruction Register Adderss】追踪指令运行到哪里
- 指令寄存器【Instruction Register】
## cpu指令运行过程

![](../pic/184aff0d.png)
1. fetch phase 取指令阶段
    1. 指令地址寄存器连到RAM，指令寄存器为0，读RAM地址0的指令【00101110】
    2. 指令【00101110】复制到指令寄存器
2. decode phase 解码阶段
    1.【00101110】复制到指令寄存器,0010是loadA操作码，地址1110是14
    2. 指令寄存器左边是控制校验单元，校验0010的码对不对
3. execute phase 执行阶段
    1. 控制校验单元的结果true，打开RAM的允许读线，把14传过去
    2. RAM连到ABCD四个寄存器，控制校验单元的结果true，打开A寄存器的写线，
    3. 把RAM地址14的值【00000011】3 写到A寄存器
    4. 关掉电路，指令地址寄存器+1【预备执行下一个指令】 
![](../pic/ea824dd8.png)
## 控制单元
![](../pic/1f5605c3.png)
- 看一下加法是怎么做的
 1. 操作码是把寄存器A和B想嫁放到寄存器A里面
 2. 寄存器A和B的值是ALU的输入，相加之后，不能直接写入寄存器A，会循环相加
 3. 控制单元【control unit】 有个寄存器暂时存计算结果，关闭ALU
 4. 讲控制单元的和写入寄存器A
 5. 关掉电路，指令地址寄存器+1【预备执行下一个指令】 
![](../pic/6c32146f.png)
## cpu时钟速度【clock speed】
等于 fetch-decode-execute的一个周期时间
以一定频率节奏控制 控制单元自增1
cpu超频过热
![](../pic/2b8ab39a.png)



## intel 4004 1971年的单芯片处理器
![](../pic/0d2192e9.png)

# 指令和程序
结合上面所知，CPU是硬件，可以被软件【RAM】修改
![](../pic/fa6b0445.png)
- jump 是拿后四位的寄存器值覆盖掉指令地址寄存器的值【反复相加算乘法】
- jump_negative ALU为负，更新指令地址寄存器 jump的停止代码 会在两者之间横跳，直至满足条件
- 除法就是反复减，程序7条指令，cpu执行了13次，循环两次，2就是商
- ALU不能做除法，除法是程序做的【不绝对 有硬件做除法】
- 指令长度【instruction length】 32 or 64
- CISC（complex instruction set computer）等长指令集 以及RISC（reduced instruction set computer）变长指令集
![](../pic/f31ca119.png)
![](../pic/b0eaa01d.png)
# 高级CPU设计
- 上文中cpu算多个减法当除法，时钟周期变长，现代cpu给ALU直接设计除法，cpu指令不断扩大，形成指令集
- 复杂cpu指令集加快了时钟周期，但是RAM到cpu的数据传输时间跟不上，cpu会空等空转，
1. 解决方法1，cpu加cache RAM可以一次传一批指令多过来 【cache hit cache miss】
    1. 但是这样cache和RAM不一致，需要dirty bit 做标记
    2. 在cache加载下一页RAM的时候，把dirty bit cache回写RAM，再加载下一页
![](../pic/e62c1b34.png)
2. 解决方法2 指令流水线 instruction pipeling
    1. parallelize 并行 带来的冲突问题
    2. 插入jump等待
    3. 等待的时候不让cpu空转，会预测指令集，推测执行【speculative execution】，错了的话清空再来【pipeline flush】，对了的话大大降低，
    4. 现代cpu有分支预测【branch prediction】
- 多核多cpu都是为了并行处理，共享结果    
    ![](../pic/981a361b.png)
# 早期编程方式 【程序怎么进入RAM】
- 织布机就最早的编程
- 打孔和插拔或开关
# ***********************软件相关****************************
# 编译原理
1. 汇编语言 Assembler 汇编器【编译器】 讲汇编语言转为机器语言 
    1 程序员还是要思考需要什么寄存器和内存地址
2. A语言 
    1. 变量就是内存地址的抽象  
# 数据结构
1. 字符串是最后一位是NULL的字符数组   
2. 数组是地址连续的
3. 多维数据就是 n个一维数组
4. 结构体是可以等同为一类数组
5. 结构体和数组内存在虚拟内存是连续的，不能动态新增or降低
    1. 编程语言的动态数组是 等数组越界，从新申请2倍的内存，copy&delete原来的
6. 没有指针灵活，node是数，next是point 下一个内存地址
7. queue enqueue/dequeue  stack  pop/push
8. 指针里 加一个leftNext，rightNext是就是tree
9. tree的leaf和node 首尾相接，就是图
# 操作系统
1. 因为cpu性能提高，但是I/O时间长，导致CPU空等，为了极速榨干CPU，引入操作系统
2. 上古时期，需要人员人工塞入卡纸，没有调度 【自动加载程序】
3. 不同I/O打印机等不适配，需要重复编程
4. OS提供API操作抽象硬件，这就是设备驱动程序【device drivers】，驱动
5. OS进程间切换需要各自的内存， 虚拟内存【virtual memory】
6. 防止程序间相互访问和恶意程序 需要 memory protection 内存保护
7. OS必备功能：
- 虚拟内存
- 内存保护
- 多任务
- 虚拟内存 动态内存分配
![](../pic/1fcdbcca.png)
# 内存【memory RAM】&& 持久存储【storage】
- 内存快，断电既不可用
- 持久存储慢，断电可用
## storage的发展历史
- 延迟存储器 电-声-声-电传输 最早的存储程序计算机
![](../pic/4cc09762.png)
- magnetic core memory 磁芯存储器
https://zhuanlan.zhihu.com/p/144628785
![](../pic/c2ade94d.png)
- 磁带
![](../pic/66e0c8f3.png)
![](../pic/a80fc3ea.png)
- 磁盘 跟磁带相似 
1. seek time 寻道时间
- 光盘
1. 光盘易读不易写
- 机械硬盘/固态硬盘
1.  SSD是Solid State Drive，即固态硬盘的缩写. 目前主流的SSD是使用半导体闪存（Flash）作为介质的存储设备，SSD有别于HDD(Hard Disk Drive)机械硬盘. 
https://zhuanlan.zhihu.com/p/43362595
# 文件
- wav 音频存储
![](../pic/20c71dcb.png)
- 图片存储
![](../pic/dbba53a6.png)
- 目录
![](../pic/a53f9049.png)
![](../pic/9fe146fe.png)
## 文件分段碎片化
![](../pic/31d9fbb3.png)
## 碎片整理
![](../pic/4a240a3e.png)
同一分区移动文件很快
![](../pic/f76859d9.png)
# 压缩
- 游程编码 run-length-encoding【无损压缩】
![](../pic/cebdbe76.png)
- 压缩音频是为了让更多人打电话
- 视频压缩 帧与帧之间背景一样
# 计算机网络 
## 模拟信号【电信号】怎么转换为数字信号【0101】
https://blog.csdn.net/ck784101777/article/details/103821767
## 局域网【LAN】 Local Area Network AND 以太网 EtherNet
- 网线连起所有PC，但不知道给谁，mac【Media Access Control Address】地址就是身份证
![](../pic/d4d30b71.png)
### 载波侦听多路访问【CSMA】
- 串行mac 传递信息会碰撞等待 加入每个mac加入随机时间 
- exponential backOff 指数退避
- mac和局域网的注册表对应
![](../pic/e75f3e8f.png)
### 交换机【隔离冲突域】 
- 交换机可以让相同和不相同的同时传输


![](../pic/d8103441.png)

## 互联网 【多个以太网组成互联网】 广域网 WAN【Wide Area Network】
- 多个以太网组成互联网，但是需要路由【router】
- 银行军队有自己专门线路，普通的需要 报文交换【message switch】
- 报文交换更更好的容错，有多条线路供选择
- 经过几个路由就是几个跳数【hot count】 分析网络问题 hot limit 限制
![](../pic/cb1a40cf.png)
### 
### 广域网的服务提供商 ISP【InterNet Service Provider】
### 分组交换【Packet Switch】 大报文拆为数据包 
- 网络拥塞机制
- ARPANET 分组交换鼻祖

###【TCP/IP协议】/UDP协议 拆分数据包，不同时序组合问题 
- 数据包在互联网传输需要符合标准 InternalNet Protocol 
- IP 联网的计算机需要IP地址 IP负责把数据送到正确的计算机
- IP 负责把
#### UDP User Datagram Protocol 【用户数据报协议】 负责把数据送到正确的程序
有 IP/Port  checkSum 【校验和】
- 无法重试，只能丢掉 适合IM 视频传输 在线射击游戏
![](../pic/194b0719.png)
#### TCP 传输控制协议【TransMission  Control Protocol】
- 比UDP 多了序列号/重试 确保顺序组合
### DNS 【Domain Name System 】 域名系统
- 互联网服务提供商
#### 树状域名 三级域名
- Top Level Domain  TLD 一级域名 com/cc/cn/org
- Second L D
- 第三级域名

![](../pic/916941aa.png)
### OSI

