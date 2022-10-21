var addBut = "show-textarea"
var newAddressDiv = "add"

var addAds = "add-address"
var adsValue = "new-address"

addButEl = document.getElementById(addBut)
newAddressDivEl = document.getElementById(newAddressDiv)
addAdsEl = document.getElementById(addAds)
adsValueEl = document.getElementById(adsValue)

var newAddressDivShow = false

function deleteAddress(id){
    Delete("/api/home/address/" + id).send(
        ok => {
            if(ok.code === 200){
                ShowInfo(ok.message)
            }else{
                ShowError(ok.message)
            }
        },

        err => {
            ShowError(err)
        }
    )
}

function setDefault(id) {
    Post("/api/home/address/" + id).send(
        ok => {
            if(ok.code === 200){
                ShowInfo(ok.message)
            }else{
                ShowError(ok.message)
            }
        },
        err => {
            ShowError(err)
        }
    )
}

addButEl.addEventListener("click", () => {

    if(newAddressDivShow){
        newAddressDivEl.style.opacity = "0"
        newAddressDivEl.style.display = "none"
        addButEl.innerText = "添加"
    }else{
        newAddressDivEl.style.display = "block"
        newAddressDivEl.style.opacity = "1"
        addButEl.innerText = "取消"
    }

    newAddressDivShow = !newAddressDivShow
})


addAdsEl.addEventListener("click", () => {
    let address = adsValueEl.value
    if(address === ""){
        ShowError("地址不能为空")
        return
    }
    Put("/api/home/address").setForm({address: address}).send(
        res => {
            if(res.code === 200){
                ShowInfo("添加成功")
            }else{
                ShowError(res.message)
            }
        },

        err => {
            ShowError(err)
        }
    )
})