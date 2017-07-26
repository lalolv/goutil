package goutil

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// ToString args: value, precision(only for float)
func ToString(args ...interface{}) (string, error) {
	value := args[0]
	// default
	precision := 12

	switch value.(type) {
	case string:
		v, _ := value.(string)
		return v, nil
	case int:
		v, _ := value.(int)
		return strconv.Itoa(v), nil
	case int32:
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10), nil
	case float32:
		v, _ := value.(float32)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(float64(v), 'f', precision, 64), nil
	case float64:
		v, _ := value.(float64)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(v, 'f', precision, 64), nil
	default:
		return "", errors.New("unknown type")
	}
}

// ToInt 转换为int类型
func ToInt(value interface{}) (int, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.Atoi(v)
	case int:
		v, _ := value.(int)
		return v, nil
	case int32:
		v, _ := value.(int32)
		return int(v), nil
	case int64:
		v, _ := value.(int64)
		return int(v), nil
	case float32:
		v, _ := value.(float32)
		return int(v), nil
	case float64:
		v, _ := value.(float64)
		return int(v), nil
	default:
		return int(0), errors.New("unknown type")
	}
}

// ToInt32 转换为Int32
func ToInt32(value interface{}) (int32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseInt(v, 10, 32)
		return int32(result), err
	case int:
		v, _ := value.(int)
		return int32(v), nil
	case int32:
		v, _ := value.(int32)
		return int32(v), nil
	case int64:
		v, _ := value.(int64)
		return int32(v), nil
	case float32:
		v, _ := value.(float32)
		return int32(v), nil
	case float64:
		v, _ := value.(float64)
		return int32(v), nil
	default:
		return int32(0), errors.New("unknown type")
	}
}

// ToInt64 转换为Int64
func ToInt64(value interface{}) (int64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseInt(v, 10, 32)
	case int:
		v, _ := value.(int)
		return int64(v), nil
	case int32:
		v, _ := value.(int32)
		return int64(v), nil
	case int64:
		v, _ := value.(int64)
		return v, nil
	case float32:
		v, _ := value.(float32)
		return int64(v), nil
	case float64:
		v, _ := value.(float64)
		return int64(v), nil
	default:
		return int64(0), errors.New("unknown type")
	}
}

// ToFloat32 转换为ToFloat32
func ToFloat32(value interface{}) (float32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseFloat(v, 32)
		return float32(result), err
	case int:
		v, _ := value.(int)
		return float32(v), nil
	case int32:
		v, _ := value.(int32)
		return float32(v), nil
	case int64:
		v, _ := value.(int64)
		return float32(v), nil
	case float32:
		v, _ := value.(float32)
		return v, nil
	case float64:
		v, _ := value.(float64)
		return float32(v), nil
	default:
		return float32(0), errors.New("unknown type")
	}
}

// ToFloat64 转换为 float64
func ToFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	default:
		return float64(0), errors.New("unknown type")
	}
}

// U8sInt uint8 转换为 int
func U8sInt(value interface{}) int {
	// nil = 0
	if value == nil {
		return 0
	}
	// err = 0
	r, err := strconv.Atoi(string(value.([]uint8)))
	if err != nil {
		r = 0
	}
	return r
}

// U82Str uint8转换为str
func U82Str(value interface{}) string {
	if t, ok := value.([]uint8); ok {
		return string(t)
	}
	return ""
}

// Bool2Int bool 转换为 int
func Bool2Int(value bool) int {
	switch value {
	case true:
		return 1
	case false:
		return 0
	default:
		return 0
	}
}

// WhenNil 当空值（Map长度为0）的时候返回相应数据
// @value 测试的值
// @nilReturn 空值返回的值
// @okReturn 不为空的时候返回的值
func WhenNil(value interface{}, nilReturn interface{}, okReturn interface{}) interface{} {
	if value == nil {
		return nilReturn
	}
	t := reflect.TypeOf(value).String()
	//只检测map和数组
	if strings.Contains(t, "map") || strings.Contains(t, "[]") {
		v := reflect.ValueOf(value)
		if v.IsNil() || v.Len() == 0 {
			return nilReturn
		}
	}
	return okReturn
}
