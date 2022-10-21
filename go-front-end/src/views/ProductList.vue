<template>
    <div>
        <header-nav/>
<!--            <h1>{{typeId}}</h1>-->
        <div class="body">
            <div class="two">
                <div class="second">
                    <img src="/logo.png">
                    <ul>
                        <li>
                            <span>{{ typeInfo.tName }}</span>
                        </li>
                        <li>
                            <p>{{ typeInfo.description }}</p>
                        </li>

                    </ul>
                </div>
            </div>
            <div class="box">
                <div class="item" v-for="(item, index) in prods" :key="index" @click="productDetail(item.prodId)">
                    <img :src="item.prodAvatar"><br>
                    <div class="info">
                        <div>{{ item.prodName }}</div>
                        <div>{{ item.prodPrice / 100 }} / {{ item.prodUnit }}</div>
                    </div>
                </div>
            </div>

            <div class="page">
                <el-pagination
                    :page-size="12"
                    layout="prev, pager, next"
                    :current-page="routerPage"
                    :page-count="totalPage"
                    @current-change="changePage">
                </el-pagination>
            </div>

        </div>
        <Footer style="position: relative;"></Footer>
    </div>
</template>

<script>
import HeaderNav from "@/components/HeaderNav";
import Footer from "@/components/Footer";
import { Get } from "@/assets/js/http"
import {ErrorNotify} from "@/assets/js/message";

export default {
    name: "ProductList",
    components: {
        HeaderNav, Footer
    },

    // 首次进入页面会触发 created 周期
    created(){
        console.log(`product list: type=${this.typeId}`)
        console.log(`product list: page=${this.routerPage}`)

        this.getProdList(this.typeId, this.routerPage)
    },

    // mounted(){
    //     console.log("product list mounted")
    // },

    // 在当前页面上进入其他页面触发 updated，且首次进入的时候不会触发这个生命周期
    // 最新进展，这样也不行，因为只要一发请求，就会触发 update，然后 update 又会发送请求，
    // 然后无限循环
    updated() {
        // console.log("product list updated")
        // console.log(`product list: type=${this.typeId}`)
        // console.log(`product list: page=${this.routerPage}`)
        // this.getProdList()
    },

    data() {
        return {

            totalPage: 1,
            currentPage: 1,
            // prods: [
            //     {
            //         prodId: "00001",
            //         name: "绿绿的大章鱼",
            //         avatar: require("@/assets/img/item0001.png"),
            //         price: 100,
            //         unit: "个"
            //     },
            //     {
            //         prodId: "00002",
            //         name: "棕色的大熊熊",
            //         avatar: require("@/assets/img/item0002.png"),
            //         price: 200,
            //         unit: "个"
            //     }
            // ]

            prods: []
        }
    },

    methods: {
        getProdList(typeId, page){
            Get(`/v1/shop/list/${typeId}/${page}`).send(
                ok => {
                    if(ok.code === 200){
                        this.prods = ok.data.products
                        this.currentPage = ok.data.currentPage
                        this.totalPage = ok.data.totalPage
                    }else{
                        ErrorNotify("错误", ok.message)
                    }
                },

                err => {
                    ErrorNotify("网络错误", err)
                }
            )
        },

        // nextPage(){
        //     console.log("nextPage")
        // },
        //
        // prevPage(){
        //     console.log("prevPage")
        // },

        changePage(page){
            console.log("changePage")
            console.log(page)

            this.$router.push({
                path: `/shop/list/${this.typeId}/${page}`
            })
        },

        productDetail(id) {
            this.$router.push({
                path: "/item/" + id
            })
        }

    },

    computed: {
        typeInfo() {
            return this.$store.getters.getTypeInfo(this.typeId)
        },

        typeId() {
            return Number.parseInt(this.$route.params.typeId)
        },

        routerPage() {
            return Number.parseInt(this.$route.params.page)
        },
    },

    watch: {
        "typeId": function(newValue, oldValue) {
            console.log(`type id change, old = ${oldValue}, new = ${newValue}`)

            this.getProdList(newValue, 1)
        },

        "routerPage": function(newValue, oldValue) {
            console.log(`page change, old = ${oldValue}, new = ${newValue}`)

            this.getProdList(this.typeId, newValue)
        }
    }
}
</script>

