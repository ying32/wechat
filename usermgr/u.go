// usermgr 用户管理
// api: https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839
package usermgr

import (
	"encoding/json"
	"fmt"

	"github.com/ying32/wechat/context"
	"github.com/ying32/wechat/util"
)

const (
	// userInfoURL 区别于oauth中的userInfoURL，此access_token为全局的那个不是通过code换的
	userInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

type TUserMgr struct {
	*context.Context
}

func NewUserMgr(context *context.Context) *TUserMgr {
	u := new(TUserMgr)
	u.Context = context
	return u
}

type TUserInfo struct {
	util.CommonError
	Subscribe     int    `json:"subscribe"`
	OpenId        string `json:"openid"`
	Nickname      string `json:"nickname"`
	Sex           int    `json:"sex"`
	Language      string `json"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	HeadImgURL    string `json:"headimgurl"`
	SubscribeTime int    `json:"subscribe_time"`
	Unionid       string `json"unionid"`
	Remark        string `json"remark"`
	GroupId       int    `json:"groupid"`
	TagidList     []int  `json"tagid_list"`
}

// GetUserInfo 获取用户信息
func (u *TUserMgr) GetUserInfo(openid string) (TUserInfo, error) {
	var result TUserInfo
	token, err := u.GetAccessToken()
	if err != nil {
		return result, err
	}
	replyData, err := util.HTTPGet(fmt.Sprintf(userInfoURL, token, openid))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(replyData, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode == 0 {
		return result, nil
	}
	return result, fmt.Errorf(fmt.Sprintf("GetUserInfo errorcode:%s, msg:%s", result.ErrCode, result.ErrMsg))
}
