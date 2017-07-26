package goutil

import "regexp"

// CheckPhone 检测手机号码
func CheckPhone(tel string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)

	return rgx.MatchString(tel)
}
