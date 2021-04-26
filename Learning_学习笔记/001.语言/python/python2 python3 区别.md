 1.print不再是语句，而是函数，比如原来是 print 'abc' 现在是 print('abc')但是 python2.6+ 可以使用 from __future__ import print_function 来实现相同功能>  
 
2. 在Python 3中，没有旧式类，只有新式类，也就是说不用再像这样 class Foobar(object): pass 显式地子类化object但是最好还是加上. 主要区别在于 old-style 是 classtype 类型而 new-style 是 type类型>

3. 原来1/2（两个整数相除）结果是0，现在是0.5了python 2+ 以上都可以使用 from __future__ import division 实现改特性, 同时注意 // 取代了之前的 / 运算>  

4. 新的字符串格式化方法format取代%错误, 从 python2.6+ 开始已经在str和unicode中有该方法, 同时 python3依然支持 % 算符> 
5. xrange重命名为range同时更改的还有一系列内置函数及方法, 都返回迭代器对象, 而不是列表或者 元组, 比如 filter, map, dict.items 等>  

7. !=取代  <   >  python2 也很少有人用  <  >  所以不算什么修改>  

8. long重命名为int不完全对, python3 彻底废弃了 long+int 双整数实现的方法, 统一为 int , 支持高精度整数运算.> 

9. except Exception, e变成except (Exception) as e只有 python2.5 及以下版本不支持该语法. python2.6 是支持的. 不算新东西>  10. exec变成函数类似 print() 的变化, 之前是语句
# 