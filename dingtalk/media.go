package dingtalk

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
)

func (ding *DingTalkEnter) MediaUpload(accessToken string, fileType string, fileName string, file multipart.File) (*ApiResponse, error) {

	if len(accessToken) == 0 {
		return nil, errors.New("accessToken is empty")
	}

	if len(fileType) == 0 {
		return nil, errors.New("fileType is empty")
	}

	if len(fileName) == 0 {
		return nil, errors.New("fileName is empty")
	}

	// 声明一个 bytes的buffer缓冲器
	body := &bytes.Buffer{}

	// 返回一个设定了一个随机边界的Writer writer，并将数据写入&body
	writer := multipart.NewWriter(body)
	defer func() {
		_ = writer.Close()
	}()

	// 创建一个 form-data类型的文件字段
	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		return nil, err
	}

	// 将接收的文件内容 copy 到创建的文件字段内
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// 添加一个新的 form-data 字段
	_ = writer.WriteField("type", fileType)

	// Close方法结束multipart信息，并将结尾的边界写入底层io.Writer接口。
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://oapi.dingtalk.com/media/upload?access_token=%s", accessToken)

	c := Client{}
	resp, err := c.PostFormData(url, writer.FormDataContentType(), body)
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
