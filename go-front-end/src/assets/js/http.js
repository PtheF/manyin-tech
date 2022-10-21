import axios from "axios"
import { ErrorMessage } from "./message"
import env from "@/../public/env.json"
import {getUser, isEmpty, setToken, getToken, removeToken, removeUser} from "./utils"
import router from "../../router"

const server = axios.create({
    baseURL: env.api[env.model].server,
    timeout: 10000
})

server.interceptors.request.use(config => {


    try {
        let user = getUser()
        if (user) {
            let token = getToken()
            if (!isEmpty(token)){
                // console.log("set authorization");
                // console.log("request interceptor token: ", token);
                config.headers['Authorization'] = token
            }
        }

    } catch (error) {
        console.log("error: ", error);
    }

    return config

}, err => {
    console.log("err, ", err);
    Promise.reject(err)
})


// 这里出了一个小 bug，就是后面的 then(res => ) 怎么也拿不到res，结果发现是这里的过滤器
// 过滤完了以后没有把 res 返回，也就是直接把 res 拦截了。
// 这里有个注意事项，就是默认情况下，axios 只能获取默认的响应头，比如 Content-Type, Content-Length 这种
// 要获取我们自定义的响应头，首先需要服务端添加一个 Access-Control-Expose-Headers 这个头
// 用于指定暴露哪些我们自定义的响应头，然后才能获取。
// 再有就是，这里的响应头是小写，不是 Authorization，而是 authorization，即便服务端是大写的。
server.interceptors.response.use(
    res => {

        // console.log("response interceptor");

        // console.log(res);

        if (res.headers["authorization"] !== undefined && !isEmpty(res.headers["authorization"])) {
            console.log("set token");
            setToken(res.headers["authorization"])
        }

        if (res.data.code === 302){
            console.log("require api but login expire");
            removeToken()
            removeUser()
            router.replace({
                path: "/login?expire=true&uri=" + window.location.pathname
            })
        }

        return res
    },
    err => {
        ErrorMessage(err + "233")
        return err
    }
)

function newHttp() {
    let http = {

        url: "",
        method: "",
        json: {},
        form: {},
        header: {},

        isJson: false,
        isForm: false,

        formHeader: "application/x-www-form-urlencoded;charset=utf-8",
        jsonHeader: "application/json;charset=utf-8",

        setJson: function (jsonData) {
            if (this.isForm) {
                throw Error("不能同时传表单和json")
            }
            this.json = jsonData
            this.isJson = true
            return this
        },

        setForm: function (formData) {
            if (this.isJson) {
                throw Error("不能同时传表单和json")
            }

            this.isForm = true
            this.form = formData
            return this
        },

        setHeader: function (headers) {
            this.header = headers
            return this
        },



        send: function (ok, err) {
            let axiosConfig = {
                url: http.url,
                method: http.method
            }

            if (http.isForm) {
                http.header["Content-Type"] = http.formHeader
                axiosConfig.data = http.form
                axiosConfig.transformRequest = [function (data) {
                    // Do whatever you want to transform the data
                    let ret = ''
                    for (let it in data) {
                        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
                    }
                    return ret
                }]
            } else if (http.isJson) {
                http.header["Content-Type"] = http.jsonHeader
                axiosConfig.data = http.json
            }

            if (http.header !== {}) {
                axiosConfig.headers = http.header
            }

            // console.log("axios config:");
            // console.log(axiosConfig);

            server(axiosConfig).then(res => {
                // console.log("ok: ", res);
                ok(res.data)
            }).catch(e => err(e))
        }
    }

    return http
}

export const http = server

export const Get = function (url) {
    let http = newHttp()
    http.url = url
    http.method = "GET"

    return http
}

export const Post = function (url) {
    let http = newHttp()
    http.url = url
    http.method = "POST"

    return http
}

export const Put = function (url) {
    let http = newHttp()
    http.url = url
    http.method = "PUT"
    return http
}

export const Delete = function (url) {
    let http = newHttp()
    http.url = url
    http.method = "DELETE"
    return http
}

