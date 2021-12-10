<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [设计模式](#%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)
  - [行为型模式](#%E8%A1%8C%E4%B8%BA%E5%9E%8B%E6%A8%A1%E5%BC%8F)
    - [观察者模式](#%E8%A7%82%E5%AF%9F%E8%80%85%E6%A8%A1%E5%BC%8F)
    - [策略模式](#%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F)
    - [迭代器模式](#%E8%BF%AD%E4%BB%A3%E5%99%A8%E6%A8%A1%E5%BC%8F)
    - [命令模式](#%E5%91%BD%E4%BB%A4%E6%A8%A1%E5%BC%8F)
  - [创建型模式](#%E5%88%9B%E5%BB%BA%E5%9E%8B%E6%A8%A1%E5%BC%8F)
    - [简单工厂模式](#%E7%AE%80%E5%8D%95%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)
    - [工厂模式](#%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)
    - [抽象工厂模式](#%E6%8A%BD%E8%B1%A1%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)
    - [单例模式](#%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F)
  - [结构型模式](#%E7%BB%93%E6%9E%84%E5%9E%8B%E6%A8%A1%E5%BC%8F)
    - [适配器模式](#%E9%80%82%E9%85%8D%E5%99%A8%E6%A8%A1%E5%BC%8F)
    - [组合模式](#%E7%BB%84%E5%90%88%E6%A8%A1%E5%BC%8F)
    - [外观模式](#%E5%A4%96%E8%A7%82%E6%A8%A1%E5%BC%8F)
    - [代理模式](#%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

[TOC]

# 设计模式
## 行为型模式
### 观察者模式
eg：当订单创建后，系统会发送邮件和短信，并保存日志记录。
在观察者模式中，被观察者完全不需要关心观察者，在自身状态有变化是，遍历执行观察者update()方法即完成通知。
https://www.awaimai.com/patterns/observer

```

<?php
/**
 * 被观察者接口
 */
interface Observable
{
    // 添加/注册观察者
    public function attach(Observer $observer);
    // 删除观察者
    public function detach(Observer $observer);
    // 触发通知
    public function notify();
}

/**
 * 被观察者
 * 职责：添加观察者到$observers属性中，
 * 有变动时通过notify()方法运行通知
 */
class Order implements Observable
{
    // 保存观察者
    private $observers = array();
    // 订单状态
    private $state = 0;

    // 添加（注册）观察者
    public function attach(Observer $observer)
    {
        $key = array_search($observer, $this->observers);
        if ($key === false) {
            $this->observers[] = $observer;
        }
    }

    // 移除观察者
    public function detach(Observer $observer)
    {
        $key = array_search($observer, $this->observers);
        if ($key !== false) {
            unset($this->observers[$key]);
        }
    }

    // 遍历调用观察者的update()方法进行通知，不关心其具体实现方式
    public function notify()
    {
        foreach ($this->observers as $observer) {
            // 把本类对象传给观察者，以便观察者获取当前类对象的信息
            $observer->update($this);
        }
    }

    // 订单状态有变化时发送通知
    public function addOrder()
    {
        $this->state = 1;
        $this->notify();
    }

    // 获取提供给观察者的状态
    public function getState()
    {
        return $this->state;
    }
}
/**
 * 观察者接口
 */
interface Observer
{
    // 接收到通知的处理方法
    public function update(Observable $observable);
}

/**
 * 观察者1：发送邮件
 */
class Email implements Observer
{
    public function update(Observable $observable)
    {
        $state = $observable->getState();
        if ($state) {
            echo '发送邮件：您已经成功下单。';
        } else {
            echo '发送邮件：下单失败，请重试。';
        }
    }
}

/**
 * 观察者2：短信通知
 */
class Message implements Observer
{
    public function update(Observable $observable)
    {
        $state = $observable->getState();
        if ($state) {
            echo '短信通知：您已下单成功。';
        } else {
            echo '短信通知：下单失败，请重试。';
        }
    }
}

/**
 * 观察者3：记录日志
 */
class Log implements Observer
{
    public function update(Observable $observable)
    {
        echo '记录日志：生成了一个订单记录。';
    }
}
// 创建观察者对象
$email = new Email();
$message = new Message();
$log = new Log();
// 创建订单对象
$order = new Order();

// 向订单对象中注册3个观察者：发送邮件、短信通知、记录日志
$order->attach($email);
$order->attach($message);
$order->attach($log);
// 添加订单，添加时会自动发送通知给观察者
$order->addOrder();

echo '<br />';

// 删除记录日志观察者
$order->detach($log);
// 添加另一个订单，会再次发送通知给观察着
$order->addOrder();

```

### 策略模式
策略模式定义了一族相同类型的算法，算法之间独立封装，并且可以互换代替。

这些算法是同一类型问题的多种处理方式，他们具体行为有差别。

每一个算法、或说每一种处理方式称为一个策略。

在应用中，就可以根据环境的不同，选择不同的策略来处理问题。

 
```
//那如何用策略模式解决这个问题呢？

//策略模式将各种方案分离开来，让操作者根据具体的需求，动态地选择不同的策略方案。

//2.1 策略类
//首先，定义一系列的策略类，它们独立封装，并且遵循统一的接口。

/**
 * 策略接口
 */
interface OutputStrategy
{
    public function render($array);
}

/**
 * 策略类1：返回序列化字符串
 */
class SerializeStrategy implements OutputStrategy
{
    public function render($array)
    {
        return serialize($array);
    }
}

/**
 * 策略类2：返回JSON编码后的字符串
 */
class JsonStrategy implements OutputStrategy
{
    public function render($array)
    {
        return json_encode($array);
    }
}

/**
 * 策略类3：直接返回数组
 */
class ArrayStrategy implements OutputStrategy
{
    public function render($array)
    {
        return $array;
    }
}
///以后的维护过程中，以上代码都不需修改了。

//如果需要增加输出方式，重新建一个类就可以了。

///（根据FIG-PSR规范，一个类就是一个独立的PHP文件。）

//2.2 环境类
//环境角色用来管理策略，实现不同策略的切换功能。

//同样，一旦写好，环境角色类以后也不需要修改了。

/**
 * 环境角色类
 */
class Output
{
    private $outputStrategy;

    // 传入的参数必须是策略接口的子类或子类的实例
    public function __construct(OutputStrategy $outputStrategy)
    {
        $this->outputStrategy = $outputStrategy;
    }

    public function renderOutput($array)
    {
        return $this->outputStrategy->render($array);
    }
}
//2.3 客户端代码
//在客户端中，策略模式通过给予不同的具体策略，来获取不同的结果。
//
/**
 * 客户端代码
 */
$test = ['a', 'b', 'c'];

// 需要返回数组
$output = new Output(new ArrayStrategy());
$data = $output->renderOutput($test);
var_dump($data);
// 需要返回JSON
$output = new Output(new JsonStrategy());
$data = $output->renderOutput($test);
var_dump($data);
//对于较为复杂的业务逻辑显得更为直观，扩展也更为方便。
```


以数组输出为例。

数组的输出有序列化输出、JSON字符串输出和数组格式输出等方式。



每种输出方式都可以独立封装起来，作为一个策略。

应用时，如要把数组保存到数据库中，可以用序列化方式输出。

要提供给APP作接口，可以用JSON字符串输出。

其他程序调用，则直接输出数组格式。
### 迭代器模式
不关心内部实现  遍历对象

 1.访问一个聚合对象的内容而无需暴露它的内部表示

 2.支持对聚合对象的多种遍历

 3.为遍历不同的聚合结构提供一个统一的接口

```
<?php
/**
 * Created by PhpStorm.
 * User: Jiang
 * Date: 2015/6/8
 * Time: 21:31
 */
 
//抽象迭代器
abstract class IIterator
{
    public abstract function First();
    public abstract function Next();
    public abstract function IsDone();
    public abstract function CurrentItem();
}
 
//具体迭代器
class ConcreteIterator extends IIterator
{
    private $aggre;
    private $current = 0;
    public function __construct(array $_aggre)
    {
        $this->aggre = $_aggre;
    }
    //返回第一个
    public function First()
    {
        return $this->aggre[0];
    }
 
    //返回下一个
    public function  Next()
    {
        $this->current++;
        if($this->current<count($this->aggre))
        {
            return $this->aggre[$this->current];
        }
        return false;
    }
 
    //返回是否IsDone
    public function IsDone()
    {
        return $this->current>=count($this->aggre)?true:false;
    }
 
    //返回当前聚集对象
    public function CurrentItem()
    {
        return $this->aggre[$this->current];
    }
}

header("Content-Type:text/html;charset=utf-8");
//--------------------------迭代器模式-------------------
//require_once "./Iterator/Iterator.php";
$iterator= new ConcreteIterator(array('周杰伦','王菲','周润发'));
$item = $iterator->First();
echo $item."<br/>";
while(!$iterator->IsDone())
{
    echo "{$iterator->CurrentItem()}：请买票！<br/>";
    $iterator->Next();

}

```

### 命令模式

这个模式主要由 命令类、用户请求数据类、业务逻辑类、命令类工厂类及调用类构成，各个类的作用概括如下：

1、命令类：调用用户请求数据类和业务逻辑类；

2、用户请求数据类：获取用户请求数据及保存后台处理后返回的结果；

3、业务逻辑类：如以下的示例中验证用户登陆信息是否正确的功能等；

4、命令工厂类(我自己取的名字,哈哈)：生成命令类的实例，这个类第一次看的时候我觉得有点屌，当然看了几遍了还是觉得很屌 ：）；

5、调用类：调用命令类，生成视图；
```
//命令类
abstract class Command {
　　abstract function execute(CommandContext $context);
}
 
class LoginCommand extends Command{　　　　　　　//处理用户登陆信息的命令类
　　function execute (CommandCotext $context){　　　　//CommandCotext 是一个处理用户请求数据和后台回馈数据的类
　　　　$manager = Registry::getAccessManager();　　//原文代码中并没有具体的实现，但说明了这是一个处理用户登陆信息的业务逻辑类
　　　　$user = $context->get('username');
　　　　$pass = $context->get('pass');
　　　　$user_obj = $manager->login($user,$pass);
　　　　if(is_null($user_obj)){
　　　　　　$context->setError($manager->getError);
　　　　　　return false;
　　　　}
　　　　$context->addParam('user',$user_obj);
　　　　return true;　　　　　　　　　　　　　　　//用户登陆成功返回true
　　}
}
 
class FeedbackCommand extends Command{　　　　　　　　//发送邮件的命令类
　　function execute(CommandContext $context){
　　　　$msgSystem = Registry::getMessageSystem();
　　　　$email = $context->get('email');
　　　　$msg = $context->get('msg');
　　　　$topic = $context->get('topci');
　　　　$result = $msgSystem->send($email,$msg,$topic);
　　　　if(!$result){
　　　　　　$context->setError($msgSystem->getError());
　　　　　　return false;
　　　　}
　　　　return true;
　　}
}
 
//用户请求数据类  
class CommandContext {
　　private $params = array();
　　private $error = '';
 
　　function __construct (){
　　$this->params = $_REQUEST;
}
 
function addParam($key,$val){
　　$this->params[$key] = $val;
}
 
function get($key){
　　return $this->params[$key];
}
 
function setError($error){
　　$this->error = $error;
}
 
function getError(){
　　return $this->error;
}
}
 
 
//命令类工厂,这个类根据用户请求数据中的action来生成命令类
class CommandNotFoundException extends Exception {}
 
class CommandFactory {
　　private static $dir = 'commands';
 
　　static function getCommand($action='Default'){
　　　　if(preg_match('/\w',$action)){
　　　　　　throw new Exception("illegal characters in action");
　　　　}
　　　　$class = UCFirst(strtolower($action))."Command";
　　　　$file = self::$dir.DIRECTORY_SEPARATOR."{$class}.php"; //DIRECTORY_SEPARATOR代表'/',这是一个命令类文件的路径
　　　　if(!file_exists($file)){
　　　　　　throw new CommandNotFoundException("could not find '$file'");
　　　　}
　　　　require_once($file);
　　　　if(!class_exists($class)){
　　　　　　throw new CommandNotFoundException("no '$class' class located");
　　　　}
　　　　$cmd = new $class();
　　　　return $cmd;
　　}
}
 
//调用者类,相当于一个司令部它统筹所有的资源
class Controller{
　　private $context;
　　function __construct(){
　　　　$this->context = new CommandContext();  //用户请求数据
　　}
　　function getContext(){
　　　　return $this->context;
　　}
 
　　function process(){
　　　　$cmd = CommandFactory::getCommand($this->context->get('action'));　　　　//通过命令工厂类来获取命令类
　　　　if(!$comd->execute($this->context)){　　　　　　　　　　　　　　　　　　　　　　
　　　　　　//处理失败
　　　　} else {
　　　　　　//成功
　　　　　　// 分发视图
　　　　}
　　}
}
 
// 客户端
$controller = new Controller();
//伪造用户请求,真实的场景中这些参数应该是通过post或get的方式获取的，貌似又废话了：）
$context = $controller->getContext();
$context->addParam('action','login');
$context->addParam('username','bob');
$context->addParam('pass','tiddles');
$controller->process();
```

## 创建型模式
https://www.cnblogs.com/mingaixin/p/4324156.html
### 简单工厂模式
PHP工厂模式概念：工厂模式是一种类，它具有为您创建对象的某些方法。您可以使用工厂类创建对象，而不直接使用 new。这样，如果您想要更改所创建的对象类型，只需更改该工厂即可。使用该工厂的所有代码会自动更改。
根据抽象程度不同，PHP工厂模式分为：简单工厂模式、工厂方法模式和抽象工厂模式
```
/**
 *简单工厂模式与工厂方法模式比较。
 *简单工厂又叫静态工厂方法模式，这样理解可以确定，简单工厂模式是通过一个静态方法创建对象的。 
 */
interface  people {
    function  jiehun();
}
class man implements people{
    function jiehun() {
        echo '送玫瑰，送戒指！<br>';
    }
}
 
class women implements people {
    function jiehun() {
        echo '穿婚纱！<br>';
    }
}
 
class SimpleFactoty {
    // 简单工厂里的静态方法
    static function createMan() {
        return new     man;
    }
    static function createWomen() {
        return new     women;
    }
    
}
 
$man = SimpleFactoty::createMan();
$man->jiehun();
$man = SimpleFactoty::createWomen();
$man->jiehun();
```
### 工厂模式
工厂模式是一种类，它具有为您创建对象的某些方法。您可以使用工厂类创建对象，而不直接使用 new。这样，如果您想要更改所创建的对象类型，只需更改该工厂即可。使用该工厂的所有代码会自动更改



```
<?php
/*
 *工厂方法模式：
 *定义一个创建对象的接口，让子类决定哪个类实例化。 他可以解决简单工厂模式中的封闭开放原则问题。<www.phpddt.com整理>
 */
interface  people {
    function  jiehun();
}
class man implements people{
    function jiehun() {
        echo '送玫瑰，送戒指！<br>';
    }
}
 
class women implements people {
    function jiehun() {
        echo '穿婚纱！<br>';
    }
}
 
interface  createMan {  // 注意了，这里是简单工厂本质区别所在，将对象的创建抽象成一个接口。
    function create();
 
}
class FactoryMan implements createMan{
    function create() {
        return  new man;
    }
}
class FactoryWomen implements createMan {
    function create() {
        return new women;
    }
}
 
class  Client {
    // 工厂方法
    function test() {
        $Factory =  new  FactoryMan;
        $man = $Factory->create();
        $man->jiehun();
        
        $Factory =  new  FactoryWomen;
        $man = $Factory->create();
        $man->jiehun();
    }
}
 
$f = new Client;
$f->test();
```

### 抽象工厂模式
工厂模式 ：用来生产同一等级结构中的固定产品。（支持增加任意产品）   
抽象工厂 ：用来生产不同产品族的全部产品。（对于增加新的产品，无能为力；支持增加产品族）

```
<?php
/*
抽象工厂：提供一个创建一系列相关或相互依赖对象的接口。 
注意：这里和工厂方法的区别是：一系列，而工厂方法则是一个。
那么，我们是否就可以想到在接口create里再增加创建“一系列”对象的方法呢？
*/
interface  people {
    function  jiehun();
}
class Oman implements people{
    function jiehun() {
        echo '美女，我送你玫瑰和戒指！<br>';
    }
}
class Iman implements people{
    function jiehun() {
        echo '我偷偷喜欢你<br>';
    }
}
 
class Owomen implements people {
    function jiehun() {
        echo '我要穿婚纱！<br>';
    }
}
 
class Iwomen implements people {
    function jiehun() {
        echo '我好害羞哦！！<br>';
    }
}
 
interface  createMan {  // 注意了，这里是本质区别所在，将对象的创建抽象成一个接口。
    function createOpen(); //分为 内敛的和外向的
    function createIntro(); //内向
 
}
class FactoryMan implements createMan{
    function createOpen() {
        return  new  Oman;
    }
    function createIntro() {
        return  new Iman;
    }
}
class FactoryWomen implements createMan {
    function createOpen() {
        return  new  Owomen;
    }
    function createIntro() {
        return  new Iwomen;
    }
}
 
class  Client {
    // 抽象方法
    function test() {
        $Factory =  new  FactoryMan;
        $man = $Factory->createOpen();
        $man->jiehun();
        
        $man = $Factory->createIntro();
        $man->jiehun();
        
        
        $Factory =  new  FactoryWomen;
        $man = $Factory->createOpen();
        $man->jiehun();
        
        $man = $Factory->createIntro();
        $man->jiehun();
        
    }
}
 
$f = new Client;
$f->test();
```


### 单例模式
举一个小例子，在我们的windows桌面上，==我们打开了一个回收站，当我们试图再次打开一个新的回收站时，Windows系统并不会为你弹出一个新的回收站窗口==。，也就是说在整个系统运行的过程中，系统只维护一个回收站的实例。这就是一个典型的单例模式运用。

  再举一个例子，网站的计数器，一般也是采用单例模式实现，如果你存在多个计数器，每一个用户的访问都刷新计数器的值，这样的话你的实计数的值是难以同步的。但是如果采用单例模式实现就不会存在这样的问题，而且还可以避免线程安全问题。同样多线程的线程池的设计一般也是采用单例模式，这是由于线程池需要方便对池中的线程进行控制
单例模式，正如其名，允许我们创建一个而且只能创建一个对象的类。

这在整个系统的协同工作中非常有用，特别明确了只需一个类对象的时候。

那么，为什么要实现这么奇怪的类，只实例化一次？

在很多场景下会用到，如：配置类、Session类、Database类、Cache类、File类等等。

这些只需要实例化一次，就可以在应用全局中使用。
单例模式的特点是4私1公：一个私有静态属性，构造方法私有，克隆方法私有，重建方法私有，一个公共静态方法。

其他方法根据需要增加。

最基础的单例模式代码如下：
```
class Singleton
{
    private static $instance = null;

    public static function getInstance()
    {
        if(self::$instance == null) {
            self::$instance = new self();
        }
        return self::$instance;
    }

    private function __construct(){}
    private function __clone(){}
    private function __wakeup(){}
}
```

## 结构型模式
### 适配器模式
适配器模式，即根据客户端需要，将某个类的接口转换成特定样式的接口，以解决类之间的兼容问题。

如果我们的代码依赖一些外部的API，或者依赖一些可能会经常更改的类，那么应该考虑用适配器模式。

下面我们以集成支付宝支付功能为例。

```
/**
 * 支付宝支付类
 */
class Alipay
{
    public function sendPayment()
    {
        echo '使用支付宝支付。';
    }
}

// 客户端代码
$alipay = new Alipay();
$alipay->sendPayment();
```
我们直接实例化Alipay类完成支付功能，这样的客户端代码可能很多。

一段时间后，如果支付宝的Alipay类升级，方法名由sendPayment()变成goPayment()会怎样？

所有用了sendPayment()的客户端代码都要改变。

如果Alipay类频繁升级，或者客户端在很多地方使用，这会是极大的工作量。

2 解决
现在我们用适配器模式来解决。

我们在客户端和Alipay类之间加一个中间类，也就是适配器类，转换原始的Alipay为客户端需要的形式。

为让客户端能调用到统一的类方法，我们先定义一个适配器接口：

```
/**
 * 适配器接口，所有的支付适配器都需实现这个接口。
 * 不管第三方支付实现方式如何，对于客户端来说，都
 * 用pay()方法完成支付
 */
interface PayAdapter
{
    public function pay();
}
因为Alipay类我们无法控制，而且它有可能经常更新，所以我们不对它做任何修改。

我们新建一个AlipayAdapter适配器类，在pay()中转换Alipay的支付功能，如下：

/**
 * 支付宝适配器
 */
class AlipayAdapter implements PayAdapter
{
    public function pay()
    {
        // 实例化Alipay类，并用Alipay的方法实现支付
        $alipay = new Alipay();
        $alipay->sendPayment();
    }
}
```

### 组合模式 
  https://www.awaimai.com/patterns/composite
```
abstract class File
{
    abstract function getSize();
}

class TextFile extends File
{
    public function getSize()
    {
        return 2;
    }
}

class ImageFile extends File
{
    public function getSize()
    {
        return 100;
    }
}
样，在创建文本或图片对象后，就可以通过getSize()方法获取到它们的大小。

然后，我们创建一个目录类，它可以把文件组合起来：
class Dir
{
    private $files = [];

    // 传入参数必须为File文件对象
    public function addFile(File $file)
    {
        $this->files[] = $file;
    }

    public function getSize()
    {
        $size = 0;
        foreach ($this->files as $file) {
            $size += $file->getSize();
        }

        return $size;
    }
}
class NewDir
{
    private $files = [];
    private $dirs = [];

    public function addFile(File $file)
    {
        $this->files[] = $file;
    }

    public function addDir(NewDir $newDir) {
        $this->dirs = $newDir;
    }

    public function getSize()
    {
        $size = 0;
        foreach ($this->files as $file) {
            $size += $file->getSize();
        }

        foreach ($this->dirs as $dir) {
            $size += $dir->getSize();
        }

        return $size;
    }
}


```
是不是比之前又复杂了些？

这还不算，我们还需要修改原来的类，可能无意间又影响原来的功能 。

另外，如果我们现在要计算多级子目录的大小、或者从目录中删除目录，是不是还需要修改原有类？

显然，这个模型无法实现这些复杂的功能，我们需要一个更加灵活的模型。

2 组合模式
组合模式的解决方法是，用抽象类规范统一的对外接口。

然后，让文件类和目录类实现这个接口，并在目录类中递归计算文件的大小。

同时，目录类比文件类两个方法：add()和remove()，用以管理文件对象。

这样，目录类就能用同样的方式获取自身的大小。

并且，还能灵活从目录总增删子目录和文件。
 
组合模式的解决方法是，用抽象类规范统一的对外接口。

然后，让文件类和目录类实现这个接口，并在目录类中递归计算文件的大小。

同时，目录类比文件类两个方法：add()和remove()，用以管理文件对象。

这样，目录类就能用同样的方式获取自身的大小。

并且，还能灵活从目录总增删子目录和文件。

这个接口中规范的方法要根据需求来定义，并且同时要考虑独立对象拥有的功能。

如果独立对象之间有差异的功能，不适合聚合在一起，则不能放在组合类中。

在组合模式中，组合对象和独立对象必须实现一个接口。

其中，组合对象必须包含添加和删除节点对象。

组合模式通过和装饰模式有着类似的结构图，但是组合模式旨在构造类，而装饰模式重在不生成子类即可给对象添加职责。

并且，装饰模式重在修饰，而组合模式重在表示。
 
### 外观模式
实门面模式就是把几个子系统(实例或者类.统一一个统一的接口进行执行,客户端不用关注子系统,只用门面即可 
https://blog.csdn.net/u010861514/article/details/45167219
```
<?php
header("content-type:text/html;charset=utf-8");
// ==================php  门面模式(外观模式)  =============================
 /* 其实门面模式就是把几个子系统(实例或者类.统一一个统一的接口进行执行,客户端不用关注子系统,只用门面即可 )*/
 
// 门面抽象接口 
interface Facade{   
    public function turnOn() ;
    public function turnOff() ;
}
 
// (1) 关闭显示器
class PcLight {
    public function turnOn() {}
    public function turnOff() {
        echo 'turn off PcLight <br>' ;
    }   
}
 
//(2) pc 机器
class Pcmachine {
    public function turnOn() {} 
    public function turnOff() {
        echo 'turn off PcMathion <br>' ;
    }
}
 
// (3) 关闭电源
class Power {
    public function turnOn() {} 
    public function turnOff() {
        echo 'turn off Power <br>' ;
    }
}
 
// 关机的门面角色 
class PcFacade implements Facade{
     
    private  $PcLight ; 
    private  $Pcmachine ; 
    private  $Power ; 
  
    public function __construct(){
     $this->PcLight = new PcLight();
         $this->Pcmachine = new Pcmachine();
         $this->Power = new Power();
    }
     
    // 门面角色的应用
     public function turnOff() { 
          $this->PcLight ->turnOff();
          $this->Pcmachine ->turnOff();
          $this->Power ->turnOff();
     }
      public function turnOn() {}
}
 
// 应用
$button = new PcFacade(); 
$button ->turnOff(); 


```


### 代理模式
   代理模式为其他对象提供一种代理以控制对这个对象的访问。在某些情况下，一个对象不适合或者不能直接引用另一个对象，而代理对象可以在客户端和目标对象之间起到中介的作用。
   
```
<?php
/**
 * Created by PhpStorm.
 * User: LYL
 * Date: 2015/5/16
 * Time: 16:33
 */
 
/**顶层接口
 * Interface IGiveGift
 */
interface IGiveGift
{
    function giveRose();
    function giveChocolate();
}
 
/**追求者
 * Class Follower
 */
class Follower implements IGiveGift
{
    private $girlName;
 
    function __construct($name='Girl')
    {
        $this->girlName=$name;
    }
 
    function giveRose()
    {
        echo "{$this->girlName}:这是我送你的玫瑰，望你能喜欢。<br/>";
    }
 
    function giveChocolate()
    {
        echo "{$this->girlName}:这是我送你的巧克力，望你能收下。<br/>";
    }
}
 
/**代理
 * Class Proxy
 */
class Proxy implements IGiveGift
{
    private $follower;
 
    function __construct($name='Girl')
    {
        $this->follower=new Follower($name);
    }
 
    function giveRose()
    {
        $this->follower->giveRose();
    }
 
    function giveChocolate()
    {
        $this->follower->giveChocolate();
    }
}

header("Content-Type:text/html;charset=utf-8");
//------------------------代理模式测试代码------------------
require_once "./Proxy/Proxy.php";
$proxy=new Proxy('范冰冰');
$proxy->giveRose();
$proxy->giveChocolate();


```





