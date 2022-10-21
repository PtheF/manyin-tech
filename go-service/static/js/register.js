var regButton = "reg"               // 注册按钮id
var id = "id"                       // 用户id
var pw = "pw"                       // 密码框id
var captchaButton = "cpBt"    // 验证码按钮id
var captcha = "cp"             // 验证码框id

var captchaButtonEl = document.getElementById(captchaButton)
var regButtonEl = document.getElementById(regButton)

// 获取验证码
captchaButtonEl.addEventListener("click", () => {
    if (!ValidateForm(id, "电话或邮箱不符合规范", /^1[3|4|5|8][0-9]\d{8}$/, /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)){
        return
    }

    Get("/api/captcha?id=" + document.getElementById(id).value).send(
        (res, _) => {
            setTimeout(() => {
                alert("由于团队太穷没法真正发短信和邮件\n所以用这种方法直接返回验证码\n :(\n 验证码：" + res.message)
            }, 2000)
        },
        err => {
            alert(err)
        }
    )

    ShowInfo("验证码已发送，5分钟内有效")

})


// 注册事件
regButtonEl.addEventListener("click", () => {
    if(
        !ValidateForm(id, "电话或邮箱不符合规范", /^1[3|4|5|8][0-9]\d{8}$/, /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/) ||
        !ValidateForm(pw, "密码不符合规范:包含数字和字母，长度8-20", /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$/) ||
        !ValidateForm(captcha, "验证码格式错误", /^\d{4}$/)
    ){
        return
    }

    Post("/api/register").setForm({
        id: document.getElementById(id).value,
        pw: document.getElementById(pw).value,
        cp: document.getElementById(captcha).value
    }).send(
        res => {
            if(res.code === 200){
                ShowInfo("注册成功")
                setTimeout(() => {
                    location.href = "/"
                }, 500)
            }else{
                ShowError(res.message)
            }
        },

        err => {

        }
    )
})