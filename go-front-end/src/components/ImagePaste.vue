<template>
    <div>
        <el-input
            @paste.native.capture.prevent="pasteImage"
            style="width: 150px;"
            v-model="imageName"
            @input="removeImg"
            placeholder="粘贴图片"
            clearable></el-input>
        <div class="show-image">
            <div class="paste"></div>
        </div>
    </div>
</template>

<script>

import {ErrorNotify} from "@/assets/js/message";
import login from "@/views/Login";

export default {
    name: "ImagePaste",

    data(){
        return {
            imageData: "",
            imageName: "",
            imgNode: null,

            base64Prefix: "data:image/png;base64,"
        }
    },

    methods: {

        removeImg(){
            console.log("change")
            if(this.imageName === '' && this.imgNode !== null){
                console.log("remove image")
                // document.getElementById("paste").removeChild(this.imgNode)
                this.$el.getElementsByClassName("paste")[0].removeChild(this.imgNode)
            }
        },

        pasteImage(e){

            console.log("paste image")

            if(this.imageName !== '') {
                ErrorNotify("错误", "不能重复粘贴")
                return
            }

            let file = null;
            const items = (e.clipboardData || window.clipboardData).items;

            if(items && items.length) {
                for(let i = 0; i < items.length; i++){
                    console.log(`items[${i}]`)
                    console.log(items[i])
                    if(items[i].type.indexOf("image") !== -1){
                        file = items[i].getAsFile()
                    }
                }
            }

            console.log("file:")
            console.log(file)

            if(file) {

                let th = this

                let reader = new FileReader()

                reader.onload = function(event) {
                    let img = document.createElement("img");

                    // th.imageData = event.target.result.substring(th.base64Prefix.length, event.target.result.length - th.base64Prefix.length)

                    th.imageData = event.target.result
                    img.src = th.imageData

                    // img.src = th.base64Prefix + th.imageData

                    // img.width = 100;
                    // img.height = 100;

                    img.style.width = 'auto'
                    img.style.height = '100%'
                    // img.style.padding = '10px'
                    // img.style.border = '1px solid #69c7ea'
                    th.$el.getElementsByClassName("paste")[0].appendChild(img)

                    th.imgNode = img

                    th.$emit("getImg", th.imageData)
                    th.imageName = "粘贴的图片"

                }

                reader.readAsDataURL(file)

            } else {
                ErrorNotify("错误", "粘贴的不是图片")
            }
        }
    }
}
</script>

<style scoped>

div.paste {
    /*width: 280px;*/
    max-width: 90%;
    /*height: 300px;*/
    /*max-height: 200px;*/
    overflow: hidden;
    /*border: 1px solid grey;*/

    padding-left: 20px;
    border-left: 4px solid #69c7ea;
    /*border-radius: 20px;*/
}

div.show-image {
    margin-top: 20px;
    /*padding: 10px;*/
    /*border: 1px solid #69c7ea;*/
    /*border-radius: 20px;*/

    display: flex;
}

</style>