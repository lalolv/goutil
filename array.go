package goutil

// InArray 判断一个值是否存在于数组当中
// @arrs 要判断的数组
// @value 要判断的值
func InArray(arrs []interface{}, value interface{}, equalMethod func(interface{}, interface{}) bool) bool {
	if equalMethod == nil {
		equalMethod = generalEqualMethod
	}
	for _, v := range arrs {
		if equalMethod(v, value) {
			return true
		}
	}
	return false
}

// generalEqualMethod 常规的比较方法
// @a
// @b
// @返回他们是否相等
func generalEqualMethod(a interface{}, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

// InStringArray 判断是否在String数组里面
// @arrs
// @value
func InStringArray(arrs []string, value string, equalMethod func(string, string) bool) bool {
	for _, v := range arrs {
		if equalMethod != nil {
			if equalMethod(v, value) {
				return true
			}
		} else {
			if v == value {
				return true
			}
		}
	}
	return false
}

// RemoveDuplicatesInt64 删除重复元素
func RemoveDuplicatesInt64(elements []int64) []int64 {
	// Use map to record duplicates as we find them.
	encountered := map[int64]bool{}
	result := []int64{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
