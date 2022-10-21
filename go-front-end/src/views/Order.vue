<template>
    <div style="position: relative;">
        <HeaderNav></HeaderNav>
        <div class="order-body">
            <div class="order-logo">
                <div>
<!--                    <img style="height: 60px; width: 60px; line-height: 90px;" src="/logo.png" alt="">-->
                    <h1>蔓茵商城</h1>
                    <div>我的订单</div>
                </div>
            </div>

            <div class="body">

                <div>
                    <h3><i class="el-icon-s-order"></i>&nbsp;&nbsp;全部订单</h3>
                </div>

                <div class="order-list">

                    <el-empty description="暂无订单" v-show="orderList.length === 0"></el-empty>

                    <div class="each-order" v-for="(item, index) in orderList" :key="index" v-show="orderList.length !== 0">
                        <div>
                            <span>
                                创建时间：{{item.createTime}} &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                订单号：{{item.orderIdStr}}
                            </span>
                        </div>

                        <div class="order-info">
                            <!-- product info -->
                            <div>
                                <img :src="item.prodAvatar" alt="">
                                <p class="content-center">
                                    {{item.prodName}}
                                </p>
                            </div>

                            <!-- product number -->
                            <div class="content-center">
                                x {{item.buyNum}}
                            </div>

                            <!-- addressee -->
                            <div class="content-center">

                                <el-popover
                                    width="200"
                                    placement="left"
                                    trigger="hover">

                                    <div style="padding: 5px;">
                                        <div><span style="color: black;">收件人: </span>{{item.addressee}}</div>
                                        <div style="color: black;">地址: </div>
                                        <p>{{item.address}}</p>
                                        <div><span style="color: black;">联系电话: </span>{{item.addresseePhone}}</div>
                                    </div>

                                    <div slot="reference"><i class="el-icon-user-solid"/>{{item.addressee}}</div>
                                </el-popover>

                            </div>

                            <!-- price -->
                            <div class="content-center">
                                ￥ {{item.totalPrice / 100}}
                            </div>

                            <!-- status -->
                            <div class="content-center">
                                <div style="display: block">
                                    <div :style="{ color: `${item.status === 1 ? '#606266' : '#ff5050'}`}">{{orderStatus(index)}}</div>
                                    <div @click="orderDetail(index)">查看详情</div>
                                </div>
                            </div>
                        </div>
                        <div>
                            <el-row>
                                <el-col :offset="10" :span="6">
                                    <div :style="{ color: `${item.status === 1 ? '#606266' : '#ff5050'}`}">{{orderStatus(index)}}</div>
                                </el-col>
                                <el-col :span="8">
                                    <span @click="orderDetail(index)" style="cursor: pointer;">查看详情</span>
                                </el-col>
                            </el-row>
                        </div>
                    </div>
                </div>

            </div>
        </div>

        <v-mask :show="showOrderDetail" @close="closeDetail">
            <div class="mask-order">

                <div class="mask-order-info" :style="{border: `1px solid ${currentOrder.status === 1 ? '#6bd56b' : 'red'}`}">
                    <div>
                        <el-row>
                            <el-col :span="8">订单号: </el-col>
                            <el-col :span="14">{{ currentOrder.orderIdStr }}</el-col>
                        </el-row>
                        <el-row>
                            <el-col :span="8">创建日期: </el-col>
                            <el-col :span="14">{{ currentOrder.createTime }}</el-col>
                        </el-row>
                        <el-row>
                            <el-col :span="8">数量: </el-col>
                            <el-col :span="14">x {{ currentOrder.buyNum }}</el-col>
                        </el-row>
                        <el-row>
                            <el-col :span="8">价格: </el-col>
                            <el-col :span="14">￥{{ currentOrder.totalPrice / 100 }}</el-col>
                        </el-row>
                        <el-row>
                            <el-col :span="8">订单状态: </el-col>
                            <el-col :span="14">
                                <div :style="{color: `${currentOrder.status === 1 ? '#39c939' : '#ff5050'}`}">
                                    {{orderStatus(curOrderDetail)}}
                                </div>
                                <div v-if="currentOrder.status === 0" class="mask-pay">
                                    <span style="cursor: pointer;" @click="payOrder(currentOrder.orderIdStr)">点击支付</span>
                                </div>
                            </el-col>
                        </el-row>
                    </div>
                </div>

                <div class="mask-prod-info">
                    <img :src="currentOrder.prodAvatar" alt="">
                    <p>
                        {{ currentOrder.prodName }}
                    </p>
                </div>

                <div class="addressee">
