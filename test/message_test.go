package test

import (
	"fmt"
	"github.com/tebie6/go-util/dingtalk"
	"log"
	"testing"
	"time"
)

func Test_SendWorkMessage(t *testing.T)  {

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

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	msg := &dingtalk.DingTalkMessage{
		MsgType: "text",
		Text:    &dingtalk.TextMsg{
			Content: fmt.Sprintf("这个是消息内容:%s", nowStr),
		},
	}

	workMessage := dingtalk.WorkMessage{
		AgentId:    ding.AgentId,
		UseridList: "",
		ToAllUser:  true,
		Msg:        msg,
	}

	resp, err := ding.SendWorkMessage(&workMessage)
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(fmt.Sprintf("发送成功 %v", resp))
}