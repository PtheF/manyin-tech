<template>
    <div class="mask" v-show="vShow">
        <div class="content">
            <!-- <div style="height: 1000px; background-color: red;"></div> -->
            <div style="width: 90%; margin: 20px auto 10px; cursor: pointer; border-bottom: 1px solid #a4a4a4;" @click="back">
                <span style="color: #a4a4a4; font-size: 15px;">返回</span>
            </div>
            <slot></slot>
        </div>
    </div>

</template>

<script>
export default {
    name: "VMask",

    props: {
        show: Boolean
    },

    data() {
        return {
            vShow: false,
        }
    },

    beforeDestroy() {
        // console.log("v-mask before destroy")
        document.querySelector("body").style.overflow = ""
    },

    methods: {
        back() {
            this.vShow = false
            this.$emit("close")

            document.querySelector("body").style.overflow = ""
        }
    },

    watch: {
        'show': function(newValue, oldValue) {
            // console.log(`show new value: ${newValue}, old value: ${oldValue}`)
            this.vShow = newValue

            if(this.vShow){
                console.log("v-mask open")
                let body = document.querySelector("body");
                body.style.overflow = 'hidden'
            }else{
                console.log("v-mask close")
                document.querySelector("body").style.overflow = ""
            }
        }
    }
}
</script>

<style scoped>

    @media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
        div.content {
            width: 50% !important;
        }
    }

    @media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
        div.content {
            margin: 0 auto 0 !important;
            width: 65% !important;
            height: 100% !important;
        }
    }

    @media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
        div.content {
            width: 80% !important;
        }
    }

    @media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
        div.content {
            margin: 40px auto 0 !important;
            width: 100% !important;
            height: 100% !important;
        }
    }

    @media all and (max-device-aspect-ratio: 1/1) {
        div.content {
            margin: 40px auto 0 !important;
            width: 100% !important;
            height: 100% !important;
        }
    }

    div.mask{
        z-index: 100;
        background-color: rgba(0, 0, 0, 0.5);
        width: 100vw;
        height: 100vh;
        position: fixed;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
    }

    div.content{
        background-color: white;
        width: 40%;
        margin: 40px auto 0;
        height: 90%;
        overflow-y: scroll;
        /*padding: 20px;*/
    }

    div.content::-webkit-scrollbar {
        background-color: #e1e1e1;
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.content::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }

</style>