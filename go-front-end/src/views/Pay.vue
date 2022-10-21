<template>
    <div>
<!--        <header-nav/>-->
        <div class="body" v-show="!payFinish">
            <h1>蔓茵支付</h1>
            <h2>
                <i class="el-icon-caret-right"/>声明
            </h2>
            <p>
                本网站不包含任何支付业务，没有调用支付宝、微信、银联或者任何其他第三方支付接口，
                下面的付款二维码只会跳转到另外一个页面模拟支付，<strong>不会</strong>对您的计算机产生任何不良影响，
                更<strong>不会</strong>出现诸如 植入木马、盗取信息 等违法犯罪行为。
            </p>
            <br><br>
            <hr>
            <br>
            <div class="pay">
                <div style="text-align: center;" v-show="typ !== 'prod'">
                    dev: <a :href="payUrl">点击支付(无法扫码情况下)</a>
                </div>
                <div style="text-align: center; font-size: 14px; color: #606266;">
                    请在订单创建后<strong style="color: red">30分钟</strong>内完成付款，否则将被退单
                </div>
                <div id="qr"></div>

                <div class="pay-func">
                    <el-button @click="back">
                        <el-icon class="el-icon-error"/>
                        取消支付
                    </el-button>

                    <el-button type="success" @click="done">
                        <el-icon class="el-icon-check"></el-icon>
                        完成支付
                    </el-button>
                </div>
                <div style="text-align: center; margin-top: 20px; font-size: 13px; color: #606266">
                    <i class="el-icon-info"></i>
                    我也不知道我为啥要设置一个取消支付的按钮，谁会点？都是直接退出。但是弄一个显得对称:)
                </div>
            </div>

        </div>

        <div v-show="payFinish">
            <h1 style="text-align: center; color: #3abe3a; margin-top: 50px;">付款成功</h1>
        </div>
<!--        <Footer style="position: relative;"/>-->
    </div>
</template>

<script>
import HeaderNav from "@/components/HeaderNav";
import Footer from "@/components/Footer";
import env from "@/../public/env.json"
import {Get} from "@/assets/js/http";
import {ErrorNotify} from "@/assets/js/message";

let qrcode = require("arale-qrcode")

export default {
    name: "Pay",
    components: {Footer, HeaderNav},

    data() {
        return {
            interv: null,
            payUrl: env.api[env.model].host + "/doPay/" + this.$route.params.orderId,
            typ: env.typ,
            orderId: this.$route.params.orderId,
            payFinish: false,
            timer: null,
        }
    },

    created(){
        this.interv = setInterval(() => {
            this.checkOrderStatus()
        }, 1000);
    },

    mounted() {

        let testQr = new qrcode({
            render: "svg",
            // text: "https://www.baidu.com",
            text: this.payUrl,
            size: 120
        })

        document.querySelector("#qr").appendChild(testQr)
    },

    beforeDestroy() {
        clearInterval(this.interv)
        clearTimeout(this.timer)
    },

    methods: {

        back() {
            clearInterval(this.interv)
            this.$router.back()
        },

        done() {
            this.checkOrderStatus()
        },

        checkOrderStatus() {
            Get(`/v1/order/status/${this.orderId}`).send(
                ok => {
                    if(ok.code === 200){
                        if(ok.data === 1) {
                            clearInterval(this.interv)
                            this.payFinish = true

                            this.timer = setTimeout(() => {
                                this.$router.back()
                            }, 2000)

                        } else {
                            console.log(`order: ${this.orderId} pay status: ${ok.data}`)
                        }
                    }else{
                        clearInterval(this.interv)
                        this.$alert(ok.message, '错误', {
                            confirmButtonText: '确定',
                            type: "error",
                            callback: action => {
                                this.$router.back()
                            }
                        });
                    }
                },
                err => {
                    ErrorNotify("网络错误", err)
                }
            )
        }
    }
}
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 60% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 80% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }
}
    div.body {
        width: 50%;
        margin: 30px auto 50px;
        /*min-height: 90vh;*/
    }

    div.body > h1 {
        text-align: center;
        border-bottom: 1px solid grey;
        padding-bottom: 20px;
    }

    div.body > h2 {
        margin-top: 15px;
    }

    div.body > p {
        text-indent: 30px;
        line-height: 30px;
        color: #606266;
        margin-left: 20px;
        font-size: 14px;
    }

    div#qr {
        width: 120px;
        margin: 20px auto 0;
    }

    div.pay-func {
        display: flex;
        width: 200px;
        margin: 20px auto 0;
        justify-content: space-around;
    }
</style>