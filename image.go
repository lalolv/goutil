package goutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取七牛图像比例
func QiniuRatio(qnUrl, pic string) float64 {
	// 获取图像的比例
	ratio := 1.00
	resp, err := http.Get(fmt.Sprintf("http://%s/%s?imageInfo", qnUrl, pic))
	var pInfo map[string]interface{}

	if err != nil {
		fmt.Println("get error: " + err.Error())
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("io error: " + err.Error())
		}
		json.Unmarshal(body, &pInfo)
	}
	width, _ := ToFloat64(pInfo["width"])
	height, _ := ToFloat64(pInfo["height"])
	if width == 0 || height == 1 {
		return 1.00
	}

	ratio = PrecFloat64(height/width, 6)
	if ratio > 0 {
		return ratio
	} else {
		return 1.00
	}

}
