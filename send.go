package monitor

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/utils/ws"
	"time"
)

// 发送消息
func Send(appName, content interface{}) {
	address := defaultServer.getAddress()
	wsClient, err := ws.Connect(address, 8192)
	if err != nil {
		flog.Warningf("[%s]Fops连接失败：%s", appName, err.Error())
		time.Sleep(3 * time.Second)
	}
	// 发送消息
	wsClient.Send(content)
	// 关闭连接
	wsClient.IsClose()
}
