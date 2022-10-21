import Vue from 'vue'
import Vuex from 'vuex'
import ProductTypes from "@/store/productTypes";
import Address from "@/store/address";
import Product from "@/store/product"

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    ProductTypes, Address, Product
  }
})

export default store
