
var ValidateElList = {
    els: [],
    num: 0
}

function newEl(elementStr){
    return {
        elStr: elementStr,
        top: -50,
        el: null,

        create: function() {
            document.body.insertAdjacentHTML("afterend", this.elStr)
            this.el = document.getElementById("validate-div")
        },

        moveTo: function(top) {
            this.top = top
            this.el.style.top = top + "px"
        },

        fade: function(){
            this.el.style.top = this.top + 50 + "px"
            this.el.style.opacity = "0"
            setTimeout(() => {
                this.el.remove()
            }, 100)
        }
    }
}

function newErrorEl(text) {
    const elStr = `<div id="validate-div" style="position: fixed; top: -50px; left: 0; right: 0; background: transparent; text-align: center; transition: .2s;">
        <span style="width: auto; margin: 0 auto 0; padding: 10px 30px 10px 30px; background-color: #ff8383; border-radius: 10px; color: white; text-align: center; transition: .2s;">
            ${text}
        </span>
    </div>`

    return newEl(elStr)
}

function newSucEl(text){
    const elStr = `<div id="validate-div" style="position: fixed; top: -50px; left: 0; right: 0; background: transparent; text-align: center; transition: .2s;">
        <span style="width: auto; margin: 0 auto 0; padding: 10px 30px 10px 30px; background-color: #00c51f; border-radius: 10px; color: #ffffff; text-align: center; transition: .2s;">
            ${text}
        </span>
    </div>`

    return newEl(elStr)
}

function display(curEl) {

    if(ValidateElList.num === 6){
        alert("帕金森？？？")
    }

    if(ValidateElList.els.length !== 0){
        for(let el of ValidateElList.els){
            el.moveTo(el.top + 50)
        }
    }

    curEl.create()

    setTimeout(() => {
        curEl.moveTo(50)
        ValidateElList.els.push(curEl)
        ValidateElList.num ++
    }, 100)

    setTimeout(() => {
        curEl.fade()
        ValidateElList.els.shift()
        ValidateElList.num --
    }, 2000)
}

function ShowInfo(text){
    display(newSucEl(text))
}

function ShowError(text){
    display(newErrorEl(text))
}

function ValidateForm(id, text, ...regs){
    let value = document.getElementById(id).value

    let res = false

    let i = 0
    for(; i < regs.length && !res; i++){
        res ||= regs[i].test(value)
    }

    if(!res){
        ShowError(text)
        return 0
    }

    return i
}
