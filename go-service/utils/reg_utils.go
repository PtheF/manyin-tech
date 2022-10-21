package utils

import (
	"github.com/dlclark/regexp2"
)

const (
	// 电话正则
	//PHONE_REG = `^1[3|4|5|6|7|8][0-9]\d{8}$`
	PHONE_REG = `^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`

	// 邮箱正则
	MAIL_REG = `^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`

	// 密码正则：必须包含数字和字母，长度 8-20
	PASS_REG = `^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,20}$`
)

func RegCheck(t string, regs ...string) int {

	for i, reg := range regs {

		// 这里需要用到 github 上的第三方 go 包：regexp2
		// 因为原来的 regexp 不支持上面那种 ?! 这种复杂写法，所以要用到这个包
		_reg, _ := regexp2.Compile(reg, 0)

		if res, _ := _reg.MatchString(t); res {
			return i + 1
		}
	}

	return 0

}
