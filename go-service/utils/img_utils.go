package utils

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"regexp"
)

const ImgBase64Prefix = `^data:\s*image\/(\w+);base64,`

func SaveBase64Img(imgName string, path string, imgBase64 string) (error, string) {
	if b, _ := regexp.MatchString(ImgBase64Prefix, imgBase64); !b {
		return errors.New("not a image base64 string"), ""
	}

	re, _ := regexp.Compile(ImgBase64Prefix)

	allData := re.FindAllSubmatch([]byte(imgBase64), 2)

	imgType := string(allData[0][1])

	base64Str := re.ReplaceAllString(imgBase64, "")

	by, _ := base64.StdEncoding.DecodeString(base64Str)

	f := path + imgName + "." + imgType

	err := ioutil.WriteFile(f, by, 0666)

	if err != nil {
		return err, ""
	}

	return nil, imgName + "." + imgType

}
