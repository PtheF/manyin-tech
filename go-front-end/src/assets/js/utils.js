export const isEmpty = function (val) {
    return typeof val === 'undefined' || val === null || val === '';
}

export const removeArray = function(array, index) {
    if(isNaN(index)|| index > array.length){
        return false;
    }
    for(; index < array.length-1; index++) {
        array[index] = array[index+1]
    }
    array.pop()

    return array
}

const TOKEN_KEY = "token:"
const USER_KEY = "user"
const ENTER_KEY = "enterprise:"

export const isLogin = () => {

    return !isEmpty(sessionStorage.getItem(USER_KEY)) && getUser() !== null
}

export const getUser = () => {
    let userJson = sessionStorage.getItem(USER_KEY)
    return JSON.parse(userJson)
}

export const setUser = (user) => {
    sessionStorage.setItem(USER_KEY, JSON.stringify(user))
}

export const removeUser = () => {
    sessionStorage.removeItem(USER_KEY)
}

export const getToken = () => {
    return sessionStorage.getItem(TOKEN_KEY)
}

export const setToken = (token) => {
    sessionStorage.setItem(TOKEN_KEY, token)
}

export const removeToken = () => {
    sessionStorage.removeItem(TOKEN_KEY)
}

export const getEnterprise = () => {
    return JSON.parse(sessionStorage.getItem(ENTER_KEY + getUser().uuid))
}

export const setEnterprise = (enterprise) => {
    sessionStorage.setItem(ENTER_KEY + getUser().uuid, JSON.stringify(enterprise))
}

export const debounce = function(callback, delay_) {
    let delay = delay_ || 200;
    let timer;

    return function() {

        let th = this
        let args = arguments

        if(timer) {
            clearTimeout(timer)
        }

        timer = setTimeout(() => {

            console.log("debounce function set timeout inner this:")
            console.log(this)

            callback.apply(this, args)
        }, delay)
    }
}


