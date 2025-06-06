package monitor

import (
	"strconv"
	"time"

	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/utils/ws"
)

// SendContentVO monitor 发送实体
type SendContentVO struct {
	AppId   string                              // 项目Id
	AppName string                              // 项目名称
	Keys    collections.Dictionary[string, any] // 键值对
}

// 所有连接共享一个wsClient
var wsClientMonitor *ws.Client

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
			// 连接ws
			if !connectWs() {
				time.Sleep(3 * time.Second)
				continue
			}
			for {
				if dic := monitorFn(); dic.Count() > 0 {
					// 发送消息
					err := wsClientMonitor.Send(SendContentVO{
						AppId:   strconv.FormatInt(core.AppId, 10),
						AppName: core.AppName,
						Keys:    dic,
					})
					if err != nil {
						wsClientMonitor = nil
						flog.Warningf("[%s]监控发送消息失败：%s", core.AppName, err.Error())
						break
					}
				}
				time.Sleep(interval)
			}
			// 断开后重连
			time.Sleep(3 * time.Second)
		}
	}()
}

// connectWs 连接ws
func connectWs() bool {
	var err error
	if wsClientMonitor == nil || wsClientMonitor.IsClose() {
		address := defaultServer.getAddress()
		wsClientMonitor, err = ws.Connect(address, 8192)
		wsClientMonitor.AutoExit = false
		if err != nil {
			wsClientMonitor = nil
			flog.Warningf("[%s]wsmonitor连接fops失败：%s", core.AppName, err.Error())
		}
	}
	return wsClientMonitor != nil && !wsClientMonitor.IsClose()
}
