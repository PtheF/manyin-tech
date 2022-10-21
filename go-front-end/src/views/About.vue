<template>
    <div class="about">
        <HeaderNav/>
        <div class="content">
            <h1>关于我们</h1>
            <hr>
            <h3><i class="el-icon-caret-right"/>关于项目</h3>
            <p>
                开发 “城市矿山”智能回收运营商，组建“蔓茵回收-建筑公司-小区物业”三位一体区域管理层，在建筑垃圾回收、管理、再生等产业链的上下游和再生品的产销过程中实现全方位智能管理，每一个步骤都有专人负责，实现用户放心，产品贴心的目标。
            </p>

            <h3><i class="el-icon-caret-right"/>关于技术</h3>
            <p>
                本网站技术比较简单：前端 Vue，后端 beego，数据库 MySQL，还有 Redis 和 JWT 这些的，没了。
                而且本项目未来计划开源，等着 Github 地址就行了。
            </p>

            <h3><i class="el-icon-caret-right"/>关于开发者</h3>
            <div style="text-align: center; margin-top: 30px;">
                <img :src="sharkbee" alt="" width="200px" height="auto">
            </div>
            <div style="text-align: center; color: grey; margin-bottom: 50px;">
                这位就是"鲨bee"开发者
            </div>

        </div>
        <Footer class="about-footer"/>
    </div>
</template>

<script>

import HeaderNav from "@/components/HeaderNav"
import Footer from "@/components/Footer"
// import sharkbee from "@/../public/sharbee.jpg"

export default {

    name: "About",
    components: {
        HeaderNav, Footer
    },
    data() {
        return {
            sharkbee: require("@/../public/sharkbee.jpg")
        }
    }

};
</script>

<style scoped>

    @media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
        div.content {
            width: 90% !important;
            padding: 0 20px 0 20px;
        }

        .about-footer {
            position: relative;
        }
    }

    @media all and (max-device-aspect-ratio: 1/1) {
        div.content {
            width: 90% !important;
            padding: 0 20px 0 20px;
        }

        .about-footer {
            position: relative;
        }
    }

    div.about{
        /* position: relative; */
        /* overflow: hidden; */
        color: #606266;
    }

    div.content{
        width: 60%;
        /* border: 1px solid grey; */
        margin: 20px auto 0;
    }

    div.content > h1 {
        text-align: center;
        margin-bottom: 10px;
    }

    div.content > p {
        text-indent: 40px;
        margin-top: 10px;
    }

    div.content > h3 {
        margin-top: 20px;
    }

</style>
