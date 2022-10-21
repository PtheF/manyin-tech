import {Get} from "@/assets/js/http";
import {getUser, isLogin, removeArray} from "@/assets/js/utils";
import {ErrorNotify} from "@/assets/js/message";

const ADDRESS = 'addresses:'
const EMPTY_ADDRESS_FLAG = -1

export default {
    state: {
        addresses: [],
    },

    mutations: {
        setAddresses(state, addresses) {

            let userId = getUser().uuid

            if(addresses === null || addresses === []){
                localStorage.setItem(ADDRESS + userId, EMPTY_ADDRESS_FLAG)
                state.addresses = []
                console.log("set addresses, addresses empty")
                return
            }

            state.addresses = addresses

            localStorage.setItem(ADDRESS + userId, JSON.stringify(addresses))
        },

        setDefaultAddress(state, index) {

            for(let i = 0; i < state.addresses.length; i++) {
                if(state.addresses[i].isDefault) {
                    state.addresses[i].isDefault = false
                }
            }
            state.addresses[index].isDefault = true

            this.commit("setAddresses", state.addresses)
        },

        addAddresses(state, address) {
            state.addresses.push(address)

            this.commit("setAddresses", state.addresses)
        },

        removeAddress(state, index) {
            removeArray(state.addresses, index)
            this.commit("setAddresses", state.addresses)
        }
    },

    getters: {
        getAddresses(state){
            return state.addresses
        },

        getDefaultAddress(state) {
            for (let address of state.addresses) {
                if (address.isDefault) {
                    return address
                }
            }

            return null
        }
    },

    actions: {
        fetchAddresses({commit}) {

            if(!isLogin()) {
                console.log("user not login")
                return
            }

            let userId = getUser().uuid

            let addresses = JSON.parse(localStorage.getItem(ADDRESS + userId))

            if(addresses !== null && addresses !== EMPTY_ADDRESS_FLAG){
                console.log("address cache hit: ")
                console.log(addresses)
                commit("setAddresses", addresses)
                return
            }

            if(addresses === EMPTY_ADDRESS_FLAG) {
                console.log("address cache hit but empty")
                commit("setAddresses", [])
                return
            }

            console.log("fetch addresses")

            Get("/v1/space").send(
                ok => {
                    // if(ok.data.addresses !== null && ok.data.addresses !== []){
                    //     console.log("fetch addresses not empty, commit")
                    //     commit("setAddresses", ok.data.addresses)
                    // }else{
                    //     console.log("fetch addresses empty")
                    //     localStorage.setItem(ADDRESS + userId, EMPTY_ADDRESS_FLAG)
                    // }

                    commit("setAddresses", ok.data.addresses)
                },

                err => {
                    ErrorNotify("错误", "拉取地址信息错误:" + err)
                }
            )
        }
    }
}