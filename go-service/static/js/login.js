var loginButton = "login-but"
var id = "id"
var pw = "pw"

var loginButtonEl = document.getElementById(loginButton)

let params = getQueries()

if(params.expire === "true"){
    ShowError("登录超时，请重新登录")
}

let u = params.uri

if(u === undefined){
    u = "/"
}

console.log("uri=" + u)

loginButtonEl.addEventListener("click", () => {

    let i = ValidateForm(id, "手机号或邮箱不符合规范", /^1[3|4|5|8][0-9]\d{8}$/, /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)

    if(i === 0) {
        return
    }

    if(!ValidateForm(pw, "密码不符合规则", /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$/)){
        return
    }

    let display = loginButtonEl.style.display
    loginButtonEl.style.display = "none"

    var loginData = {
        id: document.getElementById(id).value,
        pw: document.getElementById(pw).value,
        tp: i
    }

    Post("/api/login").setForm(loginData).refreshToken("Authorization", (jwt) => {
        sessionStorage.setItem("token", jwt)
    }).send(
        (res) => {

            if(res.code === 200){
                // console.log("getResponseHeader: ")
                // console.log(xhr.getAllResponseHeaders())

                ShowInfo("登录成功")
                console.log("user")
                console.log(res.data)

                sessionStorage.setItem("user", res.data)
                //
                setTimeout(() => {
                    location.href = u
                }, 1000)
            }else{
                ShowError(res.data.message)
                loginButtonEl.style.display = display
            }

        },

        err => {
            ShowError(err)
            loginButtonEl.style.display = display
        }
    )
})