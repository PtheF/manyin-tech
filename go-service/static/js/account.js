var saveBut = "save"
var cancelSave = "cancel-update"

var captcha = "captcha"
var doSaveBut = "update-account"
var captchaDiv = "captcha-div"

var saveButEl = document.getElementById(saveBut)
var cancelSaveEl = document.getElementById(cancelSave)

var captchaEl = document.getElementById(captcha)
var doSaveButEl = document.getElementById(doSaveBut)
var captchaDivEl = document.getElementById(captchaDiv)

saveButEl.addEventListener("click", () => {
    captchaDivEl.style.display = "flex"
})

cancelSaveEl.addEventListener("click", () => {
    captchaDivEl.style.display = "none"
})