<!--                    <h4><i class="el-icon-user-solid"></i>收件人</h4>-->
                    <el-row style="margin-top: 5px;">
                        <el-col :span="8">收件人</el-col>
                        <el-col :span="14">{{ currentOrder.addressee }}</el-col>
                    </el-row>
                    <el-row style="margin-top: 5px;">
                        <el-col :span="8">地址</el-col>
                        <el-col :span="14">
                            <p>{{ currentOrder.address }}</p>
                        </el-col>
                    </el-row>
                    <el-row style="margin-top: 5px;">
                        <el-col :span="8" style="height: 23px; line-height: 23px;">联系电话</el-col>
                        <el-col :span="14" style="height: 23px; line-height: 23px;">{{ currentOrder.addresseePhone }}</el-col>
                    </el-row>
                </div>
            </div>

        </v-mask>

        <Footer style="position: relative;"></Footer>
    </div>
</template>

<script>

import Footer from "@/components/Footer";
import HeaderNav from "@/components/HeaderNav";
import VMask from "@/components/VMask";
import {ErrorNotify} from "@/assets/js/message";
import {Get} from "@/assets/js/http";

export default {
    name: "Order",
    components: {
        VMask,
        Footer, HeaderNav
    },

    created() {
        Get("/v1/order/list").send(
            ok => {
                if(ok.code === 200) {
                    console.log("user orders")
                    console.log(ok.data)

                    this.orderList = ok.data
                }else {
                    ErrorNotify("错误", ok.message)
                }
            },

            err => {
                ErrorNotify("网络错误", err)
            }
        )
    },

    data() {
        return {

            showOrderDetail: false,
            curOrderDetail: -1,

            orderList: []
        }
    },
    methods: {
        orderDetail(index){
            this.showOrderDetail = true
            this.curOrderDetail = index
        },

        closeDetail(){
            this.showOrderDetail = false
        },

        payOrder(orderIdStr){
            this.$router.push({
                path: `/pay/${orderIdStr}`
            })
        }
    },
    computed: {
        currentOrder() {
            if(this.curOrderDetail === -1) {
                return {

                }
            }
            return this.orderList[this.curOrderDetail]
        },

        /**
         * status:
         *   1 = 完成
         *   2 = 未支付
         */
        orderStatus() {
            return (orderIndex) => {

                if(this.orderList.length === 0 || orderIndex < 0) {
                    return "无订单"
                }

                let status = this.orderList[orderIndex].status

                if(status === 1)
                    return "已完成"
                if(status === 0)
                    return "未支付"
            }
        }
    }
}
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 80% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 90% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }

    div.each-order > div:nth-child(1) > span {
        font-size: 12px !important;
    }

    div.each-order > div:nth-child(3) {
        display: block !important;
    }

    div.order-info > div:nth-child(1) {
        width: 100% !important;
        border-right: none !important;
        border-bottom: 1px solid #bebebe;
        padding-bottom: 10px;
    }

    div.order-info > div:nth-child(n + 1){
        display: none;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }

    div.each-order > div:nth-child(1) > span {
        font-size: 12px !important;
    }

    div.each-order > div:nth-child(3) {
        display: block !important;
    }

    div.order-info > div:nth-child(1) {
        width: 100% !important;
        border-right: none !important;
        border-bottom: 1px solid #bebebe;
        padding-bottom: 10px;
    }

    div.order-info > div:nth-child(n + 1){
        display: none;
    }
}

.content-center {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
}

