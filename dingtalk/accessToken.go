package dingtalk

import (
	"errors"
	"fmt"
)

// GetAccessToken 获取AccessToken 一切操作的前提
// Method: GET
// URL：https://oapi.dingtalk.com/gettoken?appkey=key&appsecret=secret
// 参数：
func (ding *DingTalkEnter) GetAccessToken() (accessToken string, err error) {

	// 通过API获取AccessToken
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", ding.AppKey, ding.AppSecret)

	c := &Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "", err
	}
	if resp.Errcode != 0 {
		msg := fmt.Sprintf("获取AccessToken出错, 错误代码是:%d,错误消息是:%s",
			resp.Errcode, resp.Errmsg)
		return "", errors.New(msg)
	}

	return resp.AccessToken, nil
}
