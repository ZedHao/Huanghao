查询分类
[TOC]
# 简单查询
# 联合查询（UNION）
## 一、UNION（消除重复行）和UNION ALL（不消除重复行）

```
UNION 用于合并两个或多个 SELECT 语句的结果集，并消去表中任何重复行。
```

```
select staff_uid,role from tblStaff  union select staff_uid,role from tblStaffInitInfo;
```

1. +------------+------+
1. | staff_uid  | role |
1. +------------+------+
1. | 2135343000 |    1 |
1. |    1234567 |    1 |
1. |    1850040 |    4 |
1. |    8879041 |    0 |
1. | 2290116733 |    1 |
1. |    8879041 |    4 |
1. | 2290116733 |    4 |
1. +------------+------+

# 连接查询
## 内连接 (只连接匹配的行)
### 等值连接

```
select tblStaff.*,tblStaffBaseInfo.* from tblStaff inner join tblStaffBaseInfo on tblStaff.staff_uid = tblStaffBaseInfo.staff_uid\G;
```

### 自然连接（select 确定值）
自然连接：基于两个表的同名的一个或多个列。

【注意】自然连接是根据两个表中同名的列而进行连接的，当列不同名时，自然连接将失去意义。且语法中没有on。

### 不等连接（>）


## 外连接
使用关键字outer join。用于检索一个表的所有记录和另一个表中的匹配行。
### 左外连接 (left outer join 无限满足左表)

```
select  tblStaff.staff_uid,tblStaffBaseInfo.staff_uid,staff_name,class  from tblStaff  left outer  join tblStaffBaseInfo   on tblStaff.staff_uid = tblStaffBaseInfo.staff_uid;
```

### 右外连接 (right outer join 无限满足右表)


### 全外连接 FULL OUTER JOIN

## 交叉连接
念:没有where条件的交叉连接将产生连接表所涉及的笛卡尔积。即TableA的行数*TableB的行数的结果集。

