package goutil

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// PrecFloat64 设置精度
func PrecFloat64(f float64, p int) float64 {
	// 设置精度
	strR := strconv.FormatFloat(f, 'f', p, 64)
	// 转换为64为小数
	nr, _ := strconv.ParseFloat(strR, 64)

	return nr
}

// RandomFloat64 随机生成64位小数
// @min 最小值
// @max 最大值
// @p 小数精度
func RandomFloat64(min, max float64, p int) float64 {
	rand.Seed(time.Now().UnixNano())
	r0 := rand.Float64()
	r := r0*(max-min) + min
	// 设置精度
	strR := strconv.FormatFloat(r, 'f', p, 64)
	// 转换为64为小数
	nr, _ := strconv.ParseFloat(strR, 64)

	return nr
}

// RandomSpec Creates a random string based on a variety of options, using
// supplied source of randomness.
//
// If start and end are both 0, start and end are set
// to ' ' and 'z', the ASCII printable
// characters, will be used, unless letters and numbers are both
// false, in which case, start and end are set to 0 and math.MaxInt32.
//
// If set is not nil, characters between start and end are chosen.
//
// This method accepts a user-supplied rand.Rand
// instance to use as a source of randomness. By seeding a single
// rand.Rand instance with a fixed seed and using it for each call,
// the same random sequence of strings can be generated repeatedly
// and predictably.
func RandomSpec(count uint, start, end int, letters, numbers bool, chars []rune) string {
	// default random
	var defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	// count
	if count == 0 {
		return ""
	}

	if start == 0 && end == 0 {
		end = 'z' + 1
		start = ' '
		if !letters && !numbers {
			start = 0
			end = math.MaxInt32
		}
	}

	buffer := make([]rune, count)
	gap := end - start
	for count != 0 {
		count--
		var ch rune
		if len(chars) == 0 {
			ch = rune(defaultRand.Intn(gap) + start)
		} else {
			ch = chars[defaultRand.Intn(gap)+start]
		}

		if letters && ((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) ||
			numbers && (ch >= '0' && ch <= '9') ||
			(!letters && !numbers) {
			if ch >= rune(56320) && ch <= rune(57343) {
				if count == 0 {
					count++
				} else {
					buffer[count] = ch
					count--
					buffer[count] = rune(55296 + defaultRand.Intn(128))
				}

			} else if ch >= rune(55296) && ch <= rune(56191) {
				if count == 0 {
					count++
				} else {
					// high surrogate, insert low surrogate before putting it in
					buffer[count] = rune(56320 + defaultRand.Intn(128))
					count--
					buffer[count] = ch
				}
			} else if ch >= rune(56192) && ch <= rune(56319) {
				// private high surrogate, no effing clue, so skip it
				count++
			} else {
				buffer[count] = ch
			}
		} else {
			count++
		}
	}

	return string(buffer)
}

// TinyNo 短单号
func TinyNo() string {
	now := time.Now()
	tt := now.Local().Format("06-01-02-15-04-05")

	var outstr string
	for _, t := range strings.Split(tt, "-") {
		tn, _ := ToInt(t)
		if tn >= 10 && tn < 26 {
			tr := rune(tn + 65)
			outstr += string(tr)
		} else if tn >= 48 && tn <= 57 {
			tr := rune(tn)
			outstr += string(tr)
		} else {
			outstr += strings.TrimLeft(t, "0")
		}
	}

	return outstr
}
