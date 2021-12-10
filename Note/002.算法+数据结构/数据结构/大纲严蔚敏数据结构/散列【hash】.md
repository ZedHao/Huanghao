![image](https://img-blog.csdn.net/20140517233813890?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQveWFuZ195dWxlaQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
# 场景
给定a、b两个文件，各存放50亿个url，每个url各占64字节，内存限制是4G，让你找出a、b文件共同的url？
可以估计每个文件安的大小为5G×64=320G，远远大于内存限制的4G。所以不可能将其完全加载到内存中处理。考虑采取分而治之的方法。

 分而治之/hash映射：

遍历文件a，对每个url求取，然后根据所取得的值将url分别存储到1000个小文件（记为，这里漏写个了a1）中。这样每个小文件的大约为300M。遍历文件b，采取和a相同的方式将url分别存储到1000小文件中（记为）。这样处理后，所有可能相同的url都在对应的小文件（）中，不对应的小文件不可能有相同的url。然后我们只要求出1000对小文件中相同的url即可。
