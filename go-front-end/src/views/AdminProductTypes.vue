<template>
    <div class="body">
        <h2>商品类型修改</h2>

        <br><br>

        <div v-for="(_, index) in prodTypes" :key="index" class="prod-type">
            <el-row>
                <el-col :span="6">类型名</el-col>
                <el-col :span="14">
                    <el-input
                        v-model="prodTypes[index].tName"
                        maxlength="50"
                        show-word-limit></el-input>
                </el-col>
            </el-row>

            <el-row>
                <el-col :span="6">类型简介</el-col>
                <el-col :span="14">
                    <el-input
                        type="textarea"
                        v-model="prodTypes[index].subText"
                        maxlength="50"
                        show-word-limit></el-input>
                </el-col>
            </el-row>

            <el-row>
                <el-col :span="6">类型详情</el-col>
                <el-col :span="14">
                    <el-input type="textarea" :rows="4" v-model="prodTypes[index].description"></el-input>
                </el-col>
            </el-row>
            <el-row>
                <el-col :offset="6">
                    <el-button type="primary" @click="updateProductType(index)">修改</el-button>
                </el-col>
            </el-row>
        </div>


    </div>
</template>

<script>

import {Post} from "@/assets/js/http";
import {OkNotify, ErrorNotify} from "@/assets/js/message";

export default {
    name: "AdminUpdateProdType",

    created() {
        this.$store.dispatch("getProductTypes")
        this.prodTypes = this.$store.getters.getAllProductTypes
    },

    data() {
        return {
            prodTypes: null
        }
    },

    methods: {

        /**
         * type ProductType struct {
         * 	ProdTId     int    `json:"prodTypeId"`
         * 	TName       string `json:"tName"`
         * 	SubText     string `json:"subText"`
         * 	Description string `json:"description"`
         * }
         */

        updateProductType(index){
            let prodType = this.prodTypes[index]
            Post("/v1/admin/types").setJson(prodType).setHeader({
                "Admin-Token": sessionStorage.getItem("adminToken")
            }).send(
                ok => {
                    if(ok.code === 200){
                        OkNotify("成功", "修改商品类型成功")
                        this.$store.commit("removeTypesCache")
                        this.$store.commit("setProductTypes", this.prodTypes)
                    } else if (ok.code === 10302){
                        this.$router.push({
                            path: "/ceoLogin"
                        })
                    } else{
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

div.body {
    width: 50%;
    margin: 0 auto;
}

div.prod-type {
    margin: 20px 0 20px 0;
    padding: 20px;
    border: 1px solid grey;
    border-radius: 20px;
}

</style>