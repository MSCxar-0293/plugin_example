//Package example
package example

import {
	"math/rand"
	"strconv"
	"strings"

	control "guthb.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/Zerobot/message"

	"github.com/FloatTech/Zerobot-Plugin/order"
}

func init(){
	engine := control.Register("example",order.PrioChoose, &control.Options{
		DisableOnDefault: false,
		Help: "example\n" +
		"- hello, world! "
		"- 使用test以验证你的第一个插件。"
	})
	engine.OnFullMatch ("test").SetBlock(true).
	Handle(func(*zero.Ctx){
		ctx.SendChain(message.Text("hello, world! "))
	})
	engine.OnFullMatch ("hello, world!").SetBlock(true).
	Handle(func(*zero.Ctx){
		ctx.SendChain(message.Text("成功！"))
	})
	initChatList(func(){
		engine.OnFullMatchGroup(chatList, zero.OnlyToMe).SetBlock(true).Handle(
			func(ctx *zero.Ctx){
				key:=ctx.MessageString()
				val := *kimomap[key]
				text:=val[rand.Intn(len(val))]
				ctx.SendChain(message.Reply(ctx.Event.MessageID), message.Text(text))
			})
		})
}