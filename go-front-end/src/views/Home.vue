<template>
    <div class="home">
<!--        <header-nav :bg-color="`transparent`" :text-color="`white`"/>-->

        <header-nav></header-nav>
<!--        /img/home_bg1.png-->
        <div class="bottom-img" style="background-image: url('/img/home_bg1.jpg');"></div>

        <div class="top" style="transform: translateY(-50vh);">
            <div class="logo">
                <div>
                    <img src="/logo.png" alt="" class="logo">
                    <h1 style="text-align: center; color: white; font-size: 50px;">蔓茵科技</h1>
                    <h3 style="text-align: center; color: #dcfff0; margin-top: 20px;">打破传统回收之囹圄</h3>
                </div>
            </div>
        </div>

        <div class="middle basic-info">
            <h1 align="center">蔓茵科技</h1>
<!--            <p style="margin-top: 50px; text-align: center;">蔓茵科技，只是某大学生的大学生创新创业项目，主要用于展示和宣传，没有其他任何其他的功能。</p>-->

<!--            <p style="margin-top: 50px; text-align: center;">本网站不包含任何收费业务，未来也不打算投入运营，只是为了应付一下大创答辩。</p>-->

<!--            <p style="margin-top: 50px; text-align: center;">这些文本内容可能会被修改。</p>-->

            <p>责任、绿色、协调、创新</p>

            <p>作为一款绿色高效的现代城市垃圾回收企业，主要用于城市建筑垃圾回收与再利用，致力于打造中国建筑垃圾回收产业新世界。</p>

            <p>蔓茵科技致力于攻克技术难关，拿下创新高地。</p>
        </div>

        <div class="bottom-img" style="background-image: url('/img/home_bg6.jpg')"></div>

        <div class="top" style="transform: translateY(-30vh);">
            <h1 style="text-align: center; color: white; transform: scale(1.5);">蔓茵回收</h1>
        </div>

        <div class="middle basic-info">
            <h1 align="center">蔓茵回收业务</h1>

<!--            <p style="margin-top: 50px; text-align: center;">回收似乎是这个项目的主要业务，卖东西倒是次要的，但是具体怎么回收，为啥回收，咱也不知道，毕竟我只是个开发。</p>-->

<!--            <p style="margin-top: 50px; text-align: center;">但是这里还是得占位，后期这里可能有大段文本，这里提前看看放大段文本是个啥效果，后面再改。</p>-->

<!--            <p style="margin-top: 50px; text-align: center;">-->
<!--                很好，没得写了，不如说说他们这个大创，这个大创是个省级大创，咱也不懂什么大创才能使国家级，可能是研究核反应堆吧-->
<!--            </p>-->

