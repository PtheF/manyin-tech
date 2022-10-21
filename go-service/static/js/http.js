function newHttp() {
    let http = {

        url: "",
        method: "",
        json: {},
        form: {},
        header: {},
        auth: "",

        refreshTokenHeader: "",
        refreshTokenSaver: null,

        isJson: false,
        isForm: false,

        formHeader: "application/x-www-form-urlencoded;charset=utf-8",
        jsonHeader: "application/json;charset=utf-8",

        setJson: function(jsonData) {
            if(this.isForm) {
                throw Error("不能同时传表单和json")
            }
            this.json = jsonData
            this.isJson = true
            return this
        },

        setForm: function(formData) {
            if(this.isJson) {
                throw Error("不能同时传表单和json")
            }

            this.isForm = true
            this.form = formData
            return this
        },

        setHeader: function(headers) {
            this.header = headers
            return this
        },

        // callback: function() -> string
        setAuth: function(callback){
            this.auth = callback()
            return this
        },

        // callback: function(jwt) -> void
        refreshToken: function(header, saver){
            this.isRefreshToken = "header"
            this.refreshTokenSaver = saver
            return this
        },

        send: function(ok, err) {

            function temp() {
                let xhr;

                if(window.XMLHttpRequest) {
                    xhr = new XMLHttpRequest();
                }else{
                    xhr = new ActiveXObject("Microsoft.XMLHttp")
                }

                console.log("xhr: ", xhr)

                xhr.onreadystatechange = function() {
                    if(this.readyState === 4 && this.status === 200) {
                        let res = JSON.parse(xhr.responseText)

                        if(http.refreshTokenHeader !== ""){
                            http.refreshTokenSaver(xhr.getResponseHeader(http.refreshTokenHeader))
                        }

                        ok(res)
                    }
                }

                xhr.open(http.method, http.url, true)

                if(http.header !== {}){
                    for(let eachHeader in http.header){
                        xhr.setRequestHeader(eachHeader, http.header[eachHeader])
                    }
                }

                if(http.auth !== ""){
                    xhr.setRequestHeader("Authorization", http.auth)
                }


                let data = ""

                if(http.method === "POST" || http.method === "PUT"){
                    if(http.isJson && !http.isForm) {
                        xhr.setRequestHeader("Content-Type", http.jsonHeader)
                        data = JSON.stringify(http.json)
                    }else if (http.isForm && !http.isJson) {
                        xhr.setRequestHeader("Content-Type", http.formHeader)

                        for(let eachItem in http.form) {
                            data += eachItem + "=" + http.form[eachItem] + "&"
                        }

                        data = data.slice(0, data.length - 1)
                    }
                }

                if(data === ""){
                    xhr.send()
                }else {
                    console.log("data: ", data)
                    xhr.send(data)
                }

            }

            try {
                temp()
            }catch(e){
                err(e)
            }


        }
    }

    return http
}

function Get(url) {
    let http = newHttp()
    http.url = url
    http.method = "GET"

    return http
}

function Post(url) {
    let http = newHttp()
    http.url = url
    http.method = "POST"

    return http
}

function Put(url) {
    let http = newHttp()
    http.url = url
    http.method = "PUT"
    return http
}

function Delete(url) {
    let http = newHttp()
    http.url = url
    http.method = "DELETE"
    return http
}