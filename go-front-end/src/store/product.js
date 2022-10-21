import {isEmpty} from "@/assets/js/utils";

const ORDER_PREFIX = "order:token:"

export default {
    state: {
    },

    getters: {
        getOrder(_){
            return (token) => {
                let orderJson = localStorage.getItem(ORDER_PREFIX + token);
                if(isEmpty(orderJson)){
                    return null
                }

                return JSON.parse(orderJson)
            }
        }
    },

    mutations: {
        setOrder(state, prod) {
            localStorage.setItem(ORDER_PREFIX + prod.orderToken, JSON.stringify(prod))
        }
    }
}