<style scoped>


@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    .box {
        width: 75% !important;
    }

    .two .second ul li:nth-child(2) p {
        width: 60% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    .box {
        width: 80% !important;
    }

    .two .second ul li:nth-child(2) p {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    .box {
        width: 90% !important;
    }

    .two .second ul li:nth-child(2) p {
        width: 80% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    .box {
        width: 95% !important;
    }
    .item {
        width: 90% !important;
        margin: 10px auto 10px !important;
    }

    .two .second ul li:nth-child(2) p {
        width: 90% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    .box {
        width: 95% !important;
    }
    .item {
        width: 90% !important;
        margin: 10px auto 10px !important;
    }

    .two .second ul li:nth-child(2) p {
        width: 90% !important;
    }
}

div.body {
    min-height: 100vh;
    /*height: 100vh;*/
    width: 100vw;
    /*background-color: #f4f4f4*/
}

.two{
    /* margin-top:44px; */
    /* height:300px; */
    height: auto;
    background: linear-gradient(-150deg, #222222 15%, #373737 70%, #3c4859 94%);
    overflow: hidden;
    /* position: relative; */
}
.box{
    width: 70%;
    /* height: 650px; */
    height: auto;
    margin: 0 auto;
    color: azure;

    padding-bottom: 40px;

    /* 用 flex 布局让里面的 item 均匀排列 */
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
}
.item{
    /* padding: 15px; */
    /* float:left; */
    /* margin: 50px; */
    margin-top: 50px;
    margin-bottom: 10px;
    width: 26%;
    border: 2px solid white;
    transition: .2s;
    padding: 10px 10px 20px 10px;
    height: 250px;

    box-shadow: 0 5px 15px 0 rgba(72, 72, 72, 0.1);
    border-radius: 20px;

    /* 鼠标悬浮时样式 */
    cursor: pointer;
}

.item:hover {
    /* transform: scale(1.05); */
    box-shadow: 0 5px 15px 0 rgba(0, 0, 0, 0.2);
}

.two .second{
    /*width: 1200px;*/
    /* height: 300px; */
    margin: 50px auto 30px;
    text-align: center;
    /* overflow: hidden; */
}
.one .head img{
    width: 40px;
    height: 40px;
    margin: 0 0;
    background-color: whitesmoke;
}
.two .second img{
    height:200px;
    /* weight:300px; */
    margin-top: 10px;
    margin-left: 5px;
    text-align: center;
}
.one .head ul{
    list-style: none;
    position: absolute;
    right: 0;
    height: 0;
    top: 20%;
    margin-top: -10px;
    font-size: 12px;
}
.two .second ul{
    list-style: none;
    margin-top: 50px;
    /* height: 30px; */
    top: 50%;
    /*margin-top: -10px;*/
    margin-bottom: 50px;
    font-size: 20px;
}
.one .head ul li{
    float: left;
    /*border-right: 2px solid #6C6C6C;*/
}
.one .head ul li span{
    color: #6C6C6C;
    margin: 0 auto;
}
.two .second ul li:nth-child(1) span{
    color: white;
    /* font-style: oblique; */
    /* margin: 0 45px; */
    font-size: 40px;

}

.two .second ul li:nth-child(2) p {
    margin: 40px auto 0;
    color: grey;
    width: 50%;

    font-size: 16px;
}

.one .head ul li a{
    text-decoration: none;
}
.item a{
    text-decoration: none;
}
.item span{
    font-size: 15px;
    color: #6C6C6C;
}
.item img{
    display: block;
    width: 95%;
    height:200px;
    margin:0 auto;
    border-radius: 20px;
}
a:hover{
    background-color: skyblue;
}

div.item > .info {
    display: flex;
    justify-content: space-between;
    /* justify-content: center; */
    color: black;
    width: 95%;
    margin: 5px auto;
}

div.item > .info > div:nth-child(1) {
    color: #606266;
}

div.item > .info > div:nth-child(2) {
    padding: 3px 20px 3px 20px;
    background-color: #f0ad4e;
    color: white;
    border-radius: 5px;
}

div.page {
    text-align: center;
    margin: 10px auto 30px;
}

</style>