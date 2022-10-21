<template>
    <div class="body">
        <h2>删除商品</h2>
        <span>是的，只能删除，不能修改，要修改你就删除以后再添加，别问为啥不能修改，就是懒</span>

        <br><br>

        <div class="prod" v-for="(item, index) in prods" :key="index">
            <el-row>
                <el-col :span="6">
                    <div class="prod-avatar" :style="{backgroundImage: `url(${item.prodAvatar})`}"></div>
                </el-col>
                <el-col :span="14">
                    <h4>{{item.prodName}}</h4>
                    <p>价格: ￥{{item.prodPrice / 100}}/{{item.prodUnit}}</p>

                    <div>
                        <el-button type="danger" @click="delProd(index)">删除</el-button>
                    </div>
                </el-col>
            </el-row>

        </div>

    </div>
</template>

<script>
import {Get, Delete} from "@/assets/js/http";
import {ErrorNotify, OkNotify} from "@/assets/js/message";
import {removeArray} from "@/assets/js/utils";


export default {
    name: "AdminProductPage",

    created() {
        this.getProdList(this.typeId, this.routerPage)
    },

    data() {
        return {
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

        delProd(index) {
            Delete("/v1/admin/product/" + this.prods[index].prodId).setHeader({
                "Admin-Token": sessionStorage.getItem("adminToken")
            }).send(
                ok => {
                    if(ok.code === 200){
                        removeArray(this.prods, index)
                        OkNotify("成功", "删除完成")
                    }else if(ok.code === 10302){
                        this.$router.push({
                            path: "/ceoLogin"
                        })
                    }else{
                        ErrorNotify("错误", ok.message)
                    }
                },
                err => {
                    ErrorNotify("网络错误", err)
                }
            )
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

div.prod {
    margin: 20px 0 10px 0;
    padding: 10px;
    border: 1px solid grey;
    border-radius: 20px;
}

div.prod-avatar {
    height: 200px;
    width: 100%;
    background-size: auto 100%;
    background-repeat: no-repeat;
    overflow: hidden;
}

</style>