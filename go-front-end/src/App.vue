<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
    export default {
        name: "App",

        async created() {
            await this.$store.dispatch("getProductTypes")
        }
    }
</script>

<style lang="less">

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  overflow: hidden;
  margin: 0;
  padding: 0;
}

</style>
