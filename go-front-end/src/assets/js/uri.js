
export const uri = function() {
    return window.location.href
}

export const getQueries = function() {
    let href = uri()

    if(href.indexOf("?") === -1){
        return {}
    }

    let queryStrings = href.substring(href.indexOf("?") + 1, href.length).split("&")

    let res = {}

    console.log("queryStrings=" + queryStrings)

    for (let queryString of queryStrings) {
        console.log(queryString)

        let kv = queryString.split("=")
        res[kv[0].trim()] = kv[1].trim()
    }

    return res
}

export const queryString = function(key) {
    return getQueries()[key]
}