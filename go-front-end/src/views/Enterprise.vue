<template>
    <div>
        <header-nav/>
            <div class="title">
                <h2>企业认证</h2>
            </div>

            <div class="moblie">
                <h2>Can not Access This Page By Cell Phone</h2>
            </div>
            
            <div class="body">
                <div>
                    <el-steps :active="finishStep" align-center :finish-status="finishStatus">
                        <el-step title="填写企业信息" index="1"></el-step>
                        <el-step title="审核认证信息" index="2"></el-step>
                        <el-step :title="'认证结果'" index="3"></el-step>
                    </el-steps>
                </div>

                <div class="enterprise-info">

                    <el-form 
                        ref="enterpriseInfo" 
                        :model="enterpriseInfo" label-width="150px"
                        label-position="left"
                        :rules="enterpriseRules">

                        <div style="padding-bottom: 30px;">
                            <h3 style="margin-bottom: 10px;">
                                <i class="el-icon-s-promotion"></i>&nbsp; &nbsp;
                                企业信息录入
                            </h3>
                            <hr>
                        </div>

                        <el-form-item label="昵称" prop="nikeName">
                            <el-col :span="10">
                                <el-input
                                    v-model="enterpriseInfo.nikeName"
                                    :disabled='!canUpdate'
                                    maxlength="50"
                                    show-word-limit></el-input>
                                <p class="notice">可以顺带改个昵称，最好与公司名差不多</p>
                            </el-col>
                            <!-- <el-input v-model="enterpriseInfo.nikeName"></el-input> -->
                        </el-form-item>

                        <el-form-item label="企业全称" prop="fullName">
                            <el-col :span="20">
                                <el-input v-model="enterpriseInfo.fullName"
                                          :disabled='!canUpdate'
                                          maxlength="50"
                                          show-word-limit></el-input>
                                <p class="notice">公司名称与营业执照一致</p>
                            </el-col>
                        </el-form-item>

                        <el-form-item label="统一社会信用代码" prop="socialCode">
                            <el-col :span="10">
                                <el-input v-model="enterpriseInfo.socialCode" :disabled='!canUpdate'></el-input>
                            </el-col>
                        </el-form-item>

                        <el-form-item label="注册资金" prop="registryMoney">
                            <el-col :span="10">
                                <el-input v-model="enterpriseInfo.registryMoney" :disabled='!canUpdate'>
                                    <template slot="append">万</template>
                                </el-input>
                            </el-col>
                        </el-form-item>

                         <el-form-item label="公司地址">
                            <el-col :span="10">
                                <el-select 
                                    v-model="enterpriseInfo.province" 
                                    @change="provChange"
                                    style="width: 100%;" 
                                    placeholder="所在省"
                                    :disabled='!canUpdate'>
                                    <el-option disabled value="">请选择</el-option>
                                    <el-option
                                      v-for="prov in allProvinces"
                                      :key="prov.name"
                                      :value="prov.name">
                                    </el-option>
                                </el-select>
                            </el-col>
                            <el-col class="line" :span="1" style="text-align: center;">-</el-col>
                            <el-col :span="10">
                                <el-select 
                                    v-model="enterpriseInfo.city" 
                                    placeholder="所在市" 
                                    style="width: 100%;"
                                    :disabled='!canUpdate'>
                                    <el-option disabled value="">请选择</el-option>
                                    <el-option
                                      v-for="city in allCities"
                                      :key="city.name"
                                      :value="city.name">
                                    </el-option>
                                </el-select>
                            </el-col>
                        </el-form-item>

                        <div style="padding-bottom: 30px;">
                            <h3 style="margin-bottom: 10px;">
                                <i class="el-icon-user-solid"></i>&nbsp; &nbsp;
                                运营信息录入
                            </h3>
                            <hr>
                        </div>

                        <el-form-item label="运营者姓名" prop="operatorName">
                            <el-col :span="10">
                                <el-input
                                    v-model="enterpriseInfo.operatorName"
                                    placeholder="输入姓名"
                                    :disabled='!canUpdate'
                                    maxlength="20"
                                    show-word-limit></el-input>
                                <p class="notice">

                                </p>
                            </el-col>
                        </el-form-item>

                        <el-form-item label="联系邮箱">
                            <el-col :span="10">
                                <el-input
                                    v-model="enterpriseInfo.operatorEmail"
                                    placeholder="输入邮箱"
                                    :disabled='!canUpdate'></el-input>
                            </el-col>
                        </el-form-item>

                        <el-form-item label="联系电话" prop="operatorPhone">
                            <el-col :span="20">
                                <el-input v-model="enterpriseInfo.operatorPhone" placeholder="输入手机号" style="width: 50%;" :disabled='!canUpdate'></el-input>
                                <p class="notice">
                                    手机号用于接收验证码，请使用有效手机号(当然目前来看随便填，毕竟现在没法发送验证码)
                                </p>
                            </el-col>
                        </el-form-item>

                        <el-form-item label="验证码" prop="captcha" v-if='canUpdate'>
                            <el-col :span="15">
                                <el-col :span="12">
                                    <el-input v-model="enterpriseInfo.captcha" placeholder="请输入验证码"></el-input>
                                    <p class="notice">验证码半小时内有效</p>
                                </el-col>
                                <el-col :span="8">
                                    <captcha
                                        @captcha='getCaptcha'
                                        @empty="captchaEmpty"
                                        :capTo='enterpriseInfo.operatorPhone'/>

                                </el-col>
                            </el-col>
                        </el-form-item>

                        <el-divider></el-divider>

                        <el-row style="display: flex; justify-content: flex-end;" v-if="canUpdate">
                            <el-button @click="contract = true">企业认证协议</el-button>
                            <el-button type="primary" @click="submitInfo('enterpriseInfo')">提交信息</el-button>
                        </el-row>

                    </el-form>
                </div>

            </div>
        <v-footer class="footer" />

        <el-dialog 
            title="协议"
            :visible.sync="contract" width='50%'>
            <div class="contract">
                <h3>企业认证协议</h3>
                <hr>
                <p>没见过这么敷衍的用户协议吧，没事，今天你就见到了~</p>
                <p>话说回来，真的有人会看这个东西么</p>
                <p>我奉劝您还是不要填真实信息，出事了概不负责</p>
                <p>说的您能填真实信息一样</p>
            </div>
            
        </el-dialog>

    </div>
