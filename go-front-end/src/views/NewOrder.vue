<template>
    <div>
        <header-nav></header-nav>

        <div class="body">
            <div class="order-logo">
                <div>
                    <!--                    <img style="height: 60px; width: 60px; line-height: 90px;" src="/logo.png" alt="">-->
                    <h1>蔓茵商城</h1>
                    <div>创建订单</div>
                </div>
            </div>
            <div class="order-body">
                <div>
<!--                    <h2>创建订单</h2>-->
                </div>

                <!-- 商品描述 -->
                <div>
                    <img :src="prodDetail.prodAvatar" alt="">
                    <div class="prod-info">
                        <div>{{ prodDetail.prodName}}</div>
                        <div>{{ prodDetail.prodDesc }}</div>
                        <div>
                            <strong>￥{{ prodDetail.prodPrice / 100 }} / {{prodDetail.prodUnit}}</strong>
                        </div>
                    </div>
                </div>

                <!-- 地址 -->
                <div>
                    <h3>收货地址</h3>

                    <el-empty v-show="receivingAddress.id === -1">
                        <el-button type="primary" @click="addAddress = true">新增地址</el-button>
                    </el-empty>

                    <div v-show="receivingAddress.id !== -1" class="rec-address">
                        <div class="address-info-row">
                            <el-row>
                                <el-col :span="6">
                                    <i class="el-icon-tickets"></i>
                                    名称
                                </el-col>
                                <el-col :span="18">
                                    {{ receivingAddress.name }}
                                </el-col>
                            </el-row>
                        </div>

                        <div class="address-info-row">
                            <el-row>
                                <el-col :span="6">
                                    <i class="el-icon-user-solid"></i>
                                    收件人
                                </el-col>
                                <el-col :span="18">
                                    {{ receivingAddress.addressee }}
                                </el-col>
                            </el-row>
                        </div>

                        <div class="address-info-row">
                            <el-row>
                                <el-col :span="6">
                                    <i class="el-icon-phone"></i>
                                    联系电话
                                </el-col>
                                <el-col :span="18">
                                    {{ receivingAddress.phone }}
                                </el-col>
                            </el-row>
                        </div>

                        <div class="address-info-row">
                            <el-row>
                                <el-col :span="6">
                                    <i class="el-icon-location"></i>
                                    收货地址
                                </el-col>
                                <el-col :span="18">
                                    {{ receivingAddress.address }}
                                </el-col>
                            </el-row>
                        </div>

                        <div class="address-func">
                            <el-button type="primary" size="mini" @click="addAddress = true">添加地址</el-button>
                            <el-button type="primary" size="mini" @click="selectAddress = true">切换地址</el-button>
                        </div>
                    </div>

                </div>

                <!-- 价格 -->
                <div>
                    <h3>价格</h3>
                    <div>
                        <span>数量：</span>
                        <span>
                            <strong>
                                {{ buyNumber }} {{ prodDetail.prodUnit }}
                            </strong>
                        </span>
                    </div>

                    <div>
                        <span>总额：</span>
                        <span>
                            <strong>
                                {{ buyNumber }} x {{ prodDetail.prodPrice / 100 }} = {{ (buyNumber * prodDetail.prodPrice) / 100 }} 元
                            </strong>
                        </span>
                    </div>


                </div>

                <!-- 条款 -->
                <div>
                    <h3>订单条款</h3>
                    <ol>
                        <li>该页面只是拟真，后台也会记录订单信息，但是并不会发货</li>
                        <li>创建订单后要求付款，付款也只是模拟，并非真实的付款</li>
                        <li>该项目没有调用任何付款接口，可以放心使用</li>
                    </ol>
                </div>

                <!-- 下单按钮 -->
                <div>
                    <span @click="createOrder">创建订单</span>
                </div>
            </div>
        </div>

        <v-mask :show="selectAddress" @close="selectAddress = false">
            <div class="addresses">
                <div v-for="(address, index) in addresses" :key="index"
                     class="address-info"
                    :style="{border: address.id === receivingAddress.id ? `1px solid #3ec466` : `1px solid rgba(96, 98, 102, 0.4)`}">
                    <div class="address-info-row">
                        <el-row>
                            <el-col :span="10">
                                <i class="el-icon-user-solid"></i>收货人
                            </el-col>
                            <el-col :span="14">{{ address.addressee }}</el-col>
                        </el-row>

                    </div>

                    <div class="address-info-row">
                        <el-row>
                            <el-col :span="10">
                                <i class="el-icon-phone"></i>联系电话
                            </el-col>
                            <el-col :span="14">{{ address.phone }}</el-col>
                        </el-row>
                    </div>

                    <div class="address-info-row">
                        <el-row>
                            <el-col :span="10">
                                <i class="el-icon-location"></i>地址
                            </el-col>
                            <el-col :span="14">{{ address.address }}</el-col>
                        </el-row>
                    </div>

                    <div class="address-info-row" style="margin-top: 20px;">
                        <el-row>
                            <el-col :offset="16">
                                <span class="select-address-but" @click="changeReceivingAddress(index)">选择地址</span>
                            </el-col>
                        </el-row>
                    </div>

                </div>
            </div>
        </v-mask>

        <address-editor :show="addAddress" @cancel="addAddress = false" @add="doAddAddress"></address-editor>

        <Footer style="position:relative"></Footer>
    </div>
