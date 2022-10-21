<template>
    <div class="body">
        <el-menu class="el-menu-demo" mode="horizontal">
            <el-menu-item index="1">
                <router-link to="/ceo/newProduct">
                    新增商品
                </router-link>
            </el-menu-item>
            <el-submenu index="2">
                <template slot="title">商品列表</template>
                <el-menu-item index="2-1" v-for="(item, index) in productTypes" :key="index">
                    <router-link :to="`/ceo/products/${item.prodTypeId}/1`">
                        {{item.tName}}
                    </router-link>
                </el-menu-item>
            </el-submenu>
            <el-menu-item index="3">
                <router-link to="/ceo/productTypes">
                    商品类型
                </router-link>
            </el-menu-item>
        </el-menu>

        <div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
export default {
    name: "Admin",

    computed: {
        productTypes(){
            return this.$store.getters.getAllProductTypes
        }
    }
}
</script>

<style scoped>

div.body {
    padding: 30px 10px 50px 10px;
}

a {
    color: #95979e;
    transition: .2s;
}

a:hover {
    color: #606266;
}

</style>