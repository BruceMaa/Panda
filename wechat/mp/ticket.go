package mp

import (
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"github.com/chanxuehong/wechat/json"
)

const (
	WechatTicketGetApi = `https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=%s` // 获取微信ticket地址
)

type Ticket struct {
	Ticket    string `json:"ticket"`     // 签名临时票据
	ExpiresIn string `json:"expires_in"` // 有效期，以秒为单位。在有效期内重复请求 ticket 不会被刷新
	common.PublicError
}

type TicketType string

const (
	WechatTicketTypeJsapi  TicketType = "jsapi"   // JSSDK获取ticket的type
	WechatTicketTypeWxcart TicketType = "wx_card" // 微信卡券
)

// 获取微信ticket
func (wm *WechatMp) GetTicket(accessToken string, ticketType TicketType) (*Ticket, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WechatTicketGetApi, accessToken, ticketType))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetTicket http get error: %+v\n", err)
		return nil, err
	}
	var ticket Ticket
	if err = json.Unmarshal(resp, &ticket); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetTicket json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &ticket, nil
}
