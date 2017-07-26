package goutil

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	nan string = "nan"
)

var white = regexp.MustCompile(`^\s*$`)

// Str2Int 字符串装换为整数类型
func Str2Int(s string) int {
	// Nan convert to 0
	if strings.ToLower(s) == nan {
		return 0
	}
	// convert to int
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}

	return int(i)
}

// Str2Int64 字符串装换为整数类型
func Str2Int64(s string) int64 {
	// Nan convert to 0
	if strings.ToLower(s) == nan {
		return 0
	}
	// convert to int
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return int64(i)
}

// Str2Uint 字符串装换为整数类型
func Str2Uint(s string) uint {
	// Nan convert to 0
	if strings.ToLower(s) == nan {
		return 0
	}
	// convert to int
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}

	return uint(i)
}

// Str2Float 字符串装换为浮点类型
func Str2Float(s string) float64 {
	// Nan convert to 0
	if strings.ToLower(s) == nan {
		return 0
	}
	// convert to float
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return float64(f)
}

// SubstrByByte 按字节截取字符串 utf-8不乱码
func SubstrByByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			var cl int
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

// IsEmpty 利用正则表达式判断一个字符串是不是空
// @val要判断的值
func IsEmpty(val string) bool {
	return white.MatchString(val)
}

// EscapeQuots 转义引号
// 单引号和双引号
func EscapeQuots(val string) string {
	str := strings.Replace(val, `'`, `\'`, -1)
	str = strings.Replace(str, `"`, `\"`, -1)

	return str
}
