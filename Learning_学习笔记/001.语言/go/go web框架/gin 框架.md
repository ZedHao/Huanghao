[toc]
# 前提
## web 服务基本概念
1. Request 
1. Response
1. Connnetion [用户的请求连接]
1. Hanlder 【处理请求和生成返回信息】
1. Router【url转handler】

## web 框架的通用功能
Clinet -> Response -> [Multiplexer(router)]->hanlder->response->clinet 
url转handler 注册 监听 解析

# net/http

![image](96A5B9D76B93479B935A6BF7428B6355)

![image](C03C20CA316542BAB443E0BCD8BC069C)

![image](FA8D08C911D244BFB2C510076E50919F)

## 注册路由函数
  所有的 path 和hander 都存在ServeMux里
  hander 是一个 interface 定制框架都是实现了该interFace
  ![image](915CC7AB467F41698FF0C946173C2D9A)
## 开始监听 处理请求 返回响应

# gin 源码解读

```
import (
	"github.com/gin-gonic/gin"

)
func WebRoot(context *gin.Context){
	context.String(200,"hello world")
}

func main() {
    // 初始化引擎
	engine := gin.Default()
	// 注册路由和处理函数
	engine.Any("/", WebRoot)
	// 绑定端口并启动应用
	engine.Run(":8999")
}
```

## 背景
net/http 暴露的函数签名不友好 w https.Resopnse  http。Request 解析请求和回写结果不方便
gin 做了什么 addRoute  生成了 engine  engine  的halder 联系了 gin 的hander serverHttps
gin 的 engine 的trees 和 RouteGroup 路由管理  中间件 Next（）

## 优点 
![image](55418F13CA304B9FA9F3CC2A78DEEED2)




 