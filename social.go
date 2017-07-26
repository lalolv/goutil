package goutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 通过 QQ 接口，获取用户资料
// @appID
// @accessToken
// @openID
func QQUserInfo(appID, accessToken, openID string) map[string]interface{} {

	// qqAPI接口
	url := "https://openmobile.qq.com/user/get_simple_userinfo?"
	url += "access_token=" + accessToken + "&"
	url += "oauth_consumer_key=" + appID + "&"
	url += "openid=" + openID + "&format=json"
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("qq api get error: " + err.Error())
	}

	// 读取数据资料
	var uInfo map[string]interface{}
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("qq api io error: " + err.Error())
		}

		json.Unmarshal(body, &uInfo)
	}

	return uInfo
}

// 通过微信接口，获取用户资料
// @accessToken
// @openID
// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID
func WechatUserInfo(accessToken, openID string) map[string]interface{} {

	// qqAPI接口
	url := "https://api.weixin.qq.com/sns/userinfo?"
	url += "access_token=" + accessToken + "&"
	url += "openid=" + openID
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("weixin api get error: " + err.Error())
	}

	// 读取数据资料
	var uInfo map[string]interface{}
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("weixin api io error: " + err.Error())
		}

		json.Unmarshal(body, &uInfo)
	}

	return uInfo
}

// 获取微博用户数据
// @accessToken
// @uid
func WeiboUserInfo(accessToken, uid string) map[string]interface{} {

	// qqAPI接口
	url := "https://api.weibo.com/2/users/show.json?"
	url += "access_token=" + accessToken + "&uid=" + uid
	// fmt.Println(url)
	resp, err := http.Get(url)
	var uInfo map[string]interface{}

	if err != nil {
		fmt.Println("get error: " + err.Error())
		uInfo = map[string]interface{}{"error": 1}
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("io error: " + err.Error())
		}
		json.Unmarshal(body, &uInfo)
	}

	return uInfo
}
