<template>
    <div>
        <el-menu mode="horizontal" :background-color="bgColor" :text-color="textColor">
            <el-menu-item class="logo">
                <router-link to="/">
                    <div>
                        <img :src="logoPng" width="40px;" height="40px;">
                        <span>&nbsp;&nbsp;&nbsp;蔓茵科技</span>
                    </div>
                </router-link>
            </el-menu-item>
            <el-submenu index="1" class="pc-nav">
                <template slot="title">商城</template>
                <el-menu-item index="1-0">
                    <router-link to="/shop/types">商品类别</router-link>
                </el-menu-item>
                <el-menu-item v-for="(item, index) in productTypes" :key="index">
                    <router-link :to="`/shop/list/${item.prodTypeId}`">
                        {{ item.tName }}
                    </router-link>
                </el-menu-item>
            </el-submenu>
            <el-menu-item class="pc-nav" @click="recovery">
                <router-link to="#">回收</router-link>
            </el-menu-item>
            <el-menu-item index="2" class="pc-nav">
                <router-link to="/response"> 公益与反馈 </router-link>
            </el-menu-item>

			<el-menu-item index="5" v-if="!isLogin">
                <router-link to="/login"> 登录 </router-link>
            </el-menu-item>

			<el-menu-item index="6" v-if="!isLogin">
                <router-link to="/register">注册</router-link>
            </el-menu-item>

			<el-submenu index="7" v-if="isLogin">
                <template slot="title">个人</template>
                <el-menu-item index="7-1">
                    <router-link to="/space" style='color: #606266'> 个人中心 </router-link>
                </el-menu-item>
                <el-menu-item index="7-3">
                    <router-link to="/order" style='color: #606266'>订单</router-link>
                </el-menu-item>
                <el-menu-item index='7-4' style="color: red;" @click='logout'>
                    登出
                </el-menu-item>
            </el-submenu>


        </el-menu>
    </div>
</template>

<script>

import { isLogin, removeUser, removeToken } from "@/assets/js/utils"
import { OkNotify, ErrorNotify } from "@/assets/js/message"
import logoPng from "../../public/logo.png"

export default {

    props: {
        bgColor: String,
        textColor: String,
    },

	data(){
		return {
			isLogin: isLogin(),
            logoPng: logoPng,
		}
	},

    // async beforeMount(){
    //     let prodTypes = this.$store.getters.getAllProductTypes
    //     console.log("header nav get type by getters")
    //     console.log(prodTypes)
    //
    //     console.log("header nav get type by state")
    //     console.log(this.$store.state.ProductTypes.productTypes)
    //     console.log("this")
    //     console.log(this)
    //     if(prodTypes.length === 0){
    //         console.log("all product types is null")
    //         await this.$store.dispatch("getProductTypes")
    //     }
    // },

    computed: {
        productTypes(){
            return this.$store.getters.getAllProductTypes
        }
    },
    methods: {
        logout() {
            removeToken()
            removeUser()
            OkNotify("账户", "已登出")
            this.isLogin = false
            this.$router.replace({
                path: "/"
            })
        },

        recovery(){
            ErrorNotify(":(", "项目负责人执意不做这个功能，呜呜 我哭了~")
        }
    }
};
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {

}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {

}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {

}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    .logo {
        display: none;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    .logo {
        /*width: 400px !important;*/
        /*!*display: none !important;*!*/
        /*margin: 0 auto;*/
        display: none;
    }

    .logo-mob {
        display: block !important;
    }

    /*.pc-nav {*/
    /*    display: none;*/
    /*}*/

    /*.logo > div {*/
    /*    !*margin-right: 50px !important;*!*/
    /*    margin-left: 10px !important;*/
    /*}*/

    .nav-mob {
        display: block !important;
    }
}

a {
    color: #95979e;
    transition: .2s;
}

a:hover {
    color: #606266;
}

.logo {
    width: 250px;
    /*margin-left: 0px;*/
}

.logo-mob {
    display: none;
}

.logo > a > div {
    margin-left: 50px;
    font-size: 20px;
}

.nav-mob {
    display: none;
}


</style>