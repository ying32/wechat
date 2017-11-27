package kfaccount

type TReplyCommon struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

// TImageReply 图像消息
type TImageReply struct {
	TReplyCommon
	Image struct {
		Media_id string `json:"media_id"`
	} `json:"image"`
}

// TTextReply 文字消息
type TTextReply struct {
	TReplyCommon
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// TVoiceReply 语音消息
type TVoiceReply struct {
	TReplyCommon
	Voice struct {
		Media_id string `json:"media_id"`
	} `json:"voice"`
}

func NewImageReply(touser, mediaid string) *TImageReply {
	im := new(TImageReply)
	im.MsgType = "image"
	im.ToUser = touser
	im.Image.Media_id = mediaid
	return im
}

func NewTextReply(touser, content string) *TTextReply {
	im := new(TTextReply)
	im.MsgType = "text"
	im.ToUser = touser
	im.Text.Content = content
	return im
}

func NewImageReply(touser, mediaid string) *TVoiceReply {
	im := new(TVoiceReply)
	im.MsgType = "voice"
	im.ToUser = touser
	im.Voice.Media_id = mediaid
	return im
}
