package monitor

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/utils/ws"
	"time"
)

// Send 发送消息
func Send(appName, content any) {
	address := defaultServer.getAddress()
	wsClient, err := ws.Connect(address, 8192)
	if err != nil {
		flog.Warningf("[%s]Fops连接失败：%s", appName, err.Error())
		time.Sleep(3 * time.Second)
	} else {
		// 发送消息
		err = wsClient.Send(content)
		if err != nil {
			flog.Warningf("[%s]监控发送消息失败：%s", appName, err.Error())
		}
		// 关闭连接
		wsClient.Close()
	}
}
