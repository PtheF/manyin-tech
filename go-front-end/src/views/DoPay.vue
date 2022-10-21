<template>
    <div>
        <div class="body" v-show="!fin">
            <h2>蔓茵支付 ￥{{orderInfo.totalPrice / 100}}</h2>

            <div @click="confirmPw = true">
                确认支付
            </div>

            <div>
                <p>1. 付款只在10分钟内有效，超时请退出重新扫码</p>
                <p>2. 密码怎么输都可以，只是走个形式</p>
                <p>3. 安全起见，密码请不要输入自己的常用密码甚至支付密码</p>
            </div>


            <el-drawer
                :visible.sync="confirmPw"
                :direction="'btt'">
                <div style="width: 80%; margin: 0 auto;">
                    <el-input v-model="pw" placeholder="请输入密码" type="password"></el-input>
                </div>
                <div style="
                width: 100px;
                padding: 10px 0 10px 0;
                color: white;
                margin: 20px auto 0;
                text-align: center;
                border-radius: 20px;
                cursor: pointer;
                background-color: #3abe3a" @click="doPay">
                    确定
                </div>
            </el-drawer>
        </div>

        <div style="margin-top: 50px;" v-show="fin">
            <h1 style="text-align: center; color: #3abe3a">付款成功</h1>
        </div>
    </div>

</template>

<script>

import {debounce, isEmpty} from "@/assets/js/utils";
import { ErrorNotify } from "@/assets/js/message";
import {Get, Post} from "@/assets/js/http";

export default {
    name: "DoPay",
    data() {
        return {
            orderId: this.$route.params.orderId,
            confirmPw: false,
            pw: "",
            fin: false,
            orderInfo: {
                totalPrice: 0
            }
        }
    },

    created() {
        Get(`/v1/pay/orderInfo/${this.orderId}`).send(
            ok => {
                if(ok.code === 200){
                    this.orderInfo = ok.data
                    console.log(this.orderInfo)
                }else{
                    ErrorNotify("错误", ok.message)
                }
            },

            err => {
                ErrorNotify("错误", err)
            }
        )
    },

    methods: {
        doPay: debounce(function() {
            if(isEmpty(this.pw)){
                ErrorNotify("错误", "密码为空")
                return
            }

            Post(`/v1/pay/${this.orderId}`).send(
                ok => {
                    if(ok.code === 200) {
                        this.fin = true
                        this.confirmPw = false
                    }else {
                        ErrorNotify("错误", ok.message)
                    }
                },

                err => {
                    ErrorNotify("网络错误", err)
                }
            )
        })
    }
}
</script>

<style scoped>

div.body {
    /*width: 50%;*/
    margin: 0 auto;
    /*height: 100vh;*/
}

div.body > h2 {
    margin: 100px auto 50px;
    text-align: center;
}

div.body > div:nth-child(2) {
    padding: 10px 50px 10px 50px;
    border-radius: 20px;
    margin: 20px auto 0;
    width: 120px;
    text-align: center;
    /*border: 1px solid grey;*/
    cursor: pointer;
    background-color: #3abe3a;
    color: white;
}

div.body > div:nth-child(3) {
    width: 300px;
    margin: 20px auto 0;
    font-size: 14px;
    color: #a8a8a8;
}

</style>