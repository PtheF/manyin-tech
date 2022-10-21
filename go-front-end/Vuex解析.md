# Vuex Debug

> create by PtheF on 2022/09/13


## 1. Vuex 复习

vuex 可以理解成整个 Vue 项目的全局空间，只要是没有刷新，那么所有页面都可以共享 Vuex 里面的数据，看看 Vuex 文件的结构：

```js
export default new Vuex.Store({
    state: {
        // state 里面存放全局变量，也就是你希望整个 Vue 项目共享的数据
        num: 10
    }, getters: {
        // 获取 state 的方法，入参是 state，state 就是上面的 state 对象
        getNum(state) {
            return state.num
        }
    }, mutations: {
        // mutations 里面存放更新 state 的方法，入参是 state 和更新的值
        setNum(state, value){
            state.num = value
        }
    }, actions: {
        // actions 进行 ajax 操作，从后端拉取信息，这里的 store 入参就是这里的 store 对象
        // 所以这里传入 store 可行，或者直接在这里结构赋值也可以
        setNumSync(store) {
            ajax(
                //...
            ).then(
                // 这里可以调用 store 的 commit 更新操作，将拉取的数据更新到 state
            )
        },
        
        // 这里就是直接结构赋值，把 传入的 store 对象的 commit 方法结构出来
        setNum2({commit}){
            ajax(
                // ...
            ).then(ok => {
                commit("setNum", ok)
            })
        }
    }, modules: {
        // 其他的模块
        UserStore
    }
})
```

调用的时候不可以直接调用，而是：
```js
// 这里的 this 自然就是 Vue Component，同时不建议
this.$store.getters.getNum // 调用 getter，注意这里没有调用

this.$store.commit("setNum", 20) // 调用 mutations 里面的 setNum 方法，传入方法名和值

this.$store.dispatch("setNumSync") // 调用 actions 里面的 setNumSync
```

## 2. getter 小操作

这里有个小需求，就是比如我的 state 里面有个聊表，我如何通过 getter 把列表的第 i 个拿出来？很显然 getter 没法传入其他参数，那咋办，可以使用高阶函数，让 getter 返回一个函数：

```js
const store = {
    state: {
        users: [
            // ......
        ]
    },
    getters: {
        getUserByIndex: (state) => {
            return (index) => {
                return state.users[index]
            }
        }
    }
}

this.$store.getters.getUserByIndex(index) // 这样就行了，直接就可以传参
```

## 3. 注意点：

Vuex 里面的东西只要不刷新就一直在，但是只要刷新 Vuex 里面的东西就会丢失，所以尽量避免刷新。

这里说一个今天遇见的小 bug，页面一条转，vuex 里面的东西就丢了，我已开始以为是 vuex 更新失败了，结果最后发现不是：是因为我跳转用的都是 a 标签，点击 a 标签就会发生跳转 + 刷新，结果刷新就会导致 vuex 数据丢失，所以咋办呢？把所有的跳转操作全部改成`router-link` 标签即可，vue 的 router 并不会发生刷新，比较的安全。

以及如果 state 里面有数组的话，可以把原数组先刷成 null，然后再更新，这样可以避免某些玄学bug。


