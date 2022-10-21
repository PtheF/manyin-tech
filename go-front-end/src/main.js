import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)

Vue.config.productionTip = false

// store.dispatch("getProductTypes").then(r => {
//   console.log("main.js vuex")
//   console.log(r)
// })

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
