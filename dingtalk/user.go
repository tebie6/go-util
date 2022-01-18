package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tebie6/go-util/tools/verify"
)

func (ding *DingTalkEnter) GetUserByMobile(accessToken string, mobile string) (*ApiResponse, error) {

	if len(accessToken) == 0 {
		return nil, errors.New("accessToken is empty")
	}

	if len(mobile) == 0 {
		return nil, errors.New("mobile is empty")
	}

	matched := verify.VerifyMobileFormat(mobile)
	if !matched {
		return nil, errors.New("wrong format of mobile")
	}

	// 请求url
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/v2/user/getbymobile?access_token=%s", accessToken)
	postData, err := json.Marshal(struct {
		Mobile string `json:"mobile"`
	}{
		Mobile: mobile,
	})
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
