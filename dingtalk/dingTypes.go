package dingtalk

// 企业内部应用
type DingTalkEnter struct {
	AgentId   int64  `json:"agent_id"`   // 钉钉应用ID
	AppKey    string `json:"app_key"`    // 应用唯一KEY
	AppSecret string `json:"app_secret"` // 应用密钥
}

type ApiResponse struct {
	Errcode     int64  `json:"errcode"`                        // 错误代码，无错误代码是0
	Errmsg      string `json:"errmsg"`                         // 错误消息
	AccessToken string `json:"access_token,omitempty"`         // Access Token
	ExpiresIn   string `json:"access_token_expires,omitempty"` // Access Token 有效期
	TaskId      int64  `json:"task_id,omitempty"`              // 任务ID，比如：发送消息
	RequestID   string `json:"request_id,omitempty"`           // 请求ID，比如：发送消息的响应
	Result      struct {
		Userid string `json:"userid,omitempty"` // 用户ID，比如：获取用户信息
	} `json:"result,omitempty"`
	MediaId string `json:"media_id,omitempty"` // 媒体文件上传后获取的唯一标识。
	Type    string `json:"type,omitempty"`     // 媒体文件类型：image：图片、voice：语音、file：普通文件、video：视频
}

// 消息类型与数据格式
// 参考文档：https://open.dingtalk.com/document/orgapp-server/message-types-and-data-format
type DingTalkMessage struct {
	MsgType string   `json:"msgtype"`        // 消息类型
	Text    *TextMsg `json:"text,omitempty"` // MsgType是text的消息内容
	File    *FileMsg `json:"file,omitempty"` // MsgType是file的消息内容
}

// 文本消息（text）
type TextMsg struct {
	Content string `json:"content"` // 文本消息内容
}

type FileMsg struct {
	MediaId string `json:"media_id"` // 文件消息内容
}

// 发送工作通知
type WorkMessage struct {
	AgentId    int64            `json:"agent_id"`    //【必填】发送消息时使用的微应用的AgentID。
	UseridList string           `json:"userid_list"` // 接收者的userid列表，最大用户列表长度100。
	ToAllUser  bool             `json:"to_all_user"` // 是否发送给企业全部用户。
	Msg        *DingTalkMessage `json:"msg"`         // 消息内容，最长不超过2048个字节。
}
