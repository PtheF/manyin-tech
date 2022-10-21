<template>

    <div class="app">

        <header-nav />

        <div class="body">
            <div class="navi">
                <div style="padding-top: 30px">
                    <!-- <h3></h3> -->
                    <a href="#private">
                        <div>
                            <i class="el-icon-user-solid" />&nbsp;&nbsp;账号信息
                        </div>
                    </a>
                    <a href="#secure">
                        <div>
                            <i class="el-icon-key" />&nbsp;&nbsp;账号安全
                        </div>
                    </a>
                    <a href="#address">
                        <div>
                            <i class="el-icon-location" />&nbsp;&nbsp;地址管理
                        </div>
                    </a>
                    <a href="#enterprise">
                        <div>
                            <i class="el-icon-s-promotion" />&nbsp;&nbsp;企业认证
                        </div>
                    </a>
                    <router-link to="/order">
                        <div><i class="el-icon-s-order" />&nbsp;&nbsp;订单</div>
                    </router-link>
                </div>
            </div>

            <div class="info">
                <div>
                    <h2>
                        <a name="private"></a>
                        <i class="el-icon-user-solid" />
                        &nbsp; &nbsp;账号信息
                    </h2>
                    <div class="account">
                        <div class="avatar-line">
                            <div class="avatar"></div>
                        </div>

                        <el-descriptions direction="vertical" :column='2' border>
                            <el-descriptions-item label="昵称">{{user.nikeName}}</el-descriptions-item>
                            <el-descriptions-item label="手机号">{{user.phone}}</el-descriptions-item>
                            <el-descriptions-item label="电子邮箱">{{user.email}}</el-descriptions-item>
                            <el-descriptions-item label="注册时间">{{user.regDate}}</el-descriptions-item>
                        </el-descriptions>

                        <el-row style='margin-top: 30px;'>
                            <div style='display:flex; justify-content: center; margin: 0 auto;'>
                                <el-button
                                    icon="el-icon-edit"
                                    type="primary"
                                    @click='editAccountInfo'>编辑</el-button>
                            </div>
                        </el-row>


                    </div>
                </div>

                <div>
                    <h2>
                        <a name='secure'></a>
                        <i class='el-icon-key'></i>&nbsp;&nbsp;账号安全
                    </h2>
                    <div class="secure-tabs">
                        <el-tabs v-model="secureTab">
                            <el-tab-pane label="密保手机" name="phone"></el-tab-pane>
                            <el-tab-pane label="密保邮箱" name="email" v-if="user.email !== ''"></el-tab-pane>
                            <el-tab-pane label="密码修改" name="password"></el-tab-pane>
                        </el-tabs>
                    </div>

                    <div class="secure">

                        <div class="_secure" v-show='secureTab === "phone"'>
                            <div v-show='!updatePhoneVerify'>
                                <el-form
                                    label-position="top"
                                    :model='verifyForm'
                                    ref='verifyForm'
                                    :rules='verifyFormRules'>
                                    <el-form-item label='注册手机号:'>
                                        <el-col :span='22'>
                                            <el-input v-model='user.phone' disabled></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='能否收到短信' prop='verifyCaptcha'>
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='verifyForm.verifyCaptcha' ></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button
                                                    @captcha='getCaptcha'
                                                    :capTo='user.phone'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='收不到短信' prop='verifyPw'>
                                        <el-col :span='22'>
                                            <el-input type='password' v-model='verifyForm.verifyPw' show-password placeholder="输入密码"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="">
                                        <el-col :span='22'>
                                            <div style='display: flex; justify-content: flex-end;'>
                                                <el-button type='primary' @click='verifyUser("verifyForm")'>验证</el-button>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-form>
                            </div>

                            <div v-show='updatePhoneVerify'>
                                <span><i class='el-icon-info'></i>&nbsp;验证通过,请在10分钟内完成换绑</span>
                                <el-form
                                    label-position="top"
                                    :model='rebindForm'
                                    ref="rebindForm"
                                    :rules='rebindFormRules'>
                                    <el-form-item label='换绑的手机号' prop='rebindPhone'>
                                        <el-col :span='22'>
                                            <el-input v-model='rebindForm.rebindPhone' placeholder="输入手机号"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='发送验证码' prop='rebindCaptcha'>
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='rebindForm.rebindCaptcha' placeholder="输入验证码"></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button
                                                    @captcha='getCaptcha'
                                                    :capTo='rebindForm.rebindPhone'
                                                    @empty='capEmpty'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="">
                                        <el-col :span='22'>
                                            <div style='display: flex; justify-content: flex-end;'>
                                                <el-button type='primary' @click='updatePhone("rebindForm")'>更新</el-button>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-form>
                            </div>

                        </div>

                        <div v-show="secureTab === 'email'">
                            <div v-show='!updateEmailVerify'>
                                <el-form
                                    label-position="top"
                                    :model='verifyForm'
                                    ref='verifyForm'
                                    :rules='verifyFormRules'>
                                    <el-form-item label='注册邮箱:'>
                                        <el-col :span='22'>
                                            <el-input v-model='user.email' disabled></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='能否收到邮件' prop="verifyCaptcha">
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='verifyForm.verifyCaptcha'></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button @captcha='getCaptcha' :capTo='user.email'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='收不到邮件' prop='verifyPw'>
                                        <el-col :span='22'>
                                            <el-input type='password' v-model='verifyForm.verifyPw' show-password placeholder="输入密码"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="">
                                        <el-col :span='22'>
                                            <div style='display: flex; justify-content: flex-end;'>
                                                <el-button type='primary' @click='verifyUser("verifyForm")'>验证</el-button>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-form>
                            </div>

                            <div v-show='updateEmailVerify'>
                                <span><i class='el-icon-info'></i>&nbsp;验证通过,请在10分钟内完成换绑</span>
                                <el-form
                                    label-position="top"
                                    :model='rebindForm'
                                    ref='rebindForm'
                                    :rules='rebindFormRules'>
                                    <el-form-item label='换绑的邮箱' prop='rebindEmail'>
                                        <el-col :span='22'>
                                            <el-input v-model='rebindForm.rebindEmail' placeholder="输入新邮箱"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='发送验证码' prop='rebindCaptcha'>
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='rebindForm.rebindCaptcha' placeholder="输入验证码"></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button
                                                    @captcha='getCaptcha'
                                                    :capTo='rebindForm.rebindEmail'
                                                    @empty='capEmpty'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="">
                                        <el-col :span='22'>
                                            <div style='display: flex; justify-content: flex-end;'>
                                                <el-button type='primary' @click='updateEmail("rebindForm")'>更新</el-button>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-form>
                            </div>
                        </div>

                        <div v-show="secureTab === 'password'">
                            <div v-show='!updatePasswordVerify'>
                                <el-form
                                    label-position="top"
                                    :model="verifyForm"
                                    ref="verifyForm"
                                    :rules="verifyFormRules">
                                    <el-form-item label='手机能否收到短信' prop="verifyCaptcha">
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='verifyForm.verifyCaptcha'></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button @captcha='getCaptcha' :capTo='user.phone'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='邮箱能否收到邮件' prop="verifyCaptcha">
                                        <el-col :span='22'>
                                            <el-col :span='13'>
                                                <el-input v-model='verifyForm.verifyCaptcha'></el-input>
                                            </el-col>
                                            <el-col :span='1' style='text-align: center'>-</el-col>
                                            <el-col :span='7'>
                                                <captcha-button @captcha='getCaptcha' :capTo='user.email'/>
                                            </el-col>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label='是否记得旧密码' prop="verifyPw">
                                        <el-col :span='22'>
                                            <el-input type='password' v-model='verifyForm.verifyPw' show-password placeholder="输入密码"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="">
                                        <el-col :span='22'>
                                            <div style='display: flex; justify-content: flex-end;'>
                                                <el-button type='primary' @click='verifyUser("verifyForm")'>验证</el-button>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-form>
                            </div>

                            <div v-show="updatePasswordVerify">
                                <span><i class='el-icon-info'></i>&nbsp;验证通过,请在10分钟内完成修改</span>

                                <el-form
                                    label-position="top"
                                    :model="resetPwForm"
                                    ref="resetPwForm"
                                    :rules="resetPwFormRules">

                                    <el-form-item label="输入新密码" prop="pw">
                                        <el-col :span="22">
                                            <el-input v-model="resetPwForm.pw" type="password" show-password placeholder="请输入新密码"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item label="请重复密码" prop="repPw">
                                        <el-col :span="22">
                                            <el-input v-model="resetPwForm.repPw" type="password" show-password placeholder="请重复密码"></el-input>
                                        </el-col>
                                    </el-form-item>

                                    <el-button type="primary" @click="updatePassword('resetPwForm')">修改</el-button>

                                </el-form>

                            </div>
                        </div>
                    </div>
                </div>

                <div>
                    <h2>
                        <a name="address" />
                        <i class="el-icon-location" />&nbsp;&nbsp;收货地址
                    </h2>
                    <div class="address">

                        <div>
                            <div @click="isAddAddress = true">添加地址</div>
                            <span style="color: grey; font-size: 16px; margin-left: 30px;">(最多十个地址)</span>
                        </div>

                        <div class="address-list" id="address-list">