</template>

<script>

import HeaderNav from "@/components/HeaderNav.vue";
import Footer from "@/components/Footer.vue";
import { getAllProvinces, getCitiesByProv } from "@/assets/js/city.js";
import CaptchaButton from "@/components/CaptchaButton.vue";
import {isEmpty, getUser, setEnterprise, setUser, setToken, getEnterprise} from "@/assets/js/utils.js";
import { OkMessage, ErrorMessage, OkNotify, ErrorNotify } from "@/assets/js/message.js";
import {Get, Post, Put} from "@/assets/js/http.js";
import {REG} from "@/assets/js/verify";

export default {
    name: "Enterprise",
    components: {
        "header-nav": HeaderNav,
        "v-footer": Footer,
        "captcha": CaptchaButton
    },

    created() {

        let user = getUser()

        let enterpriseInfo = getEnterprise()
        if(enterpriseInfo !== null){

            this.enterpriseInfo = enterpriseInfo
            this.enterpriseInfo.province = enterpriseInfo.enterAddress.split("-")[0]
            this.enterpriseInfo.city = enterpriseInfo.enterAddress.split("-")[1]

            if(enterpriseInfo.status === 1){
                this.finishStep = 1
                this.canUpdate = false
            } else if(enterpriseInfo.status === 2) {
                this.finishStep = 2
            } else if(enterpriseInfo.status === 3) {
                this.finishStep = 2
            }
        }else{
            this.enterpriseInfo.operatorPhone = user.phone
            this.enterpriseInfo.operatorEmail = user.email
        }

        this.enterpriseInfo.nikeName = user.nikeName


    },

    data() {
        return {
            enterpriseRules: {
                nikeName: [
                    { required: true, message: "信息不能为空", trigger: "blur" },
                ],
                fullName: [
                    { required: true, message: "信息不能为空", trigger: "blur" },
                ],
                socialCode: [
                    { required: true, message: "信息不能为空", trigger: "blur" },
                ],
                registryMoney: [
                    {
                        validator: (rule, value, callback) => {
                            if(!isEmpty(value) && !/^\d+\.?\d*$/.test(value)){
                                return callback(new Error("注册资金不符合规范"))
                            }

                            if(!isEmpty(value) && value.length > 10) {
                                return callback(new Error("您咋不把地球买上捏？？~"))
                            }

                            callback()

                        }, trigger: "blur"
                    }
                ],
                operatorName:[
                    { required: true, message: "信息不能为空", trigger: "blur" },
                ],
                operatorPhone: [
                    { validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("电话号不能为空"))
                        }

                        if(!REG.PHONE_REG.test(value)) {
                            return callback(new Error("电话号不符合规范"))
                        }

                        callback()
                        }, trigger: "blur" },
                ],
                operatorEmail: [
                    {
                        validator: (rule, value, callback) => {
                            if(!isEmpty(value) && !REG.EMAIL_REG.test(value)) {
                                return callback(new Error("邮箱格式错误"))
                            }

                            callback()
                        }, trigger: "blur" }
                ],
                captcha: [
                    { validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("验证码不能为空"))
                        }

                        if(!/^\d{4}$/.test(value)) {
                            return callback(new Error("验证码格式错误"))
                        }

                        callback()

                        }, trigger: "blur"
                    },
                ],

            },

            contract: false,

            enterpriseInfo:{
                nikeName: "",
                fullName: "",
                socialCode: "",
                registryMoney: "",
                province: "",
                city: "",
                operatorPhone: "",
                operatorEmail: "",
                operatorName: "",
                captcha: "",
            },

            sendCap: false,

            canUpdate: true,

            finishStep: 0,
        }
    },
    methods: {
        provChange(){
            console.log("prov change");
            this.enterpriseInfo.city = ""
        },

        getCaptcha(captcha){
            // OkMessage("发送成功")
            this.sendCap = true
            OkNotify("验证码", captcha)
        },

        captchaEmpty() {
            ErrorMessage("给谁发呀~啊？")
        },

        submitInfo(formName) {

            this.$refs[formName].validate((valid) => {
                if(valid) {

                    if(!this.sendCap){
                        ErrorNotify("嗯哼?", "不发验证码啦？？")
                        return
                    }

                    this.$confirm(
                        '您提交的信息在反馈认证结果之前将不能修改，是否确认提交信息',
                        '是否提交',
                        {
                            distinguishCancelAndClose: true,
                            confirmButton: "提交",
                            cancelButton: "取消"
                        }
                    ).then(() => {
                        console.log("enterprise info ok, can commit");

                        let enterpriseInfo = {
                            fullName: this.enterpriseInfo.fullName,
                            socialCode: this.enterpriseInfo.socialCode,
                            registryMoney: Number.parseInt(this.enterpriseInfo.registryMoney),
                            enterAddress: this.enterpriseInfo.province + "-" + this.enterpriseInfo.city,
                            operatorName: this.enterpriseInfo.operatorName,
                            operatorEmail: this.enterpriseInfo.operatorEmail,
                            operatorPhone: this.enterpriseInfo.operatorPhone
                        }

                        Put("/v1/enterprise").setJson(enterpriseInfo).send(
                            ok1 => {
                                if(ok1.code === 200) {
                                    Post("/v1/account/info").setForm({
                                        nikeName: this.enterpriseInfo.nikeName
                                    }).send(
                                        ok2 => {
                                            if(ok2.code === 200){
                                                this.canUpdate = false;
                                                this.finishStep = 1;
                                                OkNotify("成功", ok2.message)
                                                setEnterprise(enterpriseInfo)

                                                let user = getUser()
                                                user.nikeName = this.enterpriseInfo.nikeName
                                                setUser(user)
                                                setToken(ok2.data)
                                            }else{
                                                ErrorNotify("错误", ok2.message)
                                            }
                                        },

                                        err2 => {
                                            ErrorNotify("网络错误", err2)
                                        }
                                    )
                                }else{
                                    ErrorNotify("错误", ok1.message)
                                }
                            },
                            err => {
                                ErrorNotify("网络错误", err)
                            }
                        )

                    }).catch( _ => {

                    })
                }
            })
        }

    },
    computed: {
        allProvinces() {
            return getAllProvinces()
        },
        allCities(){
            return getCitiesByProv(this.enterpriseInfo.province)
        },
        captchaReady() {
            return !isEmpty(this.enterpriseInfo.operatorPhone)
        },
        finishStatus() {
            return "success"
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
    div.title {
        display: none;
    }

    div.moblie {
        display: block !important;
        text-align: center;
    }

    div.body {
        display: none;
    }

    .footer {
        position: absolute !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {

    div.title {
        display: none;
    }

    div.moblie {
        display: block !important;
        text-align: center;
    }

    div.body {
        display: none;
    }

    .footer {
        position: absolute !important;
    }
}

.footer{
    position: relative;
}

div.title{
    margin-top: 40px;
}

div.moblie {
    display: none;
}

div.title > h2 {
    text-align: center;
}

div.body {
    width: 50%;
    margin: 50px auto 0;
}

div.enterprise-info {
    margin-top: 50px;
    margin-bottom: 100px;
}

div.enterprise-info h3 {
    color: rgb(87, 87, 87);
    margin-top: 50px;
}

p.notice{
    font-size: 12px;
    color: grey;
}

div.contract {
    text-align: center;
}
div.contract > p {
    margin-top: 20px;
}

</style>