import { Get } from "@/assets/js/http"
import { ErrorNotify } from "@/assets/js/message";

const PROD_TYPES_KEY = "manyin-prod-types"

export default {
    state: {

        productTypes: []

    }, getters: {

        getAllProductTypes: state => state.productTypes,

        getTypeInfo: (state) => {
            // console.log("getters: getTypeInfo")
            // console.log(state.productTypes)
            return (typeId) => {
                // console.log("get type id:" + typeId)
                for(let t of state.productTypes){

                    // console.log("inner method called: ")
                    // console.log(t)

                    if(t.prodTypeId === typeId){
                        // console.log("prod type hit: " + t.prodTypeId)
                        return t
                    }
                }
            }
        }

    }, mutations: {
        setProductTypes(state, value) {
            state.productTypes = null
            state.productTypes = value
            sessionStorage.setItem(PROD_TYPES_KEY, JSON.stringify(value))
        },

        removeTypesCache(state, _) {
            sessionStorage.removeItem(PROD_TYPES_KEY)
        }
    }, actions: {
        getProductTypes({ commit }) {

            // if(store.state.productTypes !== []){
            //     return store.state.productTypes
            // }

            let prodTypes = JSON.parse(sessionStorage.getItem(PROD_TYPES_KEY))
            if(prodTypes === null) {
                Get("/v1/shop/types").send(
                    ok => {
                        if(ok.code === 200){
                            // console.log("get product types")
                            // console.log(ok.data)
                            commit("setProductTypes", ok.data)
                            // console.log("update vuex")
                            // console.log(state.productTypes)
                        }else{
                            console.log("get product types code !== 200")
                            console.log(ok)
                            ErrorNotify("错误", ok.message)
                        }},
                    err => {
                        console.log(err)
                        ErrorNotify("错误", err)
                    })
            }else{
                console.log("get prod types cache")
                commit("setProductTypes", prodTypes)
            }

        }

    }
}