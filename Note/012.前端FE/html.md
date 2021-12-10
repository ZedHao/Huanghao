# 概述
HTML是超文本标记语言（Hyper Text Markup Language）的缩写。它不是一种编程语言，而是一种标记语言
# 概念
## 标签
## 元素
## 属性

# 用法
## 标题  <h1>至<h6>
## 段落 <p>
## &元素容器  <div> 块级元素
div 元素是块级元素，它是可用于组合其他HTML元素的容器。与CSS一同使用，<div>元素可用于对大的内容块设置样式属性。另一个常见的用途是文档布局。
天生自带了display: block 的属性
## &文本容器 <span> 行内元素
<span>元素是内联元素，可用作文本的容器。与CSS一同使用时，<span>元素可用于为部分文本设置样式属性。
行内元素天生自带了display: inline的属性
## &样式 style
通过使用style属性直接将样式添加到HTML元素，或者间接地在CSS文件中进行定义。样式包括背景颜色、字体、颜色、尺寸、文本对齐等。
## 链接 <a>
## 表格 <table>
## 列表 ul
## &表单 <form>
表单是一个包含表单元素的区域，使用<form>标签定义。表单元素是允许用户在表单中输入信息的元素，如文本框、下拉列表、单选框、复选框等。经常被用到的表单标签是<input>输入标签，输入类型由该标签的type属性定义。

    <form name="input" action="form_action.jsp" method="get">
        Username: 
        <input type="text" name="user" />
        <input type="submit" value="Submit" />
    </form>
## 图片 <img> 
## 头部 <head>
## 主体 <body>
## 脚本 script
脚本使用<script>标签进行定义，可以使用type属性来指定脚本语言。
<script type="text/javascript">

## 事件 
 事件分为窗口事件、表单元素事件、图像事件、键盘事件、鼠标事件
浏览器内置有大量的事件处理器，这些处理器会监视特定的条件或用户行为。将某些特定的事件处理器作为属性添加给特定的标签，并在事件发生时执行对应的JS命令或函数。事件处理器的值是一系列以分号分隔的JS表达式、方法和函数调用，并用引号引起来。事件分为窗口事件、表单元素事件、图像事件、键盘事件、鼠标事件等。
```
<a href="/index.html" onmouseover="alert('Welcome');return false"></a>
```
## URL
