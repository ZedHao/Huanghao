# Html、Css和Javascript
这三个要素共同构成了Web开发的基础。

HTML：页面的结构-标题，正文，要包含的任何图像
CSS：控制该页面的外观（这将用于自定义字体，背景颜色等）
JavaScript：不可思议的第三个元素。创建结构（HTML）和美学氛围（CSS）后，JavaScript使您的网站或项目充满活力
# Javascript作用
1. 表单数据验证：表单数据验证是JavaScript最基本也是最能体现效率的功能。
2. 动态HTML（即DHTML）：动态HTML指不需要服务器介入而动态变化的网页效果，包括动态内容、动态样式、动态布局等。 比如改变盒子的尺寸，背景颜色，图片等。
3. 用户交互：用户交互指根据用户的不同操作进行的响应处理。例如：联动菜单等。
4. 数据绑定：HTML中表单和表格能够以.txt文件定义的数据源，通过对位于服务器端的数据源文件的访问，便可以将数据源中的数据传送到客户端，并将这些数据保存在客户端。
5. 少量数据查找：能够实现在当前网页中进行字符串的查找和替换。
6. AJAX核心技术：AJAX即异步JavaScript+XML。该对象提供一种支持异步请求的技术，使客户端可以使用JavaScript向服务器提出请求并处理响应，但并不影响用户在客户端的浏览。
7. Nodejs就是使用的javascript做后端，是目前为止唯一的一个既能做前端、又能做后端的语言。

# js中输入输出语句
    console.log 控制台输出日志
    console.dir 对象的形式打印一个对象
    document.write 往页面中写入内容
    alert 弹框警告
    confirm 确认框
    prompt 输入框
# JS的三个组成部分
javascript的三个组成部分：ECMAScript、BOM、DOM
## 1. ECMAScript： 定义了javascript的语法规范。
## 2. BOM： 一套操作浏览器功能的API。Browser Object Model

## 3. DOM： 一套操作页面元素的API。Document Object Model    
### 注册事件
    事件源： 触发事件的元素
    事件名称： 触发事件的名称
    事件处理函数： 触发事件时调用的函数
    var img = document.getElementById("img");
    img.onclick = function() {
      img.src = "images/2.gif";
    }