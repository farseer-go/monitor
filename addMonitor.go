package monitor

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/utils/ws"
	"time"
)

// SendContentVO monitor 发送实体
type SendContentVO struct {
	AppId   string                              // 项目Id
	AppName string                              // 项目名称
	Keys    collections.Dictionary[string, any] // 键值对
}

// 每个应用对应的ClientVO
var wsClient *ws.Client

// monitorFunc 客户端要执行的monitor
type monitorFunc func() collections.Dictionary[string, any]

// AddMonitor 添加一个监控，运行前先休眠
// interval:任务运行的间隔时间
func AddMonitor(interval time.Duration, monitorFn monitorFunc) {
	// 不立即运行，则先休眠interval时间
	if interval <= 0 {
		panic("interval参数，必须大于0")
	}
	go func() {
		for {
			var err error
			address := defaultServer.getAddress()
			wsClient, err = ws.Connect(address, 8192)
			if err != nil {
				flog.Warningf("[%s]wsmonitor连接fops失败：%s", core.AppName, err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			for {
				dic := monitorFn()
				// 发送消息
				err = wsClient.Send(SendContentVO{
					AppId:   parse.ToString(core.AppId),
					AppName: core.AppName,
					Keys:    dic,
				})
				if err != nil {
					flog.Warningf("[%s]监控发送消息失败：%s", core.AppName, err.Error())
					break
				}
				time.Sleep(interval)
			}
			// 断开后重连
			time.Sleep(3 * time.Second)
		}
	}()
}
