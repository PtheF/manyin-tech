<template>
    <div class="app">
        <div class="note">
            <div class="title">
                <h1 style="color: #11999e">蔓茵科技</h1>
                <h2 style="color: #11999e">
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;打破传统回收之囹圄
                </h2>
            </div>

            <div class="content">
                <p>责任、绿色、协调、创新</p>
                <p>一款绿色高效的现代城市垃圾回收企业</p>
                <p>致力于城市建筑垃圾回收与再利用</p>

                <h4 style="color: #48466d; margin-top: 40px">
                    打造中国建筑垃圾回收产业新世界。
                </h4>
            </div>
            <h3 style="margin-top: 50px; color: #48466d">
                蔓茵，攻克技术难关，拿下创新高地
            </h3>
        </div>

        <div class="login-body">
            <div class="login">
                <div style="margin: 40px 30px 40px 30px">
                    <div
                        style="
                            margin: 20px 0 20px 0;
                            border-bottom: 1px solid grey;
                        "
                    >
                        <div style="color: grey; cursor: pointer" @click="back">
                            返回
                        </div>
                    </div>

                    <div class="register-info">
                        <h2 class="login-title">注册</h2>
                        <div>
                            <el-form 
                                label-position="top"
                                :model='user'
                                ref='user'
                                :rules='rules'
                                label-width="60px">
                                <el-form-item label="电话或邮箱" prop='id'>
                                    <el-col :span='20'>
                                        <el-input 
                                            v-model="user.id" 
                                            placeholder="请输入邮箱或手机号" 
                                            prefix-icon='el-icon-user-solid'></el-input>
                                    </el-col>
                                </el-form-item>
                                <el-form-item label="验证码" prop='captcha'>
                                    <el-col :span='11'>
                                        <el-input 
                                            v-model='user.captcha' 
                                            placeholder="请输入验证码"
                                            prefix-icon='el-icon-message'></el-input>
                                    </el-col>
                                    <el-col :span='1'>
                                        &nbsp;
                                    </el-col>
                                    <el-col :span='6'>
                                        <captcha-button :capTo="user.id" @captcha='getCaptcha' :empty='emptyId'/>
                                    </el-col>
                                </el-form-item>

                                <el-form-item label="密码" prop='pw'>
                                    <el-col :span='20'>
                                        <el-input 
                                            type='password' 
                                            v-model='user.pw' 
                                            placeholder="请输入密码"
                                            prefix-icon='el-icon-key'
                                            show-password></el-input>
                                    </el-col>
                                </el-form-item>

                                <el-form-item label="重复密码" prop='repPw'>
                                    <el-col :span='20'>
                                        <el-input 
                                        type='password' 
                                        v-model='user.repPw' 
                                        placeholder="重复密码"
                                        prefix-icon='el-icon-key'
                                        show-password></el-input>
                                    </el-col>
                                </el-form-item>

                                <el-row>
                                    <el-col :span='20'>
                                        <div style='display: flex; justify-content: flex-end'>
                                            <el-button type='primary' @click='register("user")'>注册</el-button>
                                        </div>
                                    </el-col>
                                    
                                </el-row>
                            </el-form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import CaptchaButton from '../components/CaptchaButton.vue';
import { isEmpty } from "@/assets/js/utils"
import { Verify, REG } from "@/assets/js/verify"
import { OkNotify, ErrorNotify } from "@/assets/js/message"
import { Post } from "@/assets/js/http"

