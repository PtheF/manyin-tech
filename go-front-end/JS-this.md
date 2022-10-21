# 技术点

> create by P-F on 2022/10/16


## 1. 函数 this 指向问题

这可以说是一个基础问题，可惜我的 JS 基础并不好，今天才来补课。

### 1.1 function 作为对象

JS 里面的 function 与其说是函数，不如说是我们理解的方法，Java 里面的方法中的 this 就是指向的当前类的实例化对象，JS 的 function 也一样，在 function 作为普通方法调用的时候，function 内的 this 指向的就是当前 function 所在的对象。

```js
const user = {
    name: "jack",
    age: 20,
    hello: function () {
        console.log(`hello, i'm ${this.name}, i'm ${this.age} years old`)
    }
}
```

上面这个例子中，this 就指向的是当前 hello 方法的所属对象 user，所以就可以使用 `this.name` 这种的来访问对象的成员。

### 1.2 function 作为函数

如果说 function 没有在对象内，而是直接放在了上下文中，比如直接放在了 script 标签内，那么 function 的 this 其实也同理指向了 function 的所属对象，但是这个时候 function 的所属对象就变成了 function 的上下文。

举个例子：如果 function 直接放在了 script 标签内，那么就可以理解成这个对象是当前页面 window 对象的方法，所以 function 的 this 就会指向 window，那如果在 nodejs 环境中，function 直接写在文件内，那么 this 就指向了 nodejs 中的上下文对象 `global`。

```js
<script>
    function a() {console.log(this)}
    a()
    
    <!-- window{...} -->
</script>

// nodejs:

function a() {
    console.log(this)
}

a()

// Object [global]...
```

### 1.3 this 不可继承

如果出现了嵌套函数，那内层函数的 this 又变成谁了？

```js
const obj = {
    a: 10,
    b: 20,
    
    foo: function() {
        console.log(this)
        
        let c = function() {
            console.log(this)
        }
        
        c()
    }
}

obj.foo()
```

这里分析一下：外层的 foo 方法是 obj 对象的方法，所以 foo 方法内的 this 肯定是指向了 obj，这个没毛病，但是里面的 c 函数也是指向了 obj 么？结果显示并非，内层函数的 this 仍然是 window，似乎也好理解，因为 c 函数虽然在obj 对象的方法内部，但是本身并不属于某个对象，所以 this 还是 window。

那怎么改才能让内层的函数可以访问 obj？

```js
const obj = {
    self: this,
    
    foo: function() {
        (function(){
            console.log(self)
        })()
    }
}

obj.foo()
```

这样就可以了。

### 1.4 function 作为构造函数

这里就涉及到了 JS 面向对象的内容了，面向对象还涉及到 JS 的原型和原型链，以后会详细看的，目前还不会。只是大致说一下：function 可以作为构造函数，使用 new 关键字可以创建对象，看下面的例子：

```js
function Person(name, age) {
    return {name: name, age: age}
}

let p = Person('jack', 20)

function User(name, age) {
    
    // 这里的 this 不代表 window 了，而是代表当前要构造的对象 {  }
    this.name = name
    this.age = age
    this.hello = function() {
        console.log(`hello, i'm ${this.name}, i'm ${this.age} years old`)
    }
}

let u = new User('jack', 20)
u.hello()
```

上面的例子里面，第一个 Person 方法，也可以创建一个对象，但是略有点麻烦，那么就可以使用下面这种办法：使用 new 关键字来创建对象。

new 后面跟一个函数，然后函数会被当做构造函数执行，作为构造函数时，函数内的 this 就代表我要构造出来的这个对象。

### 1.5 箭头函数的this

而箭头函数中，this 又是另一番镜像。箭头函数中的 this 只是一个普通的变量，没错，就是一个普通的变量。根据函数作用域，当函数内访问变量时，会在函数声明的地方找这个变量，如果找不到，就会从声明位置往上找，比如：

```js
let a = 10

function a(){
    console.log(a)
}

function b() {
    let a = 20
    a()
}

b()
```

这里会打印 10，为啥？因为 b 函数内调用 a 函数，a 函数要找 a 变量，从 a 的声明位置开始找，a 函数内没有，则往外找，也就找到了全局里面的 `let a = 10`。

同理，调用一个箭头函数，对于箭头函数来说 this 就是一个普通变量，他会根据上面的原则去找这个 this 到底是个啥，比如：

```html
<script>
    console.log(this)
    
    let a = function() { console.log(this) }
    let b = () => { console.log(this) }
    
    a()
    b()
</script>
```

这俩输出的 this 全是 window，为啥？因为在 script 标签内，this 就是 window，a 函数看作是 window 的方法，所以 a 的 this 是 window，但是 b 不一样，b 函数里面找 this，没找到，就去 b 的上层作用域里面找，也就是 script 标签内，结果发现了 script 标签内有 this，还是 window，所以 b 打印 this 也是 window。

这里虽然俩 this 都是 window，但意义是不一样的，在 nodejs 环境：

```js
let a = function() { console.log(this) }
let b = () => { console.log(this) }

a()
b()
```

这回就不一样了，a 出来的 this 是 global，而 b 的 this 就是 { }，因为 nodejs 中，函数的看做是 global 的方法，但是这个 global 可不是这个 js 文件的作用域，一个 js 文件里面默认就有一个对象 this = {}，所以调用 b 函数找 this 就找到了 { }。

一句这个特性，我们就可以优化上面提到的 this 继承问题：

```js

const obj = {
    a: 10,
    foo() {
        console.log(this)
        
        let c = () => {
            console.log(this.a)
        }
        
        c()
    }
}

obj.foo()
```

这次就可以正常让内层函数正确的访问到外层方法的 this 对象，原因也很简单，c 函数调用时，碰到了 this，但是不知道这个 this 是啥声明里面也没有，就跑到 c 的上层作用域也就是 foo 方法内去找 this，结果在 foo 里面找到了 this 也就是 obj 对象，所以就可以正确访问 obj 了。


