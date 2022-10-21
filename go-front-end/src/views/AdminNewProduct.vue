<template>
    <div class="body">
        <h2>插入商品</h2>

        <br><br>

        <h4>注意事项：</h4>

        <ol style="padding-left: 40px; font-size: 14px; color:#606266;">
            <li>添加图片只能通过 截图-粘贴 的方式，截完图以后直接在输入框里面粘贴即可</li>
            <li>粘贴完了以后会显示图片，删除输入框里面的内容就会删除图片</li>
            <li>为啥要这么弄不能直接上传？最早是方便我直接从 word 文档里面直接复制图片然后粘过去，现在是纯粹懒得做文件上传</li>
            <li>图片如果太长的话会显示不全，并不是你截图没截全，不用担心</li>
            <li>图片粘贴上来看一看，太大的话就重新截一下</li>
        </ol>

        <br><br>

        <el-row>
            <el-col :span="6">商品名称</el-col>
            <el-col :span="16">
                <el-input v-model="prodName" placeholder="商品名称"></el-input>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品类别</el-col>
            <el-col :span="16">
                <el-select v-model="prodType" placeholder="请选择">
                    <el-option
                        v-for="item in prodTypes"
                        :key="item.prodTypeId"
                        :label="item.tName"
                        :value="item.prodTypeId">
                    </el-option>
                </el-select>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品描述</el-col>
            <el-col :span="16">
                <el-input v-model="prodDesc" placeholder="商品描述" type="textarea"></el-input>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品图片</el-col>
            <el-col :span="16">
                <image-paste @getImg="getProdAvatar"></image-paste>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品参数图</el-col>
            <el-col :span="16">
                <image-paste @getImg="getProdDigital"></image-paste>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品每单位价格（元）</el-col>
            <el-col :span="16">
                <el-input-number v-model="prodPrice" placeholder="请输入价格"></el-input-number>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :span="6">商品售卖单位</el-col>
            <el-col :span="16">
                <el-input v-model="prodUnit" placeholder="输入售卖单位，比如 '个', '吨' 一类"></el-input>
            </el-col>
        </el-row>
        <br>

        <el-row>
            <el-col :offset="6">
                <el-button type="primary" @click="addProduct">提交</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<script>

import ImagePaste from "@/components/ImagePaste";
import { Put } from "@/assets/js/http"
import { OkNotify, ErrorNotify } from "@/assets/js/message";

export default {
    name: "Admin",
    components: {ImagePaste},

    created() {
        this.$store.dispatch("getProductTypes")

        this.prodTypes = this.$store.getters.getAllProductTypes

        console.log(this.prodTypes)
    },

    data() {
        return {
            prodName: "",
            prodDesc: "",

            prodType: null,

            prodTypes: [],

            prodAvatar: "",
            prodDigital: "",

            prodPrice: 0,
            prodUnit: ""
        }
    },
    methods: {
        getProdAvatar(data){
            this.prodAvatar = data
        },

        getProdDigital(data) {
            this.prodDigital = data
        },

        addProduct() {
            const prod = {
                prodName: this.prodName,
                prodDesc: this.prodDesc,
                prodTypeId: this.prodType,
                prodAvatar: this.prodAvatar,
                prodDigital: this.prodDigital,
                prodPrice: this.prodPrice * 100,
                prodUnit: this.prodUnit
            }

            console.dir(prod)

            Put("/v1/admin/product").setJson(prod).setHeader({
                "Admin-Token": sessionStorage.getItem("adminToken")
            }).send(
                ok => {
                    if(ok.code === 200) {
                        OkNotify("成功", "添加成功")

                        this.prodName = ""
                        this.prodDesc = ""
                        this.prodType = null
                        this.prodTypes = []
                        this.prodAvatar = ""
                        this.prodDigital = ""
                        this.prodPrice = 0
                        this.prodUnit = ""

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

div.body {
    width: 50%;
    margin: 0 auto;
}

</style>