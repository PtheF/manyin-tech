<template>
    <div class="prod-tps-body">
        <header-nav></header-nav>
        <div class="product-types">
            <div class="logo">

                <h2>蔓茵商城</h2>
                <p style="margin-top: 20px; line-height: 20px;">{{shopSubtext}}</p>
                <div></div>
            </div>

            <div class="prod-types">
                <div class="prod-type"
                     v-for="(prodType, index) in prodTypes" :key="index"
                     :style="`background-color: ${typesColor[index]}; width: ${prodTypeWidth}`"
                     @click="toProdList(prodType.prodTypeId)">
                    <div>
                        <h3>{{ prodType.tName }}</h3>
                        <p>{{ prodType.subText }}</p>
                    </div>
                </div>
            </div>
        </div>
        <Footer style="position: relative;"></Footer>
    </div>
</template>

<script>
import HeaderNav from "@/components/HeaderNav";
import Footer from "@/components/Footer";
export default {
    name: "ProductTypes",
    components: {
        HeaderNav, Footer
    },

    data() {
        return {
            shopSubtext: `
                这里就是蔓茵商城，各类建筑垃圾回收产品的聚集地，有装配式建筑、再生材料、再生骨料、再生工艺品的四大类再生产品供您挑选，
                这里同样可以将您的意见与建议及时反馈给我们，为您服务是蔓茵商城的无上荣耀。
            `,
            typesColor: [
                "rgba(193,255,255,0.5)",
                "rgba(209,255,227,0.5)",
                "rgba(255,211,211,0.5)",
                "rgba(255,211,238,0.5)"
            ]
        }
    },

    methods: {
        toProdList(tid){
            console.log(tid)
            this.$router.push({
                path: "/shop/list/" + tid,
            })
        }
    },

    computed: {
        prodTypes() {
            return this.$store.getters.getAllProductTypes
        },
        prodTypeWidth() {
            return `${100 / this.prodTypes.length}%`
        }
    }


}
</script>

<style scoped>

@media all and (min-aspect-ratio: 5/3) {
    div.prod-tps-body::-webkit-scrollbar {
        background-color: #e1e1e1;
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.prod-tps-body::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }
}

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.prod-tps-body::-webkit-scrollbar {
        background-color: #e1e1e1;
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.prod-tps-body::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.prod-tps-body::-webkit-scrollbar {
        background-color: #e1e1e1;
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.prod-tps-body::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.prod-tps-body::-webkit-scrollbar {
        background-color: #e1e1e1;
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.prod-tps-body::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.prod-tps-body {
        height: auto !important;
    }

    div.product-types {
        height: auto !important;
        overflow: auto !important;
        background-position: center 0;
    }
    div.prod-types{
        display: block !important;
    }

    div.prod-type {
        width: 100% !important;
        height: auto !important;
        padding-top: 40px;
        padding-bottom: 40px;
    }

    div.prod-type > div > h3 {
        transform: translateY(0) !important;
    }

    div.prod-type > div > p {
        color: #494c4f !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.prod-tps-body {
        height: auto !important;
    }

    div.product-types {
        height: auto !important;
        overflow: auto !important;
        background-position: center 0;
    }
    div.prod-types{
        display: block !important;
    }

    div.prod-type {
        width: 100% !important;
        height: auto !important;
        padding-top: 40px;
        padding-bottom: 40px;
    }

    div.prod-type > div > h3 {
        transform: translateY(0) !important;
    }

    div.prod-type > div > p {
        color: #494c4f !important;
    }
}

div.prod-tps-body {
    height: 100vh;
    overflow-y: scroll;
}

div.product-types {
    background-size: cover;
    background-repeat: no-repeat;
    background-image: url("/img/home_bg6.jpg");
    /*background: black;*/
    overflow: hidden;
    height: 100vh;
}

div.prod-types {
    display: flex;
    background: none;
    justify-content: space-around;
    height: 90%;
    overflow: hidden;
}

div.prod-type{
    /*width: 25%;*/
    /*background: rgba(200, 200, 200, 0.5);*/
    transition: .3s;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
}

div.prod-type:hover {
    background-color: rgba(250, 250, 250, 0.6) !important;
    /*box-shadow: 20px 20px 50px #dcdcdc inset;*/
}

div.prod-type > div {
    transition: .3s;
    /*color: #606266;*/

}

div.prod-type > div > h3{
    text-align: center;
    transition: .3s;
    transform: translateY(50px);
}

div.prod-type > div > p {
    margin-top: 20px;
    padding: 0 40px 20px 40px;
    color: transparent;
    /*color: white;*/
    text-align: center;
    transition: .3s;
}

div.prod-type:hover > div > h3 {
    transform: translateY(0);
    color: #3b8585;
}

div.prod-type:hover > div {
    transform: scale(1.1);
}

div.prod-type:hover > div > p {
    color: #494c4f;
}

div.logo {
    /*height: 10%;*/
    background: linear-gradient(to bottom, white, rgba(255, 255, 255, 0.7));
    box-shadow: 0 5px 10px black;
    padding: 30px;
}

div.logo > h2 {
    text-align: center;
    /*color: #6e6e6e;*/
    color: #3f9161;
}

div.logo > p {
    color: #606266;
    text-align: center;
    font-size: 13px;
}


</style>