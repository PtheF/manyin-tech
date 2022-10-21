import Vue from 'vue'
import VueRouter from 'vue-router'
import { Get } from '@/assets/js/http'
import { getToken, getUser, isEmpty } from '@/assets/js/utils'
import Home from '../views/Home.vue'
import { ErrorMessage } from "@/assets/js/message"
import loginAuth from "@/assets/js/loginAuth";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: "/response",
    name: "Response",
    component: () => import("../views/Response")
  },
  {
    path: "/pay/:orderId",
    name: "Pay",
    component: () => import("../views/Pay")
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue")
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/Register.vue")
  },
  {
    path: "/prodTypes",
    name: "ProductTypes",
    component: () => import("../views/ProductTypes")
  },
  {
    path: "/space",
    name:"Space",
    component: () => import("../views/Space.vue"),
    meta: {
      login: true,
    }
  },
  {
    path: "/enterprise",
    name: "Enterprise",
    component: () => import("../views/Enterprise.vue"),
    meta: {
      login: true,
    }
  },
  {
    path: "/shop",
    name: "Shop",
    component: () => import("../views/Shop"),
    redirect: "/shop/types",
    children: [
      {
        path: "types",
        name: "ProductTypes",
        component: () => import("../views/ProductTypes.vue")
      },
      {
        path: "/shop/list/:typeId",
        redirect: "/shop/list/:typeId/1"
      },
      {
        path: "/shop/list/:typeId/:page",
        name: "ProductList",
        component: () => import("../views/ProductList.vue")
      }
    ]

  },
  {
    path: "/order",
    name: "Order",
    component: () => import("../views/Order.vue"),
    meta: {
      login: true
    }
  },
  {
    path: "/newOrder",
    name: "NewOrder",
    component: () => import("../views/NewOrder")
  },
  {
    path: "/doPay/:orderId",
    name: "DoPay",
    component: () => import("../views/DoPay"),
    meta: {

    },
  },
  {
    path: "/item/:prodId",
    name: "Product",
    component: () => import("../views/Product")
  },



  {
    path: "/ceoLogin",
    name: "CEOLogin",
    component: () => import("../views/AdminLogin")
  },
  {
    path: "/ceo",
    name: "CEO",
    component: () => import("../views/Admin"),
    redirect: "/ceo/newProduct",
    children: [
      {
        path: "/ceo/newProduct",
        name: "AdminNewProduct",
        component: () => import("../views/AdminNewProduct"),
        meta: {
          admin: true,
        },
      },
      {
        path: "/ceo/productTypes",
        name: "AdminProductTypes",
        component: () => import("../views/AdminProductTypes"),
        meta: {
          admin: true,
        },
      },
      {
        path: "/ceo/products/:typeId/:page",
        name: "CEOProducts",
        component: () => import("../views/AdminProductPage"),
        meta: {
          admin: true,
        },
      }
    ]
  }
]

const router = new VueRouter({
  routes,
  mode: "history"
})

router.beforeEach((to, from, next) => {

  // 如果需要登录
  if(to.meta.login){
    // let user = getUser()
    //
    // // 用户未登录
    // if(!user){
    //   next({
    //     path: "/login?expire=true&uri=" + to.path
    //   })
    // }else{
    //   // 看似用户登录了
    //   // 看看是不是登陆过期了
    //   Get("/v1/auth").send(ok => {
    //     if(ok.code === 200){
    //       next()
    //     }else if(ok.code === 302){
    //       ErrorMessage("登录超时，请重新登录")
    //       next({
    //         path: "/login?expire=true&uri=" + to.path
    //       })
    //     }
    //   }, err => {
    //
    //   })

    loginAuth(
        () => {
          next()
        },
        () => {
          next({
            path: "/login?uri=" + to.path
          })
        },
        () => {
          next({
            path: "/login?expire=true&uri=" + to.path
          })
        }
    )
  } else if(to.meta.admin){
    console.log("enter admin filter")
    let adminToken = sessionStorage.getItem("adminToken")
    if(isEmpty(adminToken)){
      console.log("empty admin token")
      next({
        path: "/ceoLogin"
      })
    }else{
      next()
    }
  }else{
    next()
  }

})


export default router
