// qrcode by: ying32  2017-6-15
package qrcode

import (
	"encoding/json"
	"net/url"
	//"encoding/json"
	"fmt"

	"github.com/ying32/wechat/context"
	"github.com/ying32/wechat/util"
)

const (
	// genQRCode 生成qrcode ticket
	genQRCode = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"

	// downloadQRCode 下载qrcode， 返回为一个图片
	downloadQRCode = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)

// sceneParam
type sceneParam struct {
	SceneStr string `json:"scene_str"`
}

// scene
type scene struct {
	Scene sceneParam `json:"scene"`
}

// tempQRCode 临时qrcode参数
type tempQRCode struct {
	ExpireSeconds int        `json:"expire_seconds"`
	ActionName    string     `json:"action_name"`
	ActionInfo    sceneParam `json:"action_info"`
}

// permanentQRCode  永久qrcode参数
type permanentQRCode struct {
	ActionName string     `json:"action_name"`
	ActionInfo sceneParam `json:"action_info"`
}

// Ticket
type Ticket struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	URL           string `json:"url"`
}

// 临时
// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// 永久
// {"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}

// TGenQRCode 业务模板消息
type TGenQRCode struct {
	*context.Context
}

// NewGenQRCode
func NewGenQRCode(context *context.Context) *TGenQRCode {
	gqr := new(TGenQRCode)
	gqr.Context = context
	return gqr
}

// getQRCodeTicket
func (g TGenQRCode) getQRCodeTicket(v interface{}) (*Ticket, error) {
	token, err := g.Context.GetAccessToken()
	if err != nil {
		return nil, err
	}
	retBytes, err := util.PostJSON(fmt.Sprintf(genQRCode, token), v)
	if err != nil {
		return nil, err
	}
	result := new(Ticket)
	err = json.Unmarshal(retBytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTempQRCodeTicket 获取临时QRCode Ticket
func (g TGenQRCode) GetTempQRCodeTicket(expireSeconds int, scene string) (*Ticket, error) {
	v := tempQRCode{}
	v.ExpireSeconds = expireSeconds
	v.ActionName = "QR_STR_SCENE"
	v.ActionInfo.SceneStr = scene
	return g.getQRCodeTicket(v)
}

// GetPermanentQRCodeTicket 获取永久性QRCode Ticket
func (g TGenQRCode) GetPermanentQRCodeTicket(scene string) (*Ticket, error) {
	v := permanentQRCode{}
	v.ActionName = "QR_STR_SCENE"
	v.ActionInfo.SceneStr = scene
	return g.getQRCodeTicket(v)
}

// DownloadQRCode 下载QRCode，返回的是一个图片
func (g TGenQRCode) DownloadQRCode(ticket string) {
	// 还未做 按需求需要使用urlcode编码ticket
	util.HTTPGet(fmt.Sprintf(downloadQRCode, url.QueryEscape(ticket)))
}
