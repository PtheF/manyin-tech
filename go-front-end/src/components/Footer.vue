<template>
    <footer>
        <div>联系我们</div>
        <div>
            <div>
                <a href="/about">
                    <i class='el-icon-user'/>
                    &nbsp;&nbsp;关于我们&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                </a>
            </div>
            <div>
                <i class="el-icon-message"/>
                &nbsp;&nbsp;15831362294@163.com&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            </div>
            <div>
                <i class="el-icon-location"/>
                &nbsp;&nbsp;河北省石家庄市新华区河北经贸大学
            </div>
        </div>
        <div>
            <div>Powered by
                <a href="https://cn.vuejs.org">Vue.js</a> &
                <a href="https://element.eleme.cn/#/zh-CN">Element</a> &
                <a href="https://github.com/beego/beego">beego</a>
            </div>
        </div>

        <div>
<!--            <a href="https://beian.miit.gov.cn/" target="_blank">30166479591226768</a> |-->
            <a href="https://beian.miit.gov.cn/" target="_blank">备案号：冀ICP备2022023531号-1</a>
        </div>
    </footer>
</template>

<script>
export default {
    name: "Footer"
}
</script>

<style scoped>


    @media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    
    }

    @media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
        footer > div:nth-child(2) {
            display: block !important;
            justify-content: center;
        }

        footer > div:nth-child(2) > div {
            margin-left: 20px;
        }
    }

    @media all and (max-device-aspect-ratio: 1/1) {
        footer > div:nth-child(2) {
            display: block !important;
            justify-content: center;
        }

        footer > div:nth-child(2) > div {
            margin-left: 20px;
        }
    }

    footer{
        /* background-color: #233A55; */
        /* color: #DFEEFF; */
        background-color: rgb(238, 238, 238);
        /* background: linear-gradient(to bottom, white, rgb(187, 187, 187)); */
        /* border-top: 1px solid grey; */
        position: absolute;
        box-shadow: 0 3px 5px rgb(197, 197, 197) inset;
        bottom: 0;
        left: 0;
        right: 0;
        color: #606266;
        padding-bottom: 20px;
    }

    footer > div {
        padding-top: 2px;
        font-size: 13px;
        line-height: 13px;
    }

    footer > div:nth-child(1) {
        text-align: center;
        padding-top: 20px;
    }

    footer > div:nth-child(2) {
        display: flex;
        justify-content: center;

    }

    footer > div:nth-child(3) {
        display: flex;
        justify-content: center;
        /*padding-bottom: 20px;*/
    }

    footer > div:nth-child(3) > div > a {
        color: black;
        text-decoration: underline;
    }

    footer > div:nth-child(2) > div {
        /* width: 20%; */
        height: 30px;
        line-height: 30px;
         /* color: #606266; */
    }

    footer > div:nth-child(2) > div:nth-child(1) {
       cursor: pointer;
    }

    footer > div:nth-child(2) > div:nth-child(1) > a {
       color: #606266;
    }

    footer > div:nth-child(2) > div:nth-child(1) > a:hover {
       color: rgb(47, 99, 167);
    }

    footer > div:nth-child(4) {
        /*color: #606266;*/
        margin-top: 15px;
        text-align: center;
    }

    footer > div:nth-child(4) > a {
        color: #606266;
    }
</style>
