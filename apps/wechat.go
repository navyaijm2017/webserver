package apps

import (
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/types"
	"github.com/luopengift/wechat"
)

var (
	wc    *wechat.WeChatCtx
	wcCfg *WeChatConfig
)

type WeChatConfig struct {
	CorpID         string `json:"corpID"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingAESKey"`
	AgentId        int    `json:"agentId"`
	Secret         string `json:"secret"`
}

type WeChatHandler struct {
	BaseHttpHandler
}

func (ctx *WeChatHandler) GET() {
	msg := ctx.GetQuery("msg_signature", "")
	timestamp := ctx.GetQuery("timestamp", "")
	nonce := ctx.GetQuery("nonce", "")
	echostr := ctx.GetQuery("echostr", "")
	logger.Info("%v,%v,%v,%v", msg, timestamp, nonce, echostr)
	msgCrypter, err := wechat.NewMessageCrypter(wcCfg.Token, wcCfg.EncodingAESKey, wcCfg.CorpID)
	if err != nil {
		logger.Error("%v", err)
	}
	msgDecrypt, corpID, err := msgCrypter.Decrypt(echostr)

	if err != nil {
		logger.Error("%v,%v,%v", msgDecrypt, corpID, err)
	}
	logger.Info("msgDecrypt=%v", string(msgDecrypt))
	ctx.Output(string(msgDecrypt))

}

func (ctx *WeChatHandler) POST() {
	logger.Info("%#v", ctx.GetBodyArgs())
}

type TextHandler struct {
	BaseHttpHandler
}

func (ctx *TextHandler) POST() {
	user, _ := types.ToString(ctx.GetBody("user"))
	message, _ := types.ToString(ctx.GetBody("message"))
	logger.Info("%v", message)
	qymsg := wechat.QYMessage{
		WeChatCtx: wc,
		ToUser:    user,
		ToParty:   "",
		ToTag:     "",
		MsgType:   "text",
		AgentId:   wcCfg.AgentId,
		Safe:      0,
		Text:      map[string]string{"content": message},
	}
	logger.Info("%#v", qymsg)
	ret := qymsg.SendText()
	ctx.Output(ret)
}

func WeChatInit(cfg *Config) {
	wcCfg = &cfg.WeChatConfig
	wc = wechat.NewWeChatCtx(cfg.WeChatConfig.CorpID, cfg.WeChatConfig.Secret)
	wc.SetType(wechat.QY)
	logger.Info("%v", wc.GetServerList())

}
