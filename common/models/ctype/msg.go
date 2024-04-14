package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type MsgType int8

const (
	TextMsgType MsgType = iota + 1
	ImageMsgType
	VideoMsgType
	FileMsgType
	VoiceMsgType
	VoiceCallMsgType
	VideoCallMsgType
	WithdrawMsgType
	ReplyMsgType
	QuoteMsgType
	AtMsgType
	TipMsgType
)

type Msg struct {
	Type         MsgType       `json:"type"`                   // 消息类型 和msgType一模一样
	TextMsg      *TextMsg      `json:"textMsg,omitempty"`      // 文本消息
	ImageMsg     *ImageMsg     `json:"imageMsg,omitempty"`     // 图片消息
	VideoMsg     *VideoMsg     `json:"videoMsg,omitempty"`     // 视频消息
	FileMsg      *FileMsg      `json:"fileMsg,omitempty"`      // 文件消息
	VoiceMsg     *VoiceMsg     `json:"voiceMsg,omitempty"`     // 语音消息
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg,omitempty"` // 语音通话
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg,omitempty"` // 视频通话
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg,omitempty"`  // 撤回消息
	ReplyMsg     *ReplyMsg     `json:"replyMsg,omitempty"`     // 回复消息
	QuoteMsg     *QuoteMsg     `json:"quoteMsg,omitempty"`     // 引用消息
	AtMsg        *AtMsg        `json:"atMsg,omitempty"`        // @用户的消息 群聊才有
	TipMsg       *TipMsg       `json:"tipMsg,omitempty"`       // 提示消息 一般是不入库的
}

func (msg Msg) MsgPreview() string {
	switch msg.Type {
	case 1:
		return msg.TextMsg.Content
	case 2:
		return "[图片消息] - " + msg.ImageMsg.Title
	case 3:
		return "[视频消息] - " + msg.ImageMsg.Title
	case 4:
		return "[文件消息] - " + msg.FileMsg.Title
	case 5:
		return "[语音消息]"
	case 6:
		return "[语言通话]"
	case 7:
		return "[视频通话]"
	case 8:
		return "[撤回消息] - " + msg.WithdrawMsg.Content
	case 9:
		return "[回复消息] - " + msg.ReplyMsg.Content
	case 10:
		return "[引用消息] - " + msg.QuoteMsg.Content
	case 11:
		return "[@消息] - " + msg.AtMsg.Content
	}
	return "[未知消息]"
}

// Scan 取出来的时候的数据
func (c *Msg) Scan(val interface{}) error {
	err := json.Unmarshal(val.([]byte), c)
	if err != nil {
		return err
	}
	if c.Type == WithdrawMsgType {
		// 如果这个消息是撤回消息，那就不要把原消息带出去
		if c.WithdrawMsg != nil {
			c.WithdrawMsg.OriginMsg = nil
		}
	}
	return nil
}

// Value 入库的数据
func (c Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type TextMsg struct {
	Content string `json:"content"`
}
type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}
type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Type  string `json:"type"` // 文件类型 word
}
type VoiceMsg struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
	Flag      int8      `json:"flag"`      // 标识，标识客户端弹框的模式
	Msg       string    `json:"msg"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
}

// WithdrawMsg 撤回消息
type WithdrawMsg struct {
	Content   string `json:"content"`             // 撤回的提示词
	MsgID     uint   `json:"msgID"`               // 需要撤回的消息id 入参必填
	OriginMsg *Msg   `json:"originMsg,omitempty"` // 原消息  怎么做到，转出去的时候不显示
}
type ReplyMsg struct {
	MsgID         uint      `json:"msgID"`   // 消息id
	Content       string    `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg           *Msg      `json:"msg,omitempty"`
	UserID        uint      `json:"userID"`        // 被回复人的用户id
	UserNickName  string    `json:"userNickName"`  // 被回复人的昵称
	OriginMsgDate time.Time `json:"originMsgDate"` // 原消息的时间
}
type QuoteMsg struct {
	MsgID         uint      `json:"msgID"`   // 消息id
	Content       string    `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg           *Msg      `json:"msg"`
	UserID        uint      `json:"userID"`        // 被回复人的用户id
	UserNickName  string    `json:"userNickName"`  // 被回复人的昵称
	OriginMsgDate time.Time `json:"originMsgDate"` // 原消息的时间
}

// AtMsg @消息
type AtMsg struct {
	UserID  uint   `json:"userID"`
	Content string `json:"content"` // 回复的文本消息
	Msg     *Msg   `json:"msg"`
}

type TipMsg struct {
	Status  string `json:"status"`  // error  success warning info
	Content string `json:"content"` // 提示的内容
}
