module go-service

go 1.16

require github.com/astaxie/beego v1.12.1

// 导入 go path 中 src 下的其他项目
// 就这么写就行了，首先 replace 一下，指向我们的项目目录
// 然后在下面 require 里面直接写项目名称，版本号 v0.0.0
// replace go-redis-template => ../go-redis-template
// 不过我把我的这个包传到 github 上了
//replace go-redis-template => github.com/PtheF/go-redis-template // indirect
require (
	//github.com/PtheF/go-redis-template v0.0.3
	//github.com/PtheF/go-redis-template v0.0.4-0.20220813140948-cf5fa4bfaa3e
	//github.com/PtheF/go-redis-template v0.0.4-0.20221002044323-edc7ed182234
	github.com/PtheF/go-redis-template v0.0.5
	github.com/dlclark/regexp2 v1.7.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	google.golang.org/appengine v1.6.7 // indirect
)
