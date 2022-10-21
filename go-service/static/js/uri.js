
function uri() {
    return window.location.href
}

function getQueries() {
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

function queryString(key) {
    return getQueries()[key]
}