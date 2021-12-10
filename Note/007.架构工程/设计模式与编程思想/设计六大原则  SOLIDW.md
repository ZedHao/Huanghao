[toc]
# 设计原则
https://www.cnblogs.com/huangenai/p/6219475.html
## 单一职责原则(Single Responsibility Principle)
对象应该仅具有一种单一职责的概念。

换句话说就是让一个类只做一种类型责任，当这个类需要承担其他类型的责任的时候，就需要分解这个类。
### 例子
  例如支付
  
```
namespace SOLID
{
    public class Users
    {
        /// <summary>
        /// 支付
        /// </summary>
        public void Pay(){}

        /// <summary>
        /// 数据库操作
        /// </summary>
        public void DataAccess(){}

        /// <summary>
        /// 日志操作
        /// </summary>
        public void Logger(){}
    }
}
```
在这个用户类中有这三个功能：1.支付逻辑，2数据库逻辑，3.日志操作。如果将这三个功能结合在一个类中，可能会出现修改部分代码时会破坏其他的部分。多个功能也使这个用户类难以理解，降低了内聚性。所以最好就是将这个类分离为三个分离的类，每个类仅仅有一个功能。
改进 
```
namespace SOLID
{
    /// <summary>
    /// 数据库操作
    /// </summary>
    class DataAccess { }

    /// <summary>
    /// 日志
    /// </summary>
    class Logger { }

    /// <summary>
    /// 支付
    /// </summary>
    class Pay { }
}
```
单一职责原则是实现高内聚、低耦合的指导方针，它是最简单但又最难运用的原则，需要设计人员发现类的不同职责并将其分离，而发现类的多重职责需要设计人员具有较强的分析设计能力和相关实践经验。

例子：手机与照相机，它们都有照相的功能，但是手机的功能不仅仅只有照相还可以打电话，听音乐；但是照相机就只负责照相。因此这里面的手机职能就过多，在维护，等方面就会比单一职责的困难很多。

## 开闭原则（Open Close Principle）
“软件体应该是对于扩展开放的，但是对于修改封闭的”的概念

你不是要变化吗？，那么我就让你继承实现一个对象，用一个接口来抽象你的职责，你变化越多，继承实现的子类就越多。
例如抽象的 数据库操作 是可以 扩展为mysql 或 orcle  mysqlOdp 此思想
### 例子

```
abstract class DataAccess
    {
        public abstract void OpenConnection();
        public abstract void CloseConnection();
        public abstract void ExecuteCommand();
    }

    /// <summary>
    /// SQL
    /// </summary>
    class SqlDataAccess : DataAccess
    {
        /// <summary>
        /// 打开SQL数据库
        /// </summary>
        public override void OpenConnection(){}
        /// <summary>
        /// 关闭Sql数据连接
        /// </summary>
        public override void CloseConnection(){}
        /// <summary>
        /// 执行Sql数据命令
        /// </summary>
        public override void ExecuteCommand(){}
    }
    
    /// <summary>
    /// ORACLE
    /// </summary>
    class OracleDataAccess : DataAccess
    {
        /// <summary>
        /// 打开Oracle数据连接
        /// </summary>
        public override void OpenConnection(){}
        /// <summary>
        /// 关闭Oracle数据连接
        /// </summary>
        public override void CloseConnection(){}
        /// <summary>
        /// 执行Oracle数据命令
        /// </summary>
        public override void ExecuteCommand(){}
    }
```

## 里氏代换原则 （Liskov Substitution Principle）
“程序中的对象应该是可以在不改变程序正确性的前提下被它的子类所替换的”
https://www.cnblogs.com/chenxkang/p/6657384.html
例子：企鹅是一种特殊的鸟，但是企鹅不会飞，若是在鸟这类中添加，飞这个行为，，那么企鹅将不能继承这个类。

## 接口隔离原则 （Interface Segregation Principle）
“多个特定客户端接口要好于一个宽泛用途的接口”
==不能强迫用户去依赖那些他们不使用的接口==。
### 背景和例子
 换句话说，使用多个专门的接口比使用单一的总接口总要好。注意：在代码中应用ISP并不一定意味着服务就是绝对安全的。仍然需要采用良好的编码实践，以确保正确的验证与授权。
这个原则起源于施乐公司，他们需要建立了一个新的打印机系统，可以执行诸如装订的印刷品一套，传真多种任务。此系统软件创建从底层开始编制，并实现了这些 任务功能，但是不断增长的软件功能却使软件本身越来越难适应变化和维护。每一次改变，即使是最小的变化，有人可能需要近一个小时的重新编译和重新部署。这 是几乎不可能再继续发展，所以他们聘请罗伯特Robert帮助他们。他们首先设计了一个主要类Job,几乎能够用于实现所有任务功能。只要调用Job类的 一个方法就可以实现一个功能，Job类就变动非常大，是一个胖模型啊，对于客户端如果只需要一个打印功能，但是其他无关打印的方法功能也和其耦合，ISP 原则建议在客户端和Job类之间增加一个接口层，对于不同功能有不同接口，比如打印功能就是Print接口，然后将大的Job类切分为继承不同接口的子 类，这样有一个Print Job类，等等。
```
interface IDataAccess
    {
        void OpenConnection();
        void CloseConnection();
    }

    interface ISqlDataAccess : IDataAccess
    {
        void ExecuteSqlCommand();
    }
    interface IOracleDataAccess : IDataAccess
    {
        void ExecuteOracleCommand();
    }
    class SqlDataAccess : ISqlDataAccess
    {
        /// <summary>
        /// 执行Sql数据命令
        /// </summary>
        public void ExecuteSqlCommand(){}

        /// <summary>
        /// 打开Sql数据连接
        /// </summary>
        public void OpenConnection(){}

        /// <summary>
        /// 关闭Sql数据连接
        /// </summary>
        public void CloseConnection(){}
    }
    class OracleDataAccess : IOracleDataAccess
    {
        /// <summary>
        /// 执行Oracle数据命令
        /// </summary>
        public void ExecuteOracleCommand(){}

        /// <summary>
        /// 打开Oracle数据连接
        /// </summary>
        public void OpenConnection(){}

        /// <summary>
        /// 关闭Oracle数据连接
        /// </summary>
        public void CloseConnection(){}
    }
```


