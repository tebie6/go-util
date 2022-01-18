package test

import (
	"fmt"
	"github.com/tebie6/go-util/dingtalk"
	"log"
	"testing"
)

func Test_Getbymobile(t *testing.T) {

	ding := dingtalk.DingTalkEnter{
		AgentId:     0,
		AppKey:      "",
		AppSecret:   "",
	}

	token, err := ding.GetAccessToken()
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Println("Token:", token)
	}

	resp, err := ding.GetUserByMobile(token, "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(fmt.Sprintf("获取成功 %v", resp))
}
