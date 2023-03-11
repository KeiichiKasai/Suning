package service

import (
	"encoding/json"
	"fmt"
	"main.go/model"
	"net/http"
)

var conf = model.Conf{
	ClientId:     "3eb64e7a2638c8fd61b0",
	ClientSecret: "c15d00cffa9caece4aa5a7b18eeee4a99fb4fd06",
	RedirectUrl:  "http://localhost:8080/oauth/redirect",
}

// GetTokenAuthUrl 获取地址
func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code,
	)
}

// GetToken 获取token
func GetToken(url string) (*model.Token, error) {

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodPost, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		fmt.Println("未获取到响应")
		return nil, err
	}
	fmt.Println("获取到了响应")
	fmt.Println(res.Body)
	// 将响应体解析为 token，并返回
	var token model.Token

	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(token *model.Token) (map[string]interface{}, error) {

	// 形成请求
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
