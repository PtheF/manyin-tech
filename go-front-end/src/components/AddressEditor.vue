<template>
    <v-mask :show="show" @close="cancel">
        <template>
            <div class="add-address">
                <h2 style="margin-bottom: 20px;">
                    <i class="el-icon-edit"/>
                    新增地址
                </h2>

                <hr>

                <div>
                    <el-form
                        :model="newAddress"
                        ref="newAddress"
                        label-position="top"
                        label-width="80px"
                        :rules="addressRules">
                        <el-form-item label="地址名称" prop="name">
                            <el-input
                                v-model="newAddress.name"
                                maxlength="20"
                                prefix-icon="el-icon-info"
                                show-word-limit></el-input>
                        </el-form-item>
                        <el-form-item label="收件人" prop="addressee">
                            <el-input
                                v-model="newAddress.addressee"
                                maxlength="30"
                                prefix-icon="el-icon-user-solid"
                                show-word-limit></el-input>
                        </el-form-item>
                        <el-form-item label="地址" prop="address">
                            <el-input
                                v-model="newAddress.address"
                                type="textarea"
                                maxlength="255"
                                prefix-icon="el-icon-s-home"
                                show-word-limit></el-input>
                        </el-form-item>
                        <el-form-item label="联系电话" prop="phone">
                            <el-input
                                v-model="newAddress.phone"
                                prefix-icon="el-icon-phone"></el-input>
                        </el-form-item>

                        <el-form-item>
                            <div style="display: flex; justify-content: flex-end;">
                                <el-button type="danger" icon="el-icon-error" @click="cancel">取消</el-button>
                                <el-button type="primary" icon="el-icon-upload" @click="addAddress('newAddress')">添加</el-button>
                            </div>
                        </el-form-item>
                    </el-form>
                </div>

            </div>
        </template>
    </v-mask>
</template>

<script>
import VMask from "@/components/VMask";
import {isEmpty} from "@/assets/js/utils";
import {REG} from "@/assets/js/verify";

export default {
    name: "AddressEditor",
    props: {
        show: Boolean
    },
    components: {
        "v-mask": VMask
    },

    data() {
        return {
            newAddress: {
                name: "",
                addressee: "",
                address: "",
                phone: ""
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
                    { validator: (rule, value, callback) => {
                        if(isEmpty(value)){
                            return callback(new Error("电话号不能为空"))
                        }

                        if(!REG.PHONE_REG.test(value)){
                            return callback(new Error("电话号不符合规则"))
                        }

                        callback()
                        }, trigger: "blur" }
                ]
            },
        }
    },
    methods: {

        cancel() {
            this.$emit("cancel")
        },

        addAddress(formName) {

            this.$refs[formName].validate((valid) => {
                if (valid) {

                    let emitAddress = {
                        name: this.newAddress.name,
                        addressee: this.newAddress.addressee,
                        address: this.newAddress.address,
                        phone: this.newAddress.phone
                    }

                    this.$emit("add", emitAddress)

                    this.newAddress = {
                        name: '',
                        addressee: '',
                        address: '',
                        phone: ''
                    }
                }
            });
        },
    }
}
</script>

<style scoped>


@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.add-address {
        width: 60% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.add-address {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.add-address {
        width: 80% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.add-address {
        width: 95% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.add-address {
        width: 95% !important;
    }
}

div.add-address {
    width: 60%;
    margin: 20px auto 0;
}
</style>