</template>

<script>

import {debounce, isEmpty} from "@/assets/js/utils";
import {ErrorNotify, OkNotify} from "@/assets/js/message";
import {queryString} from "@/assets/js/uri";
import Footer from "@/components/Footer";
import HeaderNav from "@/components/HeaderNav";
import VMask from "@/components/VMask";
import AddressEditor from "@/components/AddressEditor";
import { Put } from "@/assets/js/http";

export default {
    name: "newOrder",
    components: {
        AddressEditor,
        VMask,
        "Footer": Footer,
        "header-nav": HeaderNav
    },


    data() {
        return {
            // orderToken: this.$route.params.token,
            selectAddress: false,
            addAddress: false,
            orderToken: queryString("token"),
            // order: null,
            buyNumber: 0,
            prodDetail: null,
            addresses: [],

            receivingAddress: {
                id: -1,
                name: "",
                addressee: "",
                address: "",
                phone: ""
            }
        }
    },

    created() {
        console.log(`orderToken=${this.orderToken}`)

        // this.test(20)
        //
        // console.log("log this.test()")
        // console.log(this.test)

        if(isEmpty(this.orderToken)){
            this.$router.back()
            return
        }

        this.order = this.$store.getters.getOrder(this.orderToken)

        this.buyNumber = this.order.buyNumber
        this.prodDetail = this.order.prodDetail

        console.log(this.order)

        this.$store.dispatch("fetchAddresses")
        this.addresses = this.$store.getters.getAddresses
        let defaultAddress = this.$store.getters.getDefaultAddress

        if (defaultAddress !== null) {
            this.receivingAddress = defaultAddress
        }
    },

    methods: {

        // test: debounce(function(value) {
        //     console.log("debounce callback this:")
        //     console.log(this)
        //     console.log(value)
        // }),

        // test: debounce((value) => {
        //
        //     console.log("debounce callback this:")
        //     console.log(this)
        //     console.log(value)
        // }),

        changeReceivingAddress(index) {
            this.receivingAddress = this.addresses[index]
            // OkNotify("成功", "修改地址成功")
            this.selectAddress = false
        },

        createOrder: debounce(function() {
            const orderInfo = {
                prodId: this.prodDetail.prodId,
                prodName: this.prodDetail.prodName,
                prodAvatar: this.prodDetail.prodAvatar,
                buyNum: this.buyNumber,
                totalPrice: this.buyNumber * this.prodDetail.prodPrice,
                addressee: this.receivingAddress.addressee,
                address: this.receivingAddress.address,
                addresseePhone: this.receivingAddress.phone
            }

            if(this.receivingAddress.id === -1) {
                ErrorNotify("错误", "未添加收货地址")
                return
            }


            Put("/v1/order/create/" + this.orderToken).setJson(orderInfo).send(
                ok => {
                    if(ok.code === 200) {

                        // alert(`order id = ${ok.data}`)

                        this.$confirm("订单创建成功，是否现在付款？", "成功", {
                            confirmButtonText: '支付',
                            cancelButtonText: '取消',
                            type: "success"
                        }).then(() => {
                            this.$router.replace(
                                {
                                    path: "/pay/" +ok.data
                                }
                            )
                        }).catch(() => {
                            this.$router.replace(`/item/${this.prodDetail.prodId}`)
                        })
                    }else{
                        this.$alert(ok.message, "错误", {
                            confirmButtonText: '确定',
                            type: "error",
                            callback: action => {
                                this.$router.replace(`/item/${this.prodDetail.prodId}`)
                            }
                        })
                    }
                },

                err => {
                    ErrorNotify("网络错误", err)
                }
            )
        }),

        doAddAddress(address) {
            Put("/v1/address").setJson(address).send(
                ok => {
                    if(ok.code === 200){
                        OkNotify("成功", "地址添加成功")


                        if(this.receivingAddress.id === -1){
                            address.isDefault = true
                        }

                        address.id = ok.data
                        this.$store.commit("addAddresses", address)
                        this.receivingAddress = address

                        this.addAddress = false
                    }else{
                        ErrorNotify("错误", ok.message)
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
    div.order-body {
        width: 65% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.order-body {
        width: 75% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.order-body {
        width: 90% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.order-body {
        width: 100% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.order-body {
        width: 100% !important;
    }
}

div.body {
    min-height: 90vh;
}

div.order-logo {
    background: linear-gradient(-150deg, #496f73 15%, #2c4d4b 70%, #3c4859 94%);
    overflow: hidden;
}

div.order-logo > div {
    margin: 20px auto;
}

div.order-logo > div > h1 {
    text-align: center;
    color: white;
}

div.order-logo > div > div{
    margin-top: 10px;
    text-align: center;
    color: #dadada;
}

div.order-body {
    width: 50%;
    margin: 0 auto;
}

div.order-body > div:nth-child(2) {
    display: flex;
    padding: 10px 10px 20px 10px;
    margin-top: 20px;
    border-bottom: 1px solid rgba(128, 128, 128, 0.5);
}

div.order-body > div:nth-child(2) > img {
    width: 100px;
    height: 100px;
}

div.prod-info {
    margin-left: 20px;
}

div.prod-info > div:nth-child(1) {

}

div.prod-info > div:nth-child(2) {
    font-size: 14px;
    color: #606266;
    margin-top: 10px;
}

div.prod-info > div:nth-child(3) {
    margin-top: 10px;
    font-size: 15px;
}

/* 收货地址 */

div.order-body > div:nth-child(3) {
    padding: 20px 10px 20px 10px;
    border-bottom: 1px solid rgba(128, 128, 128, 0.5);
}

div.order-body > div:nth-child(3) > h3 {
    color: #47494d;
}

div.rec-address {
    /*width: 80%;*/
    margin: 20px auto 10px;
}

div.address-func {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}

/* 付款数额 */
div.order-body > div:nth-child(4) {
    padding: 20px 10px 10px 10px;
    border-bottom: 1px solid rgba(128, 128, 128, 0.5);
}

div.order-body > div:nth-child(4) > h3 {
    color: #47494d;
}

div.order-body > div:nth-child(4) > div {
    display: flex;
    justify-content: space-between;
    margin: 20px 10px 10px 10px;
}

/* 条款 */
div.order-body > div:nth-child(5) {
    padding: 20px 10px 10px 10px;
    border-bottom: 1px solid rgba(128, 128, 128, 0.5);
}

div.order-body > div:nth-child(5) > h3 {
    color: #47494d;
}

div.order-body > div:nth-child(5) > ol {
    margin: 10px 0 10px 30px;
    font-size: 15px;
    color: #606266;
}

/* 下单 */
div.order-body > div:nth-child(6) {
    padding: 20px 10px 10px 10px;
    display: flex;
    justify-content: center;
    margin-bottom: 30px;
}

div.order-body > div:nth-child(6) > span {
    padding: 10px 30px 10px 30px;
    border-radius: 10px;
    color: white;
    background-color: #eea221;
    cursor: pointer;
    transition: .2s;
}

div.order-body > div:nth-child(6) > span:hover {
    background-color: rgba(238, 162, 33, 0.8);
}

div.addresses {
    width: 90%;
    margin: 0 auto 0;
}

div.address-info {
    margin-bottom: 20px;
    padding: 10px;

    border: 1px solid #3ec466;
    border-radius: 15px;
}

div.address-info-row {
    margin-top: 5px;
    font-size: 14px;
    color: #606266;
    position: relative;
}

span.select-address-but {
    padding: 2px 20px 2px 20px;
    cursor: pointer;
    transition: .2s;
    color: #368d93;
}

span.select-address-but:hover {
    background-color: #368d93;
    color: white;
}



</style>