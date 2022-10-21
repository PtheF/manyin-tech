# Axios 请求设置请求头 BUG

> 2022/08/12 10:25

## 1. 症状

什么个情况呢？我们 axios 设置了个请求拦截器，然后在拦截器里面设置请求头，把 token 放到 Authorization 请求头上：

```js
axios.interceptors.request.use(config => {
    config.headers["Authorization"] = getToken()
})
```

结果，这个请求就死活发不出去了，依次会发两个请求，第一次是 OPTIONS 第二次是 POST，但是无一例外报错，而且 chrome 浏览器贼傻逼，不告诉原因，就告诉我 NET:ERR。

一开始我以为是我拦截器的问题，然后今天没有使用拦截器，使用原始axios发了个请求，嘿，您猜怎么着，照样发不出去，然后定位错误，发现是 headers 的问题，只要我的 headers 带上 Authorization，那这请求铁定发布出去。都给我整麻了。

这次事故持续了近3个小时，期间崩溃了10次嚎啕大哭了9次晕厥了8次。

## 2. 定位

发了两个请求，第一次是 OPTIONS，第二次才是 POST，所以初步认为这个 BUG 应该和 OPTIONS 这个请求有关系，所以我在服务端自己写了一个 OPTIONS 请求的处理器，直接给他返回字符串，嘿，您猜怎么着，即便如此，这个 OPTIONS 还是报错。

然后我打开了 firefox 浏览器，嘿，firefox 还是厚道，告诉我这么一串：

```text
CORS 预检响应的 'Access-Control-Allow-Headers'，不允许使用头 'contenttype'
```

WTF? 火狐告诉我我这个请求头不允许？虽然这个请求头不是 Authorization 而是我写错的 contentType，但是还是给我提供了思路，很可能这个 Authorization 请求头是因为某种原因被 Ban 了。那么是出于什么原因？很可能就和这个 OPTIONS 有关。

所以我再次推测，可能是第一次 OPTIONS 请求询问了服务器，我下次的 POST 可以带哪些请求头啊？结果服务器告诉我：反正不能是 Authorization，结果下次 POST 请求，我带上 Authorization 就挂了。

## 3. 解决

既然推测是这么回事，就得解决一下，因为我的服务器是 Beego 写的，所以我就查了一下 beego 解决 OPTIONS 请求的办法，真就查到了：

[beego 解决 options](https://cloud.tencent.com/developer/article/1719610)

```go
var success = []byte("SUPPORT OPTIONS")

var corsFunc = func(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Authorization")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)
}
```

然后我把这一段扔到了项目里，再试，嘿，真就行了。然后我再做了个测试，我把里面 Access-Control-Allow-Headers 里面的 Authorization 去掉了，然后再请求，果然，出现了以前的问题。

## 4. 小结

很好，这么个恶心我好几个小时的 BUG 就这么被解决了，但是有这么个问题，为什么 axios 会发送一个 options 请求？我又去网上查了一下，查到个博客，这里直接复制：

[axios 的 options请求](https://www.jianshu.com/p/9e52ca6b8818)

**简单请求**

满足下面两个条件的请求是简单请求：

**请求方式是以下三种之一：** 
- HEAD 
- GET 
- POST

**HTTP的头信息不超出以下几种字段：**
- Accept 
- Accept-Language 
- Content-Language 
- Last-Event-ID 
- Content-Type

**但是Content-Type的值，只限于三个值：**
- application/x-www-form-urlencoded
- multipart/form-data
- text/plain


**复杂请求**

非简单请求就是复杂请求。

复杂请求的CORS请求，会在正式通信之前，增加一次HTTP查询请求，称为“预检”请求（preflight）。预检请求为OPTIONS请求，用于向服务器请求权限信息。预检请求被成功响应后，才会发出真实请求，携带真实数据。

axios默认请求就是application/json,所以不需要自己加上头部（不需要在config中加headers），所以总是会发出options请求的，看看是不是配置的时候加了不必要的headers配置项。
另外，如果真的需要预检，后台也需要进行设置，允许options请求。

作者：LinkLiKang
链接：https://www.jianshu.com/p/9e52ca6b8818
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

嗯哼，就这么回事，这个 BUG 就算过去了。