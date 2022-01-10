package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SendWorkMessage 发送工作通知
// Method: POST
// URL:	https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=ACCESS_TOKEN
// 参数：
func (ding *DingTalkEnter) SendWorkMessage(accessToken string, workMessage *WorkMessage) (*ApiResponse, error) {
	var err error

	// 设置AgentId
	workMessage.AgentId = ding.AgentId

	// 发送人检测
	if len(workMessage.UseridList) == 0 {
		return nil, errors.New("UseridList is empty")
	}

	// 请求url
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", accessToken)
	postData, err := json.Marshal(workMessage)
	if err != nil {
		return nil, err
	}

	c := Client{}
	resp, err := c.Post(url, string(postData))
	if err != nil {
		return nil, err
	}

	// 对结果进行判断
	if resp.Errcode != 0 {
		msg := fmt.Sprintf("获取数据出错，错误代码:%d(%s)", resp.Errcode, resp.Errmsg)
		return nil, errors.New(msg)
	}

	return resp, nil
}
