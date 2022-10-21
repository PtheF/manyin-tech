<template>
  <div class="app">
    <div class="note">
      <div class="title">
        <h1 style="color: #11999e">蔓茵科技</h1>
        <h2 style="color: #11999e">
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;打破传统回收之囹圄
        </h2>
      </div>

      <div class="content">
        <p>责任、绿色、协调、创新</p>
        <p>一款绿色高效的现代城市垃圾回收企业</p>
        <p>致力于城市建筑垃圾回收与再利用</p>

        <h4 style="color: #48466d; margin-top: 40px">
            打造中国建筑垃圾回收产业新世界。
        </h4>
      </div>
      <h3 style="margin-top: 50px; color: #48466d">
          蔓茵，攻克技术难关，拿下创新高地
      </h3>
    </div>

    <div class="login-body">
      <div class="login">
        <div style="margin: 40px 30px 40px 30px">
          <div style="margin: 20px 0 20px 0; border-bottom: 1px solid grey">
            <div style="color: grey; cursor: pointer;" @click="back">返回</div>
          </div>

          <h2 class="login-title">登录</h2>

          <div class="input">
            <input type="text" placeholder="手机号/电子邮箱" v-model="id" />
            <input type="password" placeholder="密码" v-model="pw" />

            <div style="padding-left: 100px; margin-top: 30px; color: grey">
              没有账号？<span class="reg"> <a href="/register">立即注册</a></span>
            </div>
          </div>

          <div class="login-but" @click="doLogin">登录</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getQueries } from "@/assets/js/uri";
import {InfoMessage, OkNotify, ErrorMessage, ErrorNotify} from "@/assets/js/message";
import { Verify, REG } from "@/assets/js/verify";
import { Post } from "@/assets/js/http";
import { setUser } from "@/assets/js/utils"
import env from "@/../public/env.json"
// import axios from "axios"

export default {
  name: "Login",
  data() {
    return {
      id: "",
      pw: "",
      url: "/",
    };
  },

  created() {

      if(env.typ === "dev"){
          this.id = "13411110987"
          this.pw = "hitman47"
      }

    let params = getQueries();

    if (params.expire === "true") {
      // InfoMessage("登录超时，请重新登录");
        ErrorNotify("过期", "登录过期，请重新登录")
    }

    if (params.uri !== undefined) {
      this.url = decodeURIComponent(params.uri);
    }

    let newUserId = sessionStorage.getItem("newUserId")
    let newUserPw = sessionStorage.getItem("newUserPw")

    if(newUserId != null && newUserPw != null) {
      this.id = newUserId
      this.pw = newUserPw

        sessionStorage.removeItem("newUserId")
        sessionStorage.removeItem("newUserPw")
    }
  },

  methods: {

    back() {
      window.history.back(-1)
    },

    doLogin() {
      let i = Verify(
        this.id,
        REG.PHONE_REG,
        REG.EMAIL_REG
      );

      if (!i) {
        ErrorMessage("手机号或邮箱不符合规则")
        return;
      }

      if (!Verify(this.pw, REG.PASS_REG)) {
        ErrorMessage("密码不符合规则")
        return;
      }

      Post("/v1/login").setForm({id: this.id, pw: this.pw, tp: i}).send(
      	res => {
      		if(res.code === 200){
      			// OkMessage("登陆成功")
                OkNotify("登录", "登陆成功")

                console.log("login user: ");
                console.log(res.data);

                setUser(res.data)

                this.$store.dispatch("fetchAddresses")

                setTimeout(() => {
                    this.$router.replace({

                        path: this.url

                    })
                }, 1000)
      		}else{
      			ErrorMessage(res.message)
      		}
      	},

      	err => {
      		console.log(err);
      	}
      )

    },
  },
};
</script>

<style scoped>
div.app {
  width: 100vw;
  height: 100vh;

  background-image: url("~@/../public/green.svg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 150px 0;

  position: relative;
  overflow: hidden;
  display: flex;
}

div.note {
  top: 10%;
  left: 10%;
  background: transparent;
  position: relative;
  height: auto;

  width: 24%;
}

div.content {
  margin-top: 70px;
  margin-left: 60px;
}

p {
  margin-top: 30px;

  color: rgb(160, 160, 160);
}

div.title {
  padding: 10px;
  /* background: linear-gradient(to right, #30e3ca, transparent) */
}

div.login-body {
  position: absolute;
  right: 15%;
  height: 100%;
  width: 26%;
  background-color: rgba(242, 255, 251, 0.7);
}

div.login {
  /* width: 90%; */
  height: 100%;
  /* margin-left: 20px; */
  /* border-right: 1px solid grey; */
  overflow: hidden;
}

@media all and (max-aspect-ratio: 5/3) and (min-device-aspect-ratio: 1/1) {
  div.login-body {
    right: 10%;
    width: 35%;
  }
}

@media all and (max-aspect-ratio: 4/3) and (min-device-aspect-ratio: 1/1) {
  div.login-body {
    right: 5%;
    width: 40%;
  }
  div.note {
    width: 40%;
  }
}

@media all and (max-aspect-ratio: 1/1) and (min-device-aspect-ratio: 1/1) {
  div.login-body {
    right: 0;
    width: 50%;
  }
  div.note {
    width: 50%;
  }
}

@media all and (max-aspect-ratio: 3/4) and (min-device-aspect-ratio: 1/1) {
  div.login-body {
    right: 0;
    width: 100%;
  }
  div.note {
    display: none;
  }

  div.app {
    background-position: -300px 0px;
  }
}

@media all and (max-device-aspect-ratio: 1/1) {
  div.login-body {
    right: 0;
    width: 100%;
  }
  div.note {
    display: none;
  }
  div.app {
    background-position: -300px 0px;
  }
}

h2.login-title {
  margin-top: 50px;
  padding: 10px;
  /* background: linear-gradient(to right, #30e3ca, transparent) */
}

div.input {
  margin-top: 30px;
}

input {
  outline: none;
  border: none;
  background: none;
  display: block;

  border-bottom: 2px solid #11999e;

  width: 250px;
  height: 50px;

  font-size: 18px;
  transition: 0.3s;
}

input:hover,
input:focus {
  width: 300px;
}

span.reg {
  transition: 0.3s;
  cursor: pointer;
  color: #48466d;
}
/* 
span.reg > a {
  color: #48466d;
} */

span.reg:hover {
  border-bottom: 1px solid #48466d;
}

div.login-but {
  margin-top: 40px;

  text-align: center;

  transition: 0.2s;

  border-bottom: 1px solid #11999e;

  cursor: pointer;
}

div.login-but:hover {
  background-color: #46cdcf;
  color: white;
}
</style>
