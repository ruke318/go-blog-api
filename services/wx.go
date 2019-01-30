package services

import (
	"blog/config"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var (
	conf config.Config
	loginInfo map[string]interface{}
)

type WxApp struct {
	Appid string
	Secret string
}

func (wx *WxApp) GetConfig() *config.WxApp {
	conf = config.Load()
	return conf.WxApp
}

func (wx *WxApp) Login(code string) interface{} {
	con := wx.GetConfig()
	loginApi := "https://api.weixin.qq.com/sns/jscode2session?appid="+ con.Appid +"&secret="+ con.Secret +"&js_code="+ code +"&grant_type=authorization_code"
	resp, err := http.Get(loginApi)
	for k := range loginInfo {
			delete(loginInfo, k)
	}
	if err != nil {
		str := `{"errcode": -2, "errmsg": "请求失败"}`
		json.Unmarshal([]byte(str), &loginInfo)
	return loginInfo
	} else {
		defer resp.Body.Close()
		str, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(str), &loginInfo)
		return loginInfo
	}
}
