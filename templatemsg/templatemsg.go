// by: ying32  2017-6-15
package templatemsg

import (
	"encoding/json"
	"fmt"

	"github.com/ying32/wechat/context"
	"github.com/ying32/wechat/util"
)

const (
	templateMsg = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

// TTemplateMsg 业务模板消息
type TTemplateMsg struct {
	*context.Context
}

// TTemplateMsgBody 模板消息主导
type TTemplateMsgBody struct {
	ToUser     string      `json:"touser"`
	TemplateId string      `json:"template_id"`
	URL        string      `json:"url"`
	TopColor   string      `json:"topcolor"`
	Data       interface{} `json:"data"`
}

// TTemplateMsgError 请求返回结果
type TTemplateMsgError struct {
	util.CommonError
	MsgId int64 `json:"msgid"`
}

// TTemplateDataVal
type TTemplateDataVal struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TTemplateData struct {
	datas map[string]TTemplateDataVal
}

// NewTempldateData 新建一个模板数据
func NewTempldateData() *TTemplateData {
	td := new(TTemplateData)
	td.datas = make(map[string]TTemplateDataVal, 0)
	return td
}

// Add 添加模板数据
func (t *TTemplateData) Add(key, value, color string) {
	data := TTemplateDataVal{value, color}
	t.datas[key] = data
}

// Del 删除模板数据
func (t *TTemplateData) Del(key string) {
	delete(t.datas, key)
}

// Data 返回数据
func (t *TTemplateData) Data() interface{} {
	return t.datas
}

// NewTemplateMsg 构造
func NewTemplateMsg(context *context.Context) *TTemplateMsg {
	tmsg := new(TTemplateMsg)
	tmsg.Context = context
	return tmsg
}

// PushTo 推送业务消息
func (t *TTemplateMsg) PushTo(openId, templateId, url, topColor string, data *TTemplateData) error {
	accessKey, err := t.GetAccessToken()
	if err != nil {
		return err
	}
	var msgBody TTemplateMsgBody
	msgBody.ToUser = openId
	msgBody.TemplateId = templateId
	msgBody.TopColor = topColor
	msgBody.URL = url
	msgBody.Data = data.Data()
	response, err := util.PostJSON(fmt.Sprintf(templateMsg, accessKey), msgBody)
	var result TTemplateMsgError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return fmt.Errorf("Push template message error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
	}
	return nil
}

// PushTo2 推送业务消息， 使用默认的模板Id
func (t *TTemplateMsg) PushTo2(openId, url, topColor string, data *TTemplateData) error {
	return t.PushTo(openId, t.TemplateMsgId, url, topColor, data)
}
