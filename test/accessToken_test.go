package test

import (
	"github.com/tebie6/go-util/dingtalk"
	"log"
	"testing"
)

// 测试获取Access Token
func Test_getAccessToken(t *testing.T)  {

	ding := dingtalk.DingTalkEnter{
		AgentId:     0,
		AppKey:      "",
		AppSecret:   "",
		AccessToken: "",
	}

	if token, err := ding.GetAccessToken(); err != nil {
		t.Error(err.Error())
	} else {
		log.Println("Token:", token)
	}
}
