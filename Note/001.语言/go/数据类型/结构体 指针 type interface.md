# 指针  
## C++中指针(*)、取地址(&)、解引用(*)与引用(&)
### 如何使用指针？——解引用与指针赋值

````````
#include<cstdio>

int main(void)
{
	int num = 7;//声明num地址 地址值为7
	int *p = &num;//取num地址的值 并让 p指针指向num的地址
	printf("数值%d所在的地址是 %p\n", num, p);
	printf("指针p所指向的地址为 %p , 该地址上所保存的值为%d\n", p, *p);
	*p = 100;  //把指针指向的地址指 由 7改为100
	printf("指针p所指向的地址为 %p , 该地址上所保存的值为%d\n", p, num);
	return 0;
}
``````````

注意这里*操作符为解引用操作符，它返回指针p所指的地址所保存的值。

我们可以对*p赋值，从而改变p所指的地址上说保存的值，从而改变指向此地址的变量num的值。（num的值变为100）

当然，我们也可以给指针p赋值，使其指向另外一个地址：
```````
#include<cstdio>

int main(void)
{
	int num = 7, another = -5;
	int *p = &num;
	p = &another;
	printf("%d\n", *p);//-5
	return 0;
}
```````````````
b)指针与数组：

指针和迭代器非常像，我们可以说指针就是数组的迭代器：

 

#include<cstdio>

int val[100];

int main(void)
{
	for (int i = 0; i < 100; ++i)
		val[i] = i;
	int *p = val;///等价于int *p = &val[0];
	///数组的变量名就是一个指针
	printf("%d\n", *p); ///指针p指向val的第一个元素，即val[0]
	
	int t = 100;
	while (t--)
		///可以直接对指针进行加减运算，就和迭代器一样
		printf("%d\n", *(p++));///输出0~99
		
	///指针可以做差：
	int *p2 = &val[10], *p3 = &val[20];
	printf("%d\n", p3 - p2); //10
	printf("%d\n", p2 - p3); //-10
	
	///还可以比比较大小：
	printf("%d\n", p2 < p3 ? p3 - p2 : p2 - p3); //10
	return 0;
}

对于多维数组，只要把指向指针的指针理解好了就行：

 

#include<cstdio>

int val[100][100];

int main(void)
{
	val[2][1] = 666;
	///如何用指针取出val[2][1]？
	printf("%d", *(*(val + 2) + 1));
	return 0;
}

##  c++指针交换变量  

https://blog.csdn.net/weixin_38629529/article/details/82222878
###  函数的传参：三种情况 传值，传引用（C++中使用的，不支持C）,传地址

# type interface 
