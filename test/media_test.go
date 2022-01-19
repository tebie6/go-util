package test

import (
	"fmt"
	"github.com/tebie6/go-util/dingtalk"
	"log"
	"os"
	"testing"
)

func Test_Upload(t *testing.T) {

	path := "/Users/liumingyu/Downloads/upload_test.xlsx" //要上传文件所在路径
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	fileType := "file" // 根据钉钉文档填写

	ding := dingtalk.DingTalkEnter{
		AgentId:   0,
		AppKey:    "",
		AppSecret: "",
	}

	token, err := ding.GetAccessToken()
	if err != nil {
		t.Error(err.Error())
	} else {
		log.Println("Token:", token)
	}

	resp, err := ding.MediaUpload(token, fileType, file.Name(), file)
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(fmt.Sprintf("上传成功 %v", resp))

}
