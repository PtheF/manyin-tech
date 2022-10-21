<template>
    <div>
        <el-button 
            type="primary" 
            @click="captcha" 
            :disabled = 'left !== sec'>{{captchaBut}}</el-button>
    </div>
</template>

<script>

import { Get } from "@/assets/js/http.js";
import { ErrorNotify } from "@/assets/js/message.js"
import { isEmpty } from '@/assets/js/utils.js'

export default {
    name: "CaptchaButton",
    props: {
        capTo: String,
    },
    data() {
        return {
            sec: 60,
            left: 60,
            interval: null
        }
    },
    methods: {

        captcha(){
            if(isEmpty(this.capTo)){
                this.$emit("empty")
                return
            }

            Get("/v1/captcha?id=" + this.capTo).send(
                ok => {
                    if(ok.code === 200){
                        setTimeout(() => {
                            this.interval = setInterval(() => {
                                if(this.left == 0){
                                    this.left = this.sec
                                    clearInterval(this.interval)
                                    return
                                }
                                this.left -= 1
                            }, 1000)
                        }, 0)                           
                        this.$emit("captcha", ok.data)

                    }else{
                        ErrorNotify("错误", ok.message)
                        
                    }
                },

                err => {
                    ErrorNotify("错误", err)

                }
            )
        }
    },
    computed: {
        captchaBut() {
            return this.left == this.sec ? "验证码" : this.left + "s"
        }
    }
}
</script>