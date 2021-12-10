[TOC]
https://www.awaimai.com/patterns
# UML类图详解

## 类结构
在类的UML图中，使用长方形描述一个类的主要构成，长方形垂直地分为三层，以此放置类的名称、属性和方法。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/structure.png)

##  类关系
类与类之间的关系主要有六种：继承、实现、组合、聚合、关联和依赖，这六种关系的箭头表示
![image](https://www.awaimai.com/wp-content/uploads/2016/07/arrow.png)
## 六种类关系
六种类关系中，组合、聚合、关联这三种类关系的代码结构一样，都是用属性来保存另一个类的引用，所以要通过内容间的关系来区别。
###  继承
继承关系也称泛化关系（Generalization），用于描述父类与子类之间的关系。==父类又称作基类，子类又称作派生类。==

==继承关系中，子类继承父类的所有功能，父类所具有的属性、方法，子类应该都有==。子类中除了与父类一致的信息以外，还包括额外的信息。

![image](https://www.awaimai.com/wp-content/uploads/2016/07/realization.png)

### 实现
3.2 实现
实现关系（Implementation），主要用来==规定接口和实现类的关系==。

接口（包括抽象类）是方法的集合，在实现关系中，类实现了接口，类中的方法实现了接口声明的所有方法。

例如：汽车和轮船都是交通工具，而交通工具只是一个可移动工具的抽象概念，船和车实现了具体移动的功能。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/implement.png)
### 组合关系
组合关系（Composition）：==整体与部分的关系，但是整体与部分不可以分开。==

组合关系表示类之间整体与部分的关系，整体和部分有一致的生存期。一旦整体对象不存在，部分对象也将不存在，是同生共死的关系。

例如：人由头部和身体组成，两者不可分割，共同存在。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/composition.png)

### 聚合关系
聚合关系（Aggregation）：整体和部分的关系，整体与部分可以分开。

聚合关系也表示类之间整体与部分的关系，成员对象是整体对象的一部分，但是成员对象可以脱离整体对象独立存在。

例如：公交车司机和工衣、工帽是整体与部分的关系，但是可以分开，工衣、工帽可以穿在别的司机身上，公交司机也可以穿别的工衣、工帽。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/aggregation.png)

### 关联关系
关联关系（Association）：表示一个类的属性保存了对另一个类的一个实例（或多个实例）的引用。

关联关系是类与类之间最常用的一种关系，表示一类对象与另一类对象之间有联系。组合、聚合也属于关联关系，只是关联关系的类间关系比其他两种要弱。

关联关系有四种：双向关联、单向关联、自关联、多重数关联。

例如：汽车和司机，一辆汽车对应特定的司机，一个司机也可以开多辆车。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/association.png)
### 依赖关系
依赖关系（Dependence）：假设A类的变化引起了B类的变化，则说名B类依赖于A类。

大多数情况下，依赖关系体现在某个类的方法使用另一个类的对象作为参数。

依赖关系是一种“使用”关系，特定事物的改变有可能会影响到使用该事物的其他事物，在需要表示一个事物使用另一个事物时使用依赖关系。

例如：汽车依赖汽油，如果没有汽油，汽车将无法行驶。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/dependence.png)

## 总结
这六种类关系中，组合、聚合和关联的代码结构一样，可以从关系的强弱来理解，各类关系从强到弱依次是：继承→实现→组合→聚合→关联→依赖。如下是完整的一张UML关系图。
![image](https://www.awaimai.com/wp-content/uploads/2016/07/whole.png)