export default {
    components: { CaptchaButton },
    name: "Register",
    data() {

        var checkId = (rule, value, callback) => {
            if(isEmpty(value)){
                return callback(new Error("id不能为空"))
            }

            if(!Verify(value, REG.PHONE_REG, REG.EMAIL_REG)){
                return callback(new Error("输入不符合手机号或邮箱格式"))
            }   

            callback()
        }

        var checkPw = (rule, value, callback) => {
            if(isEmpty(value)){
                return callback(new Error("密码不能为空"))
            }

            if(!Verify(value, REG.PASS_REG)){
                return callback(new Error("密码不符合规则,长度8-20，包含数字和字母"))
            }

            callback()
        }

        var checkRepPw = (rule, value, callback) => {
            if(isEmpty(value)) {
                return callback(new Error("请重复"))
            }

            if(value !== this.user.pw){
                return callback(new Error("密码不一致"))
            }

            callback()
        }

        var checkCaptcha = (rule, value, callback) => {
            if(isEmpty(value)){
                return callback(new Error("请输入验证码"))
            }

            console.log("captcha validator: len(" + value.length + ")");

            if(!/\d{4}/.test(value)){
                return callback(new Error("验证码格式错误"))
            }

            callback()
        }

        return {

            rules: {
                id: [ { validator: checkId, trigger: "blur" } ],
                pw: [ { validator: checkPw, trigger: "blur" } ],
                captcha: [ { validator: checkCaptcha, trigger: "blur" } ],
                repPw: [ { validator: checkRepPw, trigger: "blur" } ]
            },

            user: {
                id: "",
                captcha: "",
                pw: "",
                repPw: ""
            },

            sendCaptcha: false,
        };
    },
    methods: {
        back() {
            window.history.back(-1);
        },

        getCaptcha(cap){
            this.$alert(`
                <div>
                    <p>
                        这里的验证码按理来说应该得发到你的手机或者邮箱,
                        但是因为我们太穷了，买不起服务，只能用这种办法给你验证码
                    </p>
                    <p> :( </p>
                    <strong>${cap}</strong>
                </div>
            `, '验证码', {
                dangerouslyUseHTMLString: true
            })

            this.sendCaptcha = true
        },

        register(formName){
            this.$refs[formName].validate((validate) => {
                if(validate){
                    if(!this.sendCaptcha){
                        ErrorNotify("哈？", "你没发验证呐")
                        return
                    }

                    Post('/v1/register').setForm({
                        id: this.user.id,
                        pw: this.user.pw,
                        cap: this.user.captcha
                    }).send(
                        ok => {
                            if(ok.code === 200){
                                OkNotify("账户", "账户创建成功")

                                setTimeout(() => {
                                    sessionStorage.setItem("newUserId", this.user.id)
                                    sessionStorage.setItem("newUserPw", this.user.pw)
                                    this.$router.replace({
                                        path: "/login"
                                    })
                                }, 1000)
                                
                            }else{
                                ErrorNotify("错误", ok.message)
                            }
                        },
                        
                        err => {
                            ErrorNotify("服务器故障，稍后再试")
                            ErrorNotify(err)
                        }
                    )

                }else{
                    console.log("error submit");
                    return
                }
            })
        },

        emptyId() {
            ErrorNotify("啥？", "您要给发验证码呀？？")
        }
    },
};
</script>

<style scoped>
div.app {
    width: 100vw;
    height: 100vh;

    background-image: url("~@/../public/green.svg");
    background-repeat: no-repeat;
    background-size: cover;
    background-position: 150px 0;

    position: relative;
    overflow: hidden;
    display: flex;
}

div.note {
    top: 10%;
    left: 10%;
    background: transparent;
    position: relative;
    height: auto;

    width: 24%;
}

div.content {
    margin-top: 70px;
    margin-left: 60px;
}

p {
    margin-top: 30px;

    color: rgb(160, 160, 160);
}

div.title {
    padding: 10px;
    /* background: linear-gradient(to right, #30e3ca, transparent) */
}

div.login-body {
    position: absolute;
    right: 15%;
    height: 100%;
    width: 26%;
    background-color: rgba(242, 255, 251, 0.7);
}

div.login {
    /* width: 90%; */
    height: 100%;
    /* margin-left: 20px; */
    /* border-right: 1px solid grey; */
    overflow: hidden;
}

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.login-body {
        right: 10%;
        width: 35%;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.login-body {
        right: 5%;
        width: 40%;
    }
    div.note {
        width: 40%;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.login-body {
        right: 0;
        width: 50%;
    }
    div.note {
        width: 50%;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.login-body {
        right: 0;
        width: 100%;
    }
    div.note {
        display: none;
    }

    div.app {
        background-position: -300px 0px;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.login-body {
        right: 0;
        width: 100%;
    }
    div.note {
        display: none;
    }
    div.app {
        background-position: -300px 0px;
    }
}

h2.login-title {
    margin-top: 20px;
    padding: 10px;
    /* background: linear-gradient(to right, #30e3ca, transparent) */
}

div.register-body {
    width: 90%;
    margin: 30px auto 0;
}

</style>