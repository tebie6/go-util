package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
}

func (c *Client) Get(url string) (*ApiResponse, error) {

	res := &ApiResponse{}
	resp, err := http.Get(url)
	if err != nil {
		return res, fmt.Errorf("http get request error = %v", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.Errcode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.Errmsg)
	}

	return res, nil
}

func (c *Client) Post(url string, postData string) (*ApiResponse, error) {

	res := &ApiResponse{}

	contentType := "application/json"
	resp, err := http.Post(url, contentType, strings.NewReader(postData))
	if err != nil {
		return res, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.Errcode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.Errmsg)
	}

	return res, nil
}

// PostFormData 哪位大神能帮忙写一个公共的POST 感激不尽
func (c *Client) PostFormData(url string, contentType string, postData *bytes.Buffer) (*ApiResponse, error) {

	res := &ApiResponse{}

	resp, err := http.Post(url, contentType, postData)
	if err != nil {
		return res, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.Errcode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.Errmsg)
	}

	return res, nil
}
