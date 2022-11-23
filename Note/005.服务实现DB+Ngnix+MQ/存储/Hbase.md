[toc]
# 概念
## 和hive，RDBMS区别
![image](B92B006942DD4457BE3813CD8D862C73)
hive是离线大数据分析，实时查询慢
hbase大数据实时查询

## 存储模式
![image](83A20B4653234E098BF9C1D61A07D49D)
 Hbase 中，Row-key 加上 CF 加上 Qulifier 再加上一个时间戳才可以定位到一个单元格数据（Hbase 中每个单元格默认有 3 个时间戳的版本数据）
 ![image](7B30AE3569344E6EA3B078505608C612)
逻辑存储顺序和逻辑顺序
 
## 使用场景
1. 专为海量数据计算设计
2. 缺少RDBMS特性【辅助索引，触发器，高级查询语言】
3. 高速聚合计算
## 架构体系
https://www.ibm.com/developerworks/cn/analytics/library/ba-cn-bigdata-hbase/index.html
![image](A1AA2667B259462D99BA8FEB29CAF14E)
![image](F49E15CC909F40B1BEC72755B17B960F)