## 依赖倒转原则  Dependency Inversion Principle（解耦必备原则）
我们是会用的电脑，cpu，内存等是分开的，当有那个地方出问题的时候，只需要将出问题的零件替换掉，但是想一想收音机，里面错综复杂，每个零件之间相互依赖，也许碰到一个其他的也会因此出问题。

依赖于抽象而不是一个实例

依赖反转原则： 依赖反转原则 认为一个方法应该遵从“依赖于抽象而不是一个实例” 的概念。依赖注入是该原则的一种实现方式。

依赖倒置原则(Dependency Inversion Principle，DIP)规定：代码应当取决于抽象概念，而不是具体实现。
==高层模块不应该依赖于低层模块==，二者都应该依赖于抽象 
抽象不应该依赖于细节，细节应该依赖于抽象 

类可能依赖于其他类来执行其工作。但是，它们不应当依赖于该类的特定具体实现，而应当是它的抽象。这个原则实在是太重要了，社会的分工化，标准化都 是这个设计原则的体现。显然，这一概念会大大提高系统的灵活性。如果类只关心它们用于支持特定契约而不是特定类型的组件，就可以快速而轻松地修改这些低级 服务的功能，同时最大限度地降低对系统其余部分的影响。

```
interface IBankAccount
    {
        long BankNumber { get; set; } // 卡号
        decimal Balance { get; set; } // 余额
    }

    // 转账人
    interface ITransferSource : IBankAccount
    {
        void CutPayment(decimal value);
    }

    // 收款人
    interface ITransferDestination : IBankAccount
    {
        void AddMoney(decimal value);
    }

    class BankAccout : IBankAccount, ITransferSource, ITransferDestination
    {
        public long BankNumber { get; set; }
        public decimal Balance { get; set; }
        public void CutPayment(decimal value)
        {
            Balance -= value;
        }
        public void AddMoney(decimal value)
        {
            Balance += value;
        }
    }
   
    class TransferAmount
    {
        public decimal Amount { get; set; }
        public void Transfer(ITransferSource source, ITransferDestination dest)
        {
            source.CutPayment(Amount);
            dest.AddMoney(Amount);
        }
    }
```


## 迪米特法则
迪米特法则(Law of Demeter )又叫做最少知识原则，也就是说，一个对象应当对其他对象尽可能少的了解。不和陌生人说话。英文简写为: LoD。

目的

　　在于降低类与类之间的耦合。由于每个类尽量减少对其他类的依赖，因此，很容易使得系统的功能模块功能独立，是的相互间存在尽可能少的依赖关系。

优点

　　迪米特法则的做法观念就是类间解耦，弱耦合，只有弱耦合了以后，类的复用率才可以提高。

缺点

  　  造成系统的不同模块之间的通信效率降低，使系统的不同模块之间不容易协调等缺点。

 　　因为迪米特法则要求类与类之间尽量不直接通信，如果类之间需要通信就通过第三方转发的方式，这就直接导致了系统中存在大量的中介类，大大增加了系统的复杂度。

　　解决这个问题的方式是：使用依赖倒转原则（通俗的讲就是要针对接口编程，不要针对具体编程）， 这要就可以是调用方和被调用方之间有了一个抽象层，被调用方在遵循抽象层的前提下就可以自由的变化，此时抽象层成了调用方的朋友。

    第一：在类的划分上，应当创建弱耦合的类，类与类之间的耦合越弱，就越有利于实现可复用的目标。
    第二：在类的结构设计上，每个类都应该降低成员的访问权限。
    第三：在类的设计上，只要有可能，一个类应当设计成不变的类。
    第四：在对其他类的引用上，一个对象对其他类的对象的引用应该降到最低。
    第五：尽量限制局部变量的有效范围，降低类的访问权限。

 简单示例：

　　上课时，教师让学习委员进行点名
　　
```
namespace DesignPrinciples.LowOfDemeter
{
    class Program
    {
        static void Main(string[] args)
        {
            GroupLeader group = new GroupLeader();
            Teacher teacher = new Teacher();
            teacher.Command(group);
            Console.WriteLine("==== 我是分割线 ====");
            teacher.CommandByDemeter(group);
        }
    }

    public class Teacher
    { 
        /// <summary>
        /// 教师发送命令
        /// </summary>
        /// <param name="group"></param>
        public void Command(GroupLeader group)
        {
            List<Student> students = new List<Student>();
            for (int i = 0; i < 20; i++)
            {
                students.Add(new Student());
            }
            group.Count(students);
        }

        public void CommandByDemeter(GroupLeader group)
        { 
            group.CountByDemeter();
        }
    }

    public class GroupLeader
    {
        /// <summary>
        /// 学习委员点名
        /// </summary>
        /// <param name="girls"></param>
        public void Count(List<Student> girls)
        {
            Console.WriteLine("学生人数：" + girls.Count);
        }

        public void CountByDemeter()
        {
            List<Student> students = new List<Student>();
            for (int i = 0; i < 20; i++)
            {
                students.Add(new Student());
            }
            Console.WriteLine("学生人数（迪米特）：" + students.Count);
        }
    }

    public class Student
    {
        public string Name { get; set; }
    }
}
```