<!--                            <span color="grey" v-if="addresses.length === 0" >未添加地址</span>-->
                            <el-empty description="暂无地址" v-if="addresses.length === 0"></el-empty>

                            <div class="address-info" v-for="(item, index) in addresses" :key="index">
                                <div class="address-info-row" v-if='item.isDefault'>
                                    <strong>默认地址</strong>
                                </div>
                                <div class="address-info-row">
                                    <div>名称:</div>
                                    <p>{{item.name}}</p>
                                </div>

                                <div class="address-info-row">
                                    <div style="line-height: 26px; height: 26px;">收件人:</div>
                                    <p style="line-height: 26px; height: 26px;">{{item.addressee}}</p>
                                </div>

                                <div class="address-info-row">
                                    <div>地址:</div>
                                    <p>{{item.address}}</p>
                                </div>

                                <div class="address-info-row">
                                    <div style="line-height: 26px; height: 26px;">联系电话:</div>
                                    <p style="line-height: 26px; height: 26px;">{{item.phone}}</p>
                                </div>

                                <div class="address-func" v-if='!item.isDefault'>
                                    <div @click="deleteAddress(index)">删除</div>
                                    <div @click='setDefaultAddress(index)'>默认</div>
                                    <div style="color: transparent;">占位</div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>

                <div>
                    <h2>
                        <a name="enterprise" />
                        <i class="el-icon-s-promotion" />&nbsp; &nbsp;企业认证
                    </h2>
                    <div class="enterprise">

                        <div v-if="enterpriseInfo === null">
                            <h4 style="text-align: center; color: grey;">未进行企业认证</h4>
                            <div style="text-align: center; padding-top: 30px;">
                                <el-button type="primary" @click="goConfirm">点击进行认证</el-button>
                            </div>

                        </div>

                        <div class="enterprise-info" v-if="enterpriseInfo !== null">

                            <el-descriptions :column="1" border>
                                <el-descriptions-item label="状态">
                                    <i class="el-icon-info" style="color: #1989fa;" v-if="enterpriseInfo.status === 1"></i>
                                    <i class="el-icon-success" style="color: #67C23A;" v-if="enterpriseInfo.status === 2"></i>
                                    <i class="el-icon-warning" style="color: #F56C6C;" v-if="enterpriseInfo.status === 3"></i>
                                    {{enterpriseStatus}}
                                </el-descriptions-item>
                                <el-descriptions-item label="企业全名">{{enterpriseInfo.fullName}}</el-descriptions-item>
                                <el-descriptions-item label="社保号">{{enterpriseInfo.socialCode}}</el-descriptions-item>
                                <el-descriptions-item label="运营人姓名">{{enterpriseInfo.operatorName}}</el-descriptions-item>
                                <el-descriptions-item label="运营人手机号">{{enterpriseInfo.operatorPhone}}</el-descriptions-item>
                                <el-descriptions-item label="运营人邮箱">{{enterpriseInfo.operatorEmail}}</el-descriptions-item>
                            </el-descriptions>
                            <div style="text-align: center; padding-top: 30px;">
                                <el-button type="primary" @click="goConfirm">查看详细信息</el-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <v-mask :show='isEditAccount' @close="leaveEditAccount">
            <div class='edit-account'>

                <h4 style='margin-bottom: 20px;'><i class='el-icon-edit'></i>&nbsp; &nbsp;编辑账户信息</h4>
                <hr>
                <el-form
                    :model='updateAccountForm'
                    ref='updateAccountForm'
                    label-position="left"
                    label-width="50px"
                    :rules="updateAccountFormRules"
                    style='margin-top: 40px;'>
                    <el-form-item label='昵称' prop="nikeName">
                        <el-col :span='20'>
                            <el-input v-model="updateAccountForm.nikeName"></el-input>
                        </el-col>

                    </el-form-item>

                    <el-form-item label='电话' prop="nikeName">
                        <el-col :span='20' v-if='user.phone === ""'>
                            <el-input v-model='updateAccountForm.phone' placeholder="添加密保手机号"></el-input>
                        </el-col>
                        <el-row v-if='user.phone !== ""'>
                            <el-col :span='14'>
                                <el-input v-model='user.phone' :disabled='true'></el-input>
                            </el-col>
                            <el-col :span='8'>
                                <el-button @click='updateBind'>换绑</el-button>
                            </el-col>
                        </el-row>

                    </el-form-item>

                    <el-form-item label='邮箱' prop="email">
                        <el-col :span='20' v-if='user.email === ""'>
                            <el-input v-model='updateAccountForm.email' placeholder="添加绑定邮箱"></el-input>
                        </el-col>

                        <el-row v-if='user.email !== ""'>
                            <el-col :span='14'>
                                <el-input v-model='user.email' :disabled='true'></el-input>
                            </el-col>
                            <el-col :span='8'>
                                <el-button @click='updateBind'>换绑</el-button>
                            </el-col>
                        </el-row>
                    </el-form-item>

                    <el-form-item label=''>
                        <el-col :span='20'>
                            <div  style='display:flex; justify-content: flex-end'>
                                <el-button @click='leaveEditAccount'>取消</el-button>
                                <el-button type='primary' @click='updateAccount("updateAccountForm")'>保存</el-button>
                            </div>
                        </el-col>
                    </el-form-item>
                </el-form>
            </div>
        </v-mask>

        <address-editor :show="isAddAddress" @add="addAddress" @cancel="isAddAddress=false"></address-editor>

        <Footer class="space-footer"/>

    </div>

