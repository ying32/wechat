package wechat

import (
	"net/http"
	"sync"

	"github.com/ying32/wechat/cache"
	"github.com/ying32/wechat/context"
	"github.com/ying32/wechat/js"
	"github.com/ying32/wechat/material"
	"github.com/ying32/wechat/menu"
	"github.com/ying32/wechat/oauth"
	"github.com/ying32/wechat/qrcode"
	"github.com/ying32/wechat/server"
	"github.com/ying32/wechat/templatemsg"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	// 默认的模板消息id
	TemplateMsgId string
	Cache         cache.Cache
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.TemplateMsgId = cfg.TemplateMsgId
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth(req *http.Request, writer http.ResponseWriter) *oauth.Oauth {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return oauth.NewOauth(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs(req *http.Request, writer http.ResponseWriter) *js.Js {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return js.NewJs(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu(req *http.Request, writer http.ResponseWriter) *menu.Menu {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return menu.NewMenu(wc.Context)
}

// GetTemplateMsg 业务模板消息
func (wc *Wechat) GetTemplateMsg() *templatemsg.TTemplateMsg {
	return templatemsg.NewTemplateMsg(wc.Context)
}

// GetQRCode QRCode
func (wc *Wechat) GetQRCode() *qrcode.TGenQRCode {
	return qrcode.NewGenQRCode(wc.Context)
}