div.order-body {
    min-height: 100vh;
    background-color: #f4f4f4;
    overflow: hidden;
    padding-bottom: 20px;
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

div.body {
    height: auto;
    width: 60%;
    margin: 20px auto 10px;
}

div.body > div:nth-child(1) {
    background-color: white;
    padding: 10px 10px 10px 20px;
    border-radius: 20px;
}

div.each-order {
    padding: 10px 10px 10px 20px;
    background-color: white;
    border-radius: 10px;
    margin-top: 15px;

    transition: .2s;
    box-shadow: 0 5px 15px 0 rgba(169, 169, 169, 0.1);
}

div.each-order:hover {
    box-shadow: 0 5px 15px 0 rgba(72, 72, 72, 0.1);
}

div.each-order > div:nth-child(1) > span {
    font-size: 14px;
    color: #606266;
}

div.each-order > div:nth-child(1) {
    padding-bottom: 6px;
    border-bottom: 1px solid #afafaf;
}

div.each-order > div:nth-child(3) {
    font-size: 14px;
    display: none;
}

div.order-info {
    font-size: 14px;
    /*padding: 15px 0 15px 0;*/
    margin: 15px 0 15px 0;
    display: flex;
    flex-wrap: wrap;
}

/* product info */
div.order-info > div:nth-child(1) {
    display: flex;
    width: 40%;
    border-right: 1px solid #afafaf;
    cursor: pointer;
}

div.order-info > div:nth-child(1) > img {
    width: 80px;
    height: 80px;
}

div.order-info > div:nth-child(1) > p {
    margin-left: 20px;
}

/* product number */
div.order-info > div:nth-child(2) {
    width: 10%;
    border-right: 1px solid #afafaf;
    color: #888c91;
}

/* addressee */
div.order-info > div:nth-child(3) {
    width: 15%;
    border-right: 1px solid #afafaf;
    cursor: pointer;
}

/* price */
div.order-info > div:nth-child(4) {
    width: 10%;
    border-right: 1px solid #afafaf;
    color: #888c91;
}

/* status */
div.order-info > div:nth-child(5) {
    width: 15%;
}

div.order-info > div:nth-child(5) > div > div:nth-child(1) {
    text-align: center;
    border-bottom: 1px solid #b4b4b4;
    padding-bottom: 5px;
    color: #888c91;
}

div.order-info > div:nth-child(5) > div > div:nth-child(2) {
    margin-top: 5px;
    padding: 1px 10px 1px 10px;
    border: 1px solid grey;
    cursor: pointer;
    border-radius: 5px;
    background-color: #f5f5f5;
}


/* v-mask order detail */
div.mask-order {
    padding: 10px 20px 10px 20px;
    width: 85%;
    margin: 10px auto 10px;
}

/* v-mask order-info */
div.mask-order-info {
    padding: 12px 10px 12px 20px;
    /*border: 1px solid #d9d9d9;*/
    border-radius: 20px;
    box-shadow: 0 5px 15px 0 rgba(169, 169, 169, 0.1);
}

div.mask-order-info > div:nth-child(1) {
    font-size: 14px;
    color: #606266;
}

/* v-mask order-detail product info */
div.mask-prod-info {
    margin: 10px auto 0;
    display: flex;
    padding: 12px 10px 12px 20px;
    border: 1px solid #d9d9d9;
    border-radius: 20px;
    box-shadow: 0 5px 15px 0 rgba(169, 169, 169, 0.1);
    cursor: pointer;
}

div.mask-prod-info > img {
    width: 120px;
    height: 120px;
}

div.mask-prod-info > p {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-left: 20px;
}

/* v-mask addressee info */
div.addressee {
    padding: 12px 10px 12px 20px;
    border: 1px solid #d9d9d9;
    box-shadow: 0 5px 15px 0 rgba(169, 169, 169, 0.1);
    margin-top: 12px;
    border-radius: 20px;
    color: #606266;
    font-size: 14px;
}

div.mask-pay {
    margin-top: 5px;
    color: #3a708d;
}

</style>