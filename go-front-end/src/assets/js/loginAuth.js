
import {Get} from "@/assets/js/http";
import {isLogin} from "@/assets/js/utils";

function loginAuth(loginCallback, logoutCallback, expireCallback) {
    if(isLogin()){
        Get("/v1/auth").send(
            ok => {
                if(ok.code === 200){
                    loginCallback()
                }else{
                    expireCallback()
                }
            },
            err => {
                logoutCallback()
            }
        )
    }else{
        logoutCallback()
    }
}

export default loginAuth