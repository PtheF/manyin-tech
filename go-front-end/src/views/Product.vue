<template>
    <div>
        <header-nav></header-nav>
        <div class="body">
            <div class="info">
                <div class="prod-img" :style="{backgroundImage: `url(${productDetail.prodAvatar})`}"></div>

                <div class="prod-info">
                    <div>
                        <h3>{{productDetail.prodName}}</h3>

                        <p>{{ productDetail.prodDesc }}</p>
                    </div>


                    <div>
                        <span style="color: #606266;">售价</span>：
                        <span style="color: #ff7225; font-size: 20px;">
                            <strong>￥{{ productDetail.prodPrice / 100 }} / {{ productDetail.prodUnit }}</strong>
                        </span>
                    </div>

                    <div>
                        <el-row>
                            <el-col :span="8">IP</el-col>
                            <el-col :span="12">蔓茵科技</el-col>
                        </el-row>

                        <el-row>
                            <el-col :span="8">品牌</el-col>
                            <el-col :span="12">蔓茵科技</el-col>
                        </el-row>
                    </div>

                    <div>
<!--                        <div>-->
<!--                            <el-input-number v-model="buyNum" style="width: 100px;" controls-position="right" :min="1" :max="10"></el-input-number>-->
<!--                        </div>-->
<!--                        <br/>-->
                        <span @click="buy">立即购买</span>&nbsp;&nbsp;&nbsp;&nbsp;
                        <el-input-number v-model="buyNum" style="width: 100px;" controls-position="right" :min="1" :max="999"></el-input-number>
                    </div>

                    <div>
                        <el-row>
                            <el-col :span="6">保障</el-col>

                            <el-col :span="16">
                                <i class="el-icon-circle-check"/>正品保障&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                                <i class="el-icon-circle-check"/>不支持无理由退货
                            </el-col>
                        </el-row>
                    </div>
                </div>
            </div>

            <div class="detail">
                <div><strong>商品详情</strong></div>
                <img :src="productDetail.prodDigital" alt="">
            </div>

            <div class="notice">
                <div>
                    <strong>购买须知</strong>
                </div>

                <p>
                    1. 因为本项目只是大创展示用，并没有实际业务，所以不会发货
                </p>

                <p>
                    2. 当然，您也没法付款
                </p>

                <p>
                    3. 所有产品的参数不一定正确，如有错误请谅解一下
                </p>

                <p>
                    4. 作为商家，这个购买须知才3条是不是显得不专业？
                </p>
            </div>
        </div>
        <Footer style="position: relative;"></Footer>

    </div>
</template>

<script>
import HeaderNav from "@/components/HeaderNav";
import Footer from "@/components/Footer";
import VMask from "@/components/VMask";
import {Get} from "@/assets/js/http";
import {ErrorNotify} from "@/assets/js/message";
import loginAuth from "@/assets/js/loginAuth";
import {debounce} from "@/assets/js/utils";

export default {
    name: "Product",
    components: {VMask, Footer, HeaderNav},

    created() {
        this.prodId = this.$route.params.prodId;

        Get(`/v1/shop/item/${this.prodId}`).send(
            ok => {
                if(ok.code === 200) {
                    this.productDetail = ok.data
                }else{
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
            prodId: null,

            productDetail: {
                prodId: '00001',
                prodName: "绿绿的大章鱼",
                prodUnit: "个",
                prodPrice: 100,
                prodAvatar: require("@/assets/img/item0001.png"),
                prodDesc: `特别绿的大章鱼，我也不知道为啥这么绿，可能是炒股炒的吧，赔了不少钱，反正就是挺厉害的`,
                prodDigital: require("@/assets/img/item0001-attr.png")
            },

            buyNum: 1
        }
    },

    methods: {

        buy: debounce(function(){
            loginAuth(
                () => {
                    Get("/v1/order/token").send(
                        ok => {
                            if(ok.code === 200){
                                let token = ok.message
                                let createOrderPage = "/newOrder?token=" + token

                                const order = {
                                    orderToken: token,
                                    buyNumber: this.buyNum,
                                    prodDetail: this.productDetail
                                }

                                this.$store.commit("setOrder", order)

                                this.$router.push({
                                    path: createOrderPage
                                })
                            }
                        },

                        err => {
                            ErrorNotify("网络错误", err)
                        }
                    )
                },

                () => {
                    this.$router.replace({
                        path: "/login?uri=" + this.$route.path
                    })
                },

                () => {
                    this.$router.replace({
                        path: "/login?expire=true&uri=" + this.$route.path
                    })
                }
            )
        })
    }
}
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 60vw !important;
    }

    div.prod-img {
        width: 28vw !important;
        height: 28vw !important;
    }

    div.prod-info {
        width: 28vw !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 70vw !important;
    }

    div.prod-img {
        width: 33vw !important;
        height: 33vw !important;
    }

    div.prod-info {
        width: 33vw !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 80vw !important;
    }

    div.prod-img {
        width: 38vw !important;
        height: 38vw !important;
    }

    div.prod-info {
        width: 36vw !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 95vw !important;
    }

    div.info {
        flex-wrap: wrap !important;
    }

    div.prod-img {
        width: 90vw !important;
        height: 90vw !important;
    }

    div.prod-info {
        width: 90vw !important;
        margin-left: 0 !important;
    }

    div.prod-info > div:nth-child(4) {
        text-align: center !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.body {
        width: 95vw !important;
    }

    div.info {
        flex-wrap: wrap !important;
    }

    div.prod-img {
        width: 90vw !important;
        height: 90vw !important;
    }

    div.prod-info {
        width: 90vw !important;
        margin-left: 0 !important;
    }

    div.prod-info > div:nth-child(4) {
        text-align: center !important;
    }
}

div.body {
    min-height: 90vh;
    margin: 0 auto;
    width: 50vw;
    /*border: 1px solid grey;*/
    padding: 20px;
}

div.info {
    display: flex;

    padding-bottom: 20px;
}

div.prod-img {
    background-repeat: no-repeat;
    background-size: 100% auto;
    width: 23vw;
    /*height: auto;*/
    height: 23vw;
}

div.prod-info {
    margin-left: 40px;
    width: 22vw;
}

div.prod-info > div:nth-child(1) > p {
    font-size: 15px;
    color: #606266;

    margin-top: 20px;
}

div.prod-info > div:nth-child(2) {
    margin-top: 20px;
    /*background: linear-gradient(to right, yellow, white);*/
    /*background-color: #f1f17f;*/
    background-color: rgba(241, 241, 127, 0.3);
    padding: 3px 10px 3px 30px;
}

div.prod-info > div:nth-child(3) {
    margin-top: 20px;
    color: #606266;
    font-size: 15px;
}

div.prod-info > div:nth-child(4) {
    margin-top: 60px;
}

div.prod-info > div:nth-child(4) > span {
    padding: 10px 40px 10px 40px;
    color: white;
    background-color: orange;

    cursor: pointer;
}

div.prod-info > div:nth-child(5) {
    margin-top: 30px;
    font-size: 15px;
    color: #606266;
}

div.detail > div, div.notice > div {
    padding: 10px 0 10px 40px;
    color: black;
    background: rgba(227, 227, 227, 0.7);
}

div.detail > img {
    width: 100%;
    height: auto;
}

div.notice > p {
    margin-top: 10px;
    text-indent: 20px;
    color: #606266;
    font-size: 16px;
}


</style>