<!--            <p style="text-align: center; width: 60%; margin: 50px auto 0; line-height: 50px;">-->
<!--                再说经费，这个经费有点少啊，才2000块钱，-->
<!--                关键我一分钱没有，就是为爱发电，我都感觉我在做慈善，买服务器最简单的乞丐版轻量级应用服务器，一年就得100多，如果这个项目真能保证经费，-->
<!--                直接上腾讯云服务器，直接买它100台组成集群，一年恐怕就得10个万了，然后域名买它10年的，DNS最贵的套餐买它10年的，在套上腾讯的服务器防护套餐，-->
<!--                那这个经费不得50万？-->
<!--            </p>-->

            <p>本企业通过建筑垃圾产量预测、设备研发、平台开发、基于用户需求设计产品等模块实现盈利。</p>

            <p>客户与蔓茵共同进步：主动了解、迅速满足客户需要，为客户交付高质量的产品和服务。将以客户为中心的理念贯穿于全业务流程。</p>

            <p>自然与蔓茵共同发展：通过创新优化选择对环境最有利的措施进行建筑垃圾回收，积极推行清洁生产、节约生产、绿色生产，努力实现蔓茵与自然环境的友好发展。</p>

            <p>蔓茵注重质量，质量成就蔓茵：蔓茵以高度的主人翁责任感，认真负责生产再生品，严格遵守质量规范，追求以高质量、高技术的工艺流程和再生产品来回馈顾客。</p>

        </div>

        <div class="bottom-img" style="background-image: url('/img/home_bg3.jpg')"></div>

        <div class="middle">
            <h1>技术实现</h1>
            <div class="tech-ach">
                <div>
                    <h3>前端</h3>
                    <p>
                        本网站是基于 Vue.js 实现的单页面项目，部分还使用了 Element-UI 库用于导入现成的组件，统一网站风格.
                        同时使用 axios 用于发送请求。
                    </p>
                    <br><br>
                    <h3>后端</h3>
                    <p>
                        后端使用近几年较为热门的 go 语言开发，使用 beego 框架搭建后端服务。使用 MySQL 做数据库，
                        同时使用 Redis 实现缓存，ID生成，登录令牌 等功能，登录业务使用了 JWT 令牌配合 Redis 实现
                        前后端令牌校验刷新超时自动登出功能，提供系统安全性。
                    </p>
                    <br><br>
                    <h3>部署</h3>
                    <p>
                        整个项目部署在公网服务器，使用 Nginx 部署前端项目，同时为后端服务进行反向代理。
                    </p>
                </div>

                <div style="background-image: url('/img/tech-ach.png');
                background-repeat: no-repeat;
                background-size: 100% auto;">
<!--                    <img src="/img/tech-ach.png" alt="" width="100%" height="auto">-->
                </div>
            </div>
        </div>

        <Footer class='home-footer'/>


    </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Footer from "@/components/Footer.vue"
import HeaderNav from '../components/HeaderNav.vue'
import VMask from "@/components/VMask";

export default {
    name: 'Home',
    components: {
        'header-nav': HeaderNav,
        'Footer': Footer,
        'v-mask': VMask
    },
    data(){
        return {
            drawer: false
        }
    },
    methods: {
        showMenu() {
            this.drawer = true
        }
    }
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
    div.tech-ach {
        display: block !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.tech-ach {
        display: block !important;
    }

    div.tech-ach > div:nth-child(1) {
        width: 100% !important;
    }

    div.tech-ach > div:nth-child(2) {
        width: 100% !important;
        height: 200px !important;
    }
}

	.home-footer {
		position: relative;
	}

    div.home {
        height: 100vh;
        overflow-x: hidden;
        perspective: 3px;
        color: #606266;
    }

    div.home::-webkit-scrollbar {
        background-color: #e1e1e1;
        /*background-color: transparent;*/
        width: 8px;
        box-shadow: 1px 2px 5px #424242 inset;
    }

    div.home::-webkit-scrollbar-thumb {
        width: 8px;
        background-color: #818181;
        box-shadow: 1px 2px 4px #b9b9b9 inset;
    }

    div.bottom-img {
        transform: translateZ(-1px) scale(1.5);
        background-size: cover;
        background-repeat: no-repeat;
        position: relative;
        /*z-index: -1;*/
        height: 100vh;
        z-index: -1;
    }

    div.middle {
        background-color: white;
        padding: 50px 40px;
        /*height: 100vh;*/
        min-height: 50vh;
    }

    div.top {
        transform: translateZ(1px);
        z-index: -1;
    }

    div.logo {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    img.logo {
        width: 200px;
        height: 200px;
    }

    div.tech-ach {
        margin-top: 50px;
        display: flex;
        flex-wrap: wrap;
    }

    /* 描述文本 div */
    div.tech-ach > div:nth-child(1) {
        width: 40%;
    }

    /* 描述文本 div > p */
    div.tech-ach > div:nth-child(1) > p {
        line-height: 40px;
        text-indent: 20px;
        font-size: 15px;
    }

    /* 图片 div */
    div.tech-ach > div:nth-child(2) {
        width: 60%;
        display: flex;
        flex-wrap: wrap;
        height: 400px;
    }

    div.basic-info {

    }

    div.basic-info > p {
        margin-top: 50px;
        text-align: center;
    }

</style>