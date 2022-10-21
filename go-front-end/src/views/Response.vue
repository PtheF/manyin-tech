<template>
    <div>
        <header-nav/>

        <div class="body">
            <h2 style="text-align: center; margin-top: 50px;">公益</h2>
            <p>
                建筑类垃圾资源化是顺应行动hi的市场性行为，但同时关乎到绝大多数百姓的生存状况与健康，所以亦可认为是一种公益性的事业，两者并不相悖。具体体现为：
            </p>

            <hr>

            <ol>
                <li>带来经济效益。</li>
                <li>减少建筑垃圾带来的危害。</li>
                <li>缓解建材紧缺尴尬局面。</li>
                <li>促进地方就业与带动建筑垃圾产业链发展。</li>
            </ol>

            <h2 style="text-align: center; margin-top: 30px;">反馈</h2>
            <br><br>
            <el-input
                type="textarea"
                :rows="3"
                placeholder="请输入您的反馈意见"
                v-model="useless">
            </el-input>

            <div style="display: flex; justify-content: center; margin-top: 20px;">
                <el-button type="primary" @click="submitResponse">提交</el-button>
            </div>

            <el-dialog
                title="反馈"
                :visible.sync="finSub">
                <span>感谢您的反馈，您的宝贵意见我们一定不会聆听，因为这个功能根本没做 :)</span>
                <span slot="footer" class="dialog-footer">
                    <el-button type="primary" @click="finSub = false">确 定</el-button>
                </span>
            </el-dialog>
        </div>

        <Footer style="position: relative;"/>
    </div>
</template>

<script>
import HeaderNav from "@/components/HeaderNav";
import Footer from "@/components/Footer"
export default {
    name: "Response",
    components: {HeaderNav, Footer},
    data() {
        return {
            useless: "",
            finSub: false
        }
    },
    methods: {
        submitResponse() {
            this.useless = ''

            this.finSub = true
        }
    }
}
</script>

<style scoped>

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 60% !important;
    }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 70% !important;
    }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 80% !important;
    }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }
}

@media all and (max-device-aspect-ratio: 1/1) {
    div.body {
        width: 95% !important;
    }
}

div.body {
    min-height: 90vh;
    width: 50%;
    margin: 0 auto;
    padding-bottom: 50px;
}

div.body > p {
    margin-top: 30px;
    text-indent: 20px;
    line-height: 25px;
    font-size: 20px;
    padding-bottom: 30px;
    color: #606266;
}

div.body > ol {
    margin-top: 20px;
    margin-left: 50px;

    color: #606266;
}

</style>