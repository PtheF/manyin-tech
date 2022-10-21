import { ErrorMessage } from "./message"


export const REG = {
    // PHONE_REG: /^1[3|4|5|7|8|9][0-9]\d{8}$/,
    PHONE_REG: /^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$/,

    EMAIL_REG: /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,

    /** 密码正则： 有数字和字母，长度 8 - 20 */
    PASS_REG: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$/
}

export const Verify = (value, ...regs) => {
    let res = false
    let i = 0

    console.log("verify: ", value);

    for(; i < regs.length && !res; i++){
        res ||= regs[i].test(value)
        console.log("verify res: ", res);
    }

    if(!res){
        return 0
    }

    return i
}