</template>

<script>
import { Get, Delete, Put, Post } from "@/assets/js/http";
import { OkMessage, ErrorMessage, OkNotify, ErrorNotify } from "@/assets/js/message";
import {getUser, removeArray, isEmpty, setEnterprise} from "@/assets/js/utils";
import HeaderNav from "@/components/HeaderNav.vue";
import Footer from "@/components/Footer.vue";
import VMask from "@/components/VMask.vue";
import CaptchaButton from '../components/CaptchaButton.vue';
import { REG } from "@/assets/js/verify"
import { setToken, setUser } from "@/assets/js/utils"
import AddressEditor from "@/components/AddressEditor";

export default {
    name: "Space",
    components: {
        HeaderNav, Footer, "v-mask": VMask,
        CaptchaButton,
        AddressEditor
    },
    data() {

        return {
            user: getUser(),
            addresses: [],
            showTextArea: false,
            isAddAddress: false,
            isEditAccount: false,

            enterpriseInfo: null,

            updateAccountForm: {
                nikeName: "",
                phone: "",
                email: ""
            },

            updateAccountFormRules: {
                nikeName: [ {
                    validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("昵称不能为空"))
                        }
                        if(value.length >= 50){
                            return callback(new Error("昵称过长"))
                        }

                        callback()

                    }, trigger: "blur" } ],

                phone: [ {
                    validator: (rule, value, callback) => {
                        if(!isEmpty(value) && !REG.PHONE_REG.test(value)){
                            return callback(new Error("电话格式错误"))
                        }
                        callback()
                    }, trigger: "blur"} ],

                email: [ {
                    validator: (rule, value, callback) => {
                        if(!isEmpty(value) && !REG.EMAIL_REG.test(value)) {
                            return callback(new Error("邮件格式错误"))
                        }
                        callback()
                    }, trigger: "blur"} ]
            },

            newAddress: {
                id: "",
                name: "",
                addressee: "",
                address: "",
                phone: "",
                isDefault: false
            },
            addressRules: {
                name: [
                    { required: true, message: "请输入地址名称", trigger: 'blur'}
                ],
                addressee: [
                    { required: true, message: "请输入收件人", trigger: "blur" }
                ],
                address: [
                    { required: true, message: "请输入地址", trigger: "blur"}
                ],
                phone: [
                    { required: true, message: "请输入联系电话", trigger: "blur" }
                ]
            },

            updatePhoneVerify: false,
            updateEmailVerify: false,
            updatePasswordVerify: false,

            secureTab: "phone",

            verifyForm: {
                verifyCaptcha: "",
                verifyPw: ""
            },

            verifyFormRules: {
                verifyCaptcha: [ {
                    validator: (rule, value, callback) => {
                        console.log("verifyCaptchaValidator work")
                        if(this.verifyForm.verifyPw === "" && value === ""){
                            return callback(new Error("验证码或密码不能为空"))
                        }

                        if(!isEmpty(value) && !/\d{4}/.test(value)){
                            return callback(new Error("验证码格式错误"))
                        }

                        callback()
                    }, trigger: "blur" } ],
                verifyPw: [ {
                    validator: (rule, value, callback) => {
                        if(this.verifyForm.verifyCaptcha === "" && value === ""){
                            return callback(new Error("验证码或密码不能为空"))
                        }

                        if(!isEmpty(value) && !REG.PASS_REG.test(value)){
                            return callback(new Error("密码不符合规范"))
                        }

                        callback()
                    }, trigger: "blur" } ]
            },

            rebindForm: {
                rebindPhone: "",
                rebindEmail: "",
                rebindCaptcha: ""
            },

            rebindFormRules: {
                rebindPhone: [ {
                    validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("换绑手机不能为空"))
                        }
                        if(value === this.user.phone){
                            return callback(new Error("您换了个鬼！"))
                        }
                        if(!REG.PHONE_REG.test(value)){
                            return callback(new Error("手机号不符合规范"))
                        }
                        callback()
                    }, trigger: "blur" } ],

                rebindEmail: [ {
                    validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("换绑邮箱不能为空"))
                        }
                        if(value === this.user.email) {
                            return callback(new Error("您换了个鬼！"))
                        }
                        if(!REG.EMAIL_REG.test(value)){
                            return callback(new Error("邮箱不符合规范"))
                        }
                        callback()
                    }, trigger: "blur" } ],

                rebindCaptcha: [ {
                    validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("验证码不能为空"))
                        }
                        if(!/^\d{4}$/.test(value)){
                            return callback(new Error("验证码不符合规范"))
                        }
                        callback()
                    }, trigger: "blur" } ]
            },

            resetPwFormRules: {
                pw: [ { validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("密码不能为空"))
                        }
                        if(!REG.PASS_REG.test(value)){
                            return callback(new Error("密码不符合规则"))
                        }
                        callback()
                    }, trigger: "blur" } ],
                repPw: [ { validator: (rule, value, callback) => {
                        if(!isEmpty(this.resetPwForm.pw) && isEmpty(value)){
                            return callback(new Error("请重复密码"))
                        }
                        if(!isEmpty(this.resetPwForm.pw) && this.resetPwForm.pw !== value){
                            return callback(new Error("密码不一致"))
                        }
                        callback()
                    }, trigger: "blur" } ]
            },

            resetPwForm: {
                pw: "",
                repPw: ""
            }
        };
    },

    created() {

        console.log("current user:")
        console.log(getUser())

        Get("/v1/space").send(
            (res) => {
                // console.log("space res: ");
                //
                // console.log(res);
                // if(res.data.addresses != null){
                //     this.addresses = res.data.addresses;
                // }

                if(!isEmpty(res.data.enterprise.userId)){
                    console.log("enterprise not empty")
                    this.enterpriseInfo = res.data.enterprise
                    setEnterprise(res.data.enterprise)
                }

            },
            (err) => {
                console.log(err);
            }
        );

        this.$store.dispatch("fetchAddresses")

        this.addresses = this.$store.getters.getAddresses

        this.updateAccountForm = {
            nikeName: this.user.nikeName,
            phone: this.user.phone,
            email: this.user.email
        }
    },

    methods: {

        getCaptcha(cap) {
            OkNotify("验证码", cap)
        },

        editAccountInfo() {
            // this.updateUser.newNikeName = this.user.nikeName
            this.isEditAccount = true
        },

        updateAccount(formName){
            this.$refs[formName].validate( validate => {
                if(validate) {
                    Post("/v1/account/info").setForm({
                        nikeName: this.updateAccountForm.nikeName,
                        phone: this.updateAccountForm.phone,
                        email: this.updateAccountForm.email
                    }).send(
                        ok => {
                            if(ok.code === 200){
                                OkNotify("成功", "修改信息成功")
                                setToken(ok.data)
                                this.user.nikeName = this.updateAccountForm.nikeName
                                this.user.phone = this.updateAccountForm.phone
                                this.user.email = this.updateAccountForm.email
                                setUser(this.user)
                                this.isEditAccount = false
                            }else{
                                ErrorNotify("错误", ok.message)
                            }
                        },

                        err => {
                            ErrorNotify("错误", err)
                        }
                    )
                }
            } )

        },

        verifyUser(formName) {

            this.$refs[formName].validate((validate) => {
                if(validate){

                    console.log("validate ok");
                    Post("/v1/account/verify").setForm({
                        cp: this.verifyForm.verifyCaptcha,
                        pw: this.verifyForm.verifyPw,
                        tp: this.verifyForm.verifyPw === "" ? "cp" : "pw"
                    }).send(
                        ok => {
                            if(ok.code === 200){
                                OkNotify("认证", ok.message)
                                this.verifyForm.verifyCaptcha = ""
                                this.verifyForm.verifyPw = ""

                                setTimeout(() => {
                                    switch(this.secureTab){
                                    case "phone":
                                        this.updatePhoneVerify = true
                                        break
                                    case "email":
                                        this.updateEmailVerify = true
                                        break
                                    case "password":
                                        this.updatePasswordVerify = true
                                        break
                                    default:
                                        break
                                    }
                                }, 200)
                                
                            }else{
                                ErrorNotify(ok.message)
                            }
                        },

                        err => {
                            ErrorNotify("网络错误", err)
                        }
                    )
                }

            })
        },

        clearRebindForm() {
            this.rebindForm = {
                rebindPhone: "",
                rebindEmail: "",
                rebindCaptcha: ""
            }
        },

        updatePhone(formName) {

            this.rebindForm.rebindEmail = "hold@test.com"

            this.$refs[formName].validate((validate) => {
                if(validate) {

                    console.log("validate rebind form ok")

                    Post("/v1/account/phone").setForm({
                        phone: this.rebindForm.rebindPhone,
                        cap: this.rebindForm.rebindCaptcha
                    }).send(
                        ok => {
                            if(ok.code === 200){
                                OkNotify("换绑成功")
                                setToken(ok.data)
                                this.user.phone = this.rebindForm.rebindPhone
                                setUser(this.user)
                                this.updatePhoneVerify = false
                                this.clearRebindForm()
                            }else{
                                ErrorNotify("错误", ok.message)
                                this.updatePhoneVerify = false
                                this.clearRebindForm()
                            }
                        },

                        err => {
                            ErrorNotify("错误", err)
                            this.updatePhoneVerify = false
                            this.clearRebindForm()
                        }
                    )
                }
            })

            
        },

        updateEmail(formName) {

            this.rebindForm.rebindPhone = "18230110889"

            this.$refs[formName].validate((validate) => {
                if(validate){
                    Post("/v1/account/email").setForm({
                        cap: this.rebindForm.rebindCaptcha,
                        email: this.rebindForm.rebindEmail
                    }).send(
                        ok => {
                            if(ok.code === 200) {
                                OkNotify("成功", "修改邮箱成功")
                                setToken(ok.data)
                                this.user.email = this.rebindForm.rebindEmail
                                setUser(this.user)
                                this.updateEmailVerify = false
                                this.clearRebindForm()
                            }else{
                                ErrorNotify("错误", ok.message)
                                this.updateEmailVerify = false
                                this.clearRebindForm()
                            }
                        },

                        err => {
                            ErrorNotify("错误", err)
                            this.updateEmailVerify = false
                            this.clearRebindForm()
                        }
                    )
                }
            })
        },

        updatePassword(formName) {
            this.$refs[formName].validate((validate) => {
                if(validate){
                    Post("/v1/account/pw").setForm({
                        pw: this.resetPwForm.pw
                    }).send(
                        ok => {
                            if(ok.code === 200){
                                OkNotify("成功", "修改密码成功")
                                this.updatePasswordVerify = false
                                this.resetPwForm.pw = ""
                                this.resetPwForm.repPw = ""
                            }else{
                                ErrorNotify("错误", ok.message)
                                this.updatePasswordVerify = false
                                this.resetPwForm.pw = ""
                                this.resetPwForm.repPw = ""
                            }
                        },

                        err => {
                            ErrorNotify("错误", err)
                            this.updatePasswordVerify = false
                            this.resetPwForm.pw = ""
                            this.resetPwForm.repPw = ""
                        }
                    )
                }
            })

        },

        leaveEditAccount() {
            this.isEditAccount = false
        },

        goConfirm(){
            this.$router.push({
                path: "/enterprise"
            })
        },

        updateBind() {
            this.isEditAccount = false;
            window.location.href = '/space#secure'
        },

        deleteAddress(index) {
            let addressId = this.addresses[index].id;
            Delete("/v1/address/" + addressId).send(
                (ok) => {
                    if (ok.code === 200) {
                        OkMessage("删除成功");

                        // for(let i = index; i < this.addresses.length-1; i++){
                        //     this.addresses[i] = this.addresses[i+1]
                        // }

                        // this.addresses.pop()

                        // removeArray(this.addresses, index)
                        this.$store.commit("removeAddress", index)
                        this.addresses = this.$store.getters.getAddresses

                    } else {
                        ErrorMessage(ok.message);
                    }
                },

                (err) => {
                    ErrorMessage("删除失败" + err);
                }
            );
        },
        setDefaultAddress(index) {
            let addressId = this.addresses[index].id;
            Post("/v1/address/" + addressId).send(
                (ok) => {
                    if (ok.code === 200) {
                        OkMessage(ok.message);
                        // for(let i = 0; i < this.addresses.length; i++){
                        //     if(this.addresses[i].isDefault){
                        //         this.addresses[i].isDefault = false
                        //         break
                        //     }
                        // }
                        //
                        // this.addresses[index].isDefault = true

                        this.$store.commit("setDefaultAddress", index)
                        this.addresses = this.$store.getters.getAddresses

                    } else {
                        ErrorMessage(ok.message);
                    }
                },
                (err) => {
                    ErrorMessage(err);
                }
            );
        },

        // addAddress(formName) {
        //
        //     console.log("add address: ");
        //     console.log(this.newAddress);
        //
        //     this.$refs[formName].validate((valid) => {
        //         if (valid) {
        //             Put("/v1/address")
        //                 .setJson(this.newAddress)
        //                 .send(
        //                     (ok) => {
        //                         if (ok.code === 200) {
        //                             OkMessage(ok.message);
        //
        //                             if(this.addresses.length === 0){
        //                                 this.newAddress.isDefault = true
        //                             }
        //                             this.newAddress.id = ok.data
        //                             this.addresses.push(this.newAddress)
        //                             this.newAddress = {
        //                                 name: "",
        //                                 addressee: "",
        //                                 phone: "",
        //                                 address: ""
        //                             }
        //
        //                             this.isAddAddress = false;
        //                         } else {
        //                             ErrorMessage(ok.message);
        //                         }
        //                     },
        //                     (err) => {}
        //                 );
        //         } else {
        //             console.log('error submit!!');
        //             return false;
        //         }
        //     });
        // },

        addAddress(address) {
            Put("/v1/address").setJson(address).send(
                ok => {
                    if (ok.code === 200) {
                        OkMessage(ok.message);
                        if(this.addresses.length === 0){
                            address.isDefault = true
                        }
                        address.id = ok.data
                        // this.addresses.push(address)
                        // this.isAddAddress = false;

                        this.$store.commit("addAddresses", address)
                        this.addresses = this.$store.getters.getAddresses
                        this.isAddAddress = false
                    } else {
                        ErrorMessage(ok.message);
                    }
                },

                err => {
                    ErrorNotify("错误", "网络异常")
                }
            )
        },

        capEmpty() {
            ErrorNotify("错误", "我给谁发验证码？？")
        },
    },

    computed: {
        enterpriseStatus() {
            if(this.enterpriseInfo.status === 1) {
                return `企业信息正在审核`
            }

            if(this.enterpriseInfo.status === 2) {
                return `已通过审核`
            }

            if(this.enterpriseInfo.status === 3) {
                return `信息审核不通过，请修改信息重新提交`
            }
        }
    }
};
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 85% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 90% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 100% !important;
    }

    div.secure {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.navi{
        display: none;
    }

    div.app {
        height: 100vh !important;
        overflow-y: scroll;
    }

    div.body {
        width: 100% !important;
        margin-bottom: 0 !important;
        height: auto !important;
    }

    .space-footer {
        position: relative;
    }

    div.info {
        overflow-y: hidden !important;
        -ms-overflow-y: hidden !important;
    }

    div.info-row {
        width: 95% !important;
    }

    div.info-row > div:nth-child(1){
        padding-left: 10px !important;
        border: none !important;
    }

    div.info-row > div:nth-child(2) {
        border: none !important;
    }

    /* div.info > div > h2 {
        padding-left: 10px !important;
    } */

    div.account, div.address, div.secure, div.enterprise {
        width: 95% !important;
        margin: 30px auto 0 !important;
    }

    div.edit-account {
        width: 90% !important;
        margin: 0 auto;
    }

    div.address-list {
        padding-left: 10px !important;
    }

    div.address-info {
        width: 88% !important;
        border: none !important;
        border-bottom: 1px solid grey !important;
        box-shadow: none !important;

    }

    div.address-info-row {
        font-size: 14px !important;
    }

    div.add-address {
        width: 90% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.navi{
        display: none;
    }

    div.app {
        height: 100vh !important;
        overflow-y: scroll;
    }

    div.body {
        width: 100% !important;
        margin-bottom: 0 !important;
        height: auto !important;
    }

    .space-footer {
        position: relative;
    }

    div.info {
        overflow-y: hidden !important;
        -ms-overflow-y: hidden !important;
    }

    div.info-row {
        width: 95% !important;
    }

    div.info-row > div:nth-child(1){
        padding-left: 10px !important;
        border: none !important;
    }

    div.info-row > div:nth-child(2) {
        border: none !important;
    }

    /* div.info > div > h2 {
        padding-left: 10px !important;
    } */

    div.account, div.address, div.secure, div.enterprise {
        width: 95% !important;
        margin: 30px auto 0 !important;
    }

    div.edit-account {
        width: 90% !important;
        margin: 0 auto;
    }

    div.address-list {
        padding-left: 10px !important;
    }

    div.address-info {
        width: 88% !important;
        border: none !important;
        border-bottom: 1px solid grey !important;
        box-shadow: none !important;

    }

    div.address-info-row {
        font-size: 14px !important;
    }

    div.add-address {
        width: 90% !important;
    }
}

input {
    outline: none;
    border: none;
    display: block;
    background: none;
}

div.app {
    position: relative;
    width: 100%;
    height: 100vh;
    /* overflow: hidden; */
    z-index: 0;
}

div.body {
    height: 75%;
    width: 80%;
    margin: 20px auto 30px;
    /*background: #48466d;*/

    display: flex;
    box-shadow: 3px 3px 5px grey;
    z-index: 0;
}

/* ====================== 导航 ========================= */
div.navi {
    width: 20%;
    height: 100%;
    background-color: #326d87;
}

div.navi > div h3 {
    padding-left: 20px;
    color: white;
    margin-bottom: 40px;
    padding-bottom: 20px;
    border-bottom: 2px solid white;
}

div.navi a {
    outline: none;
    text-decoration: none;
    color: black;
}

div.navi > div div {
    padding-left: 20px;
    padding-top: 20px;
    padding-bottom: 20px;
    transition: 0.2s;
    cursor: pointer;
    font-size: 15px;
    color: white;
    border-left: 4px solid transparent;
}

div.navi > div div:hover {
    border-bottom: 1px solid white;
    padding-left: 30px;
    background: white;
    color: #326d87;
    border-left: 4px solid #e34343;
}

/* ====================== 信息栏 ======================== */

div.info {
    flex-grow: 1;
    /*background-color: #11999e;*/
    overflow-y: scroll;
    -ms-overflow-y: scroll;
}

div.info > div {
    height: auto;
    width: 100%;
    position: relative;
    padding-bottom: 40px;
    transition: 0.2s;
    /* margin-bottom: 10px; */
    /* box-shadow: 2px 2px 5px grey; */
    border-bottom: 1px solid grey;
}

div.info > div > h2 {
    padding-left: 30px;
    padding-top: 30px;
    color: rgb(70, 70, 70);
}

/* =================== 账号管理 =================== */
div.account{
    margin: 20px auto;
    height: auto;
    width: 60%;
    /*background-color: antiquewhite;*/
    color: #606266;
}

div.avatar-line {
    height: 200px;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 30px;
}

div.avatar {
    width: 180px;
    height: 180px;
    border-radius: 90px;
    border: 2px solid #34aad7;
}

div.edit-account{
    width: 60%;
    margin: 30px auto 0;
}

/*=====================账号安全======================= */

div.secure-tabs{
    margin: 20px auto;
    height: auto;
    width: 90%;
    color: #606266;
}

div.secure {
    margin: 0 auto;
    height: auto;
    width: 50%;
    color: #606266;
}

/*=====================地址管理=======================*/

div.address {
    margin: 20px auto;
    height: auto;
    width: 90%;
    /*background-color: antiquewhite;*/
    color: #606266;
}

div.address > h4 {
    width: 94%;
    margin: 20px auto 0;
}

div.address > div:nth-child(1){
    padding-left: 30px;
    margin-top: 20px;
    display: flex;
}

div.address > div:nth-child(1) > div {
    width: 100px;
    text-align: center;
    cursor: pointer;
    border-bottom: 2px solid #298ab4;
    transition: .2s;
}

div.address > div:nth-child(1) > div:hover{
    background-color: #298ab4;
    color: white;
    box-shadow: 2px 2px 4px grey;
}

div.address-list {
    width: 100%;
    margin-top: 20px;
    height: auto;

    padding-left: 30px;

    /* border-left: 1px solid #d14358; */
}

div.address-info {
    width: 90%;
    /* border: 1px solid grey; */
    padding: 10px;
    margin-top: 10px;
    box-shadow: 2px 2px 7px rgb(182, 182, 182);
    border-left: 3px solid #326d87;
}

div.address-info-row {
    display: flex;
    height: auto;
    margin-top: 5px;
    font-size: 16px;
}

div.address-info-row > div {
    width: 100px;
}

div.address-info-row > p {
    color: rgb(90, 90, 90);
}

div.address-func{
    display: flex;
    position: relative;
    height: auto;
}

div.address-func > div:nth-child(1){
    position: absolute;
    right: 30px;
    width: 70px;
    color: rgb(131, 28, 28);
    cursor: pointer;
}

div.address-func > div:nth-child(2){
    position: absolute;
    right: 120px;
    width: 70px;
    color: rgb(42, 78, 124);
    cursor: pointer;
}

div.add-address {
    width: 60%;
    margin: 20px auto 0;
}

/*===================企业认证====================*/

div.enterprise {
    margin: 40px auto 0;
    height: auto;
    width: 60%;
}

</style>
