package monitor

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/utils/ws"
	"time"
)

type SendContentVO struct {
	AppId   string                                 // 项目Id
	AppName string                                 // 项目名称
	Keys    collections.Dictionary[string, string] // 键值对
}

// Send 发送消息
func Send(content SendContentVO) {
	address := defaultServer.getAddress()
	wsClient, err := ws.Connect(address, 8192)
	if err != nil {
		flog.Warningf("[%s]Fops连接失败：%s", content.AppName, err.Error())
		time.Sleep(3 * time.Second)
	} else {
		// 发送消息
		err = wsClient.Send(content)
		if err != nil {
			flog.Warningf("[%s]监控发送消息失败：%s", content.AppName, err.Error())
		}
		// 关闭连接
		wsClient.Close()
	}
}
