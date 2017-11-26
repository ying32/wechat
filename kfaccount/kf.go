// kfaccount 客服接口，主动动发送消息
// api: https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140547
package kfaccount

import (
	"encoding/json"
	"fmt"

	"github.com/ying32/wechat/context"
	"github.com/ying32/wechat/util"
)

const (
	kfCustomSend = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

type TKFAccount struct {
	*context.Context
}

func NewKFAccount(context *context.Context) *TKFAccount {
	kf := new(TKFAccount)
	kf.Context = context
	return kf
}

// CustomSend 主动发送消息给用户
func (k *TKFAccount) CustomSend(msg interface{}) error {
	token, err := k.GetAccessToken()
	if err != nil {
		return err
	}
	replyData, err := util.PostJSON(fmt.Sprintf(kfCustomSend, token), msg)
	if err != nil {
		return err
	}
	var result util.CommonError
	err = json.Unmarshal(replyData, &result)
	if err != nil {
		return err
	}
	if result.ErrCode == 0 {
		return nil
	}
	return fmt.Errorf(fmt.Sprintf("errorcode:%s, msg:%s", result.ErrCode, result.ErrMsg))
}
