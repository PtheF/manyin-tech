<template>
    <div class="body">
        <h3 style="text-align: center;">管理员登录</h3>
        <br><br>
        <ol>
            <li>管理员页面入口</li>
            <li>为啥这个管理员页面写的跟屎一样？因为一开始就没想过做这个东西</li>
            <li>为啥突然想做了？为啥？往后让我改商品我可受不了，你们自己改吧</li>
            <li>这个页面的安全性做的可谓是非常差，随时都有被攻击的风险</li>
        </ol>

        <br><br>

        <el-row>
            <el-col :offset="8" :span="8">
                <el-input v-model="pw" placeholder="输入密码" type="password"></el-input>
            </el-col>
        </el-row>

        <br>

        <el-row>
            <div style="display: flex; justify-content: center;">
                <el-button type="primary" @click="adminLogin">登录</el-button>
            </div>

        </el-row>



    </div>
</template>

<script>
import {ErrorNotify, OkNotify} from "@/assets/js/message";
import {Post} from "@/assets/js/http";

export default {
    name: "AdminLogin",
    data() {
        return {
            pw: ""
        }
    },

    methods: {
        adminLogin() {
            if(this.pw === ""){
                return
            }

            Post("/v1/admin/login").setForm({
                pw: this.pw
            }).send(
                ok => {
                    if(ok.code === 200){

                        console.log(ok)

                        sessionStorage.setItem("adminToken", ok.message)

                        this.$router.replace({
                            path: "/ceo"
                        })
                    }else{
                        ErrorNotify("错误", "登录失败")
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
    margin: 50px auto 0;
    width: 50%;
}

div.body > ol {
    width: 60%;
    margin: 20px auto 0;
}

</style>