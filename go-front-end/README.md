# 项目小技术点总结

> create by P-F on 2022/09/30

## 1. 页面防抖

啥叫页面防抖？就是有的用户是帕金森，一个按钮点半天，为了防止这个按钮被多次点击重复发送数据破坏接口幂等性，就需要接口防抖，也就是在一段时间内多次点击按钮但是只会执行一次，来防止重复提交数据。

思路就是：设置一个定时器，点击按钮后，需要 200ms 后才执行，如果在这期间多次点击按钮就会重置计时器。

### 实现：

代码如下：

```js

```
