package kfaccount

type TReply struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

type TImageReply struct {
	TReply
	Image struct {
		Media_id string `json:"media_id"`
	} `json:"image"`
}

func NewImageReply(touser, mediaid string) *TImageReply {
	im := new(TImageReply)
	im.MsgType = "image"
	im.ToUser = touser
	im.Image.Media_id = mediaid
	return im
}
