package model

type Conf struct {
	ClientId     string
	ClientSecret string //GitHub里所获取
	RedirectUrl  string //重定向URL
}

type Token struct {
	AccessToken string `json:"access_token"` //唯一有用，所以只传了这个
}
