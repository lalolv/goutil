package goutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Easemob struct {
	ID     string
	Secret string
}

// 生成环信token
func (p *Easemob) generateToken() bson.M {
	// body params
	params := bson.M{
		"grant_type":    "client_credentials",
		"client_id":     p.ID,
		"client_secret": p.Secret,
	}
	resp := p.doRequest(params, "https://a1.easemob.com/fashionmii/fashionmii/token", "", "POST")
	if resp["errcode"] == 0 {
		return resp["data"].(bson.M)
	} else {
		return bson.M{"access_token": "error", "expires_in": -1}
	}
}

// 执行post请求，返回json
// 执行请求，返回json数据
func (p *Easemob) doRequest(requestData bson.M, requestUrl, aToken, method string) bson.M {
	buf, _ := json.Marshal(requestData)
	body := bytes.NewBuffer(buf)
	req, err := http.NewRequest(method, requestUrl, body)
	// Header
	req.Header.Add("Content-Type", "application/json")
	if aToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", aToken))
	}

	// resp
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	data := bson.M{}
	errcode := 0
	// get data
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			r, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(r, &data)
		} else {
			data = bson.M{"HTTP_CODE": resp.StatusCode}
		}
	} else {
		errcode = 1
	}
	return bson.M{"data": data, "errcode": errcode}
}

// 环信注册单个用户
// @uid
// @passwd
// @nickName
func (p *Easemob) EasemobSignupSingle(uid int64, passwd, nickName string) bool {
	// body params
	params := bson.M{
		"username": uid,
		"password": passwd,
		"nickname": nickName,
	}
	// 生成token
	aToken := p.generateToken()
	if aToken["access_token"] == "error" {
		fmt.Println("生成token失败")
		return false
	}

	data := p.doRequest(
		params,
		"https://a1.easemob.com/fashionmii/fashionmii/users",
		aToken["access_token"].(string), "POST",
	)

	if data["errcode"] == 1 {
		fmt.Println("注册失败")
		return false
	}
	return true
}

// 环信修改用户昵称
// @userName
// @passwd
// @nickName
// @aToken
func (p *Easemob) EasemobNickModify(uid int64, nickName string) bool {
	// body params
	params := bson.M{
		"nickname": nickName,
	}

	aToken := p.generateToken()
	if aToken["access_token"] == "error" {
		fmt.Println("生成token失败")
		return false
	}

	data := p.doRequest(
		params,
		fmt.Sprintf("https://a1.easemob.com/fashionmii/fashionmii/users/%d", uid),
		aToken["access_token"].(string), "PUT")

	if data["errcode"] == 1 {
		return false
	}
	return true
}

// 发送消息
func (p *Easemob) PushMessage(msg string, target map[string]string) (interface{}, error) {
	// body params
	params := bson.M{
		"target_type": "users",
		"target":      target,
		"msg": map[string]string{
			"type": "txt",
			"msg":  msg,
		},
		"from": "systerm",
	}

	aToken := p.generateToken()
	if aToken["access_token"] == "error" {
		fmt.Println("生成token失败")
		return nil, errors.New("生成token失败")
	}

	data := p.doRequest(
		params,
		fmt.Sprintf("https://a1.easemob.com/fashionmii/fashionmii/messages"),
		aToken["access_token"].(string), "POST")

	return data, nil
}
