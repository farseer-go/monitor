package monitor

import (
	"strconv"
	"time"

	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
)

// Send 添加一个监控，直接发送消息
func Send(dic collections.Dictionary[string, any]) {
	if dic.Count() < 1 {
		return
	}
	for {
		// 连接ws
		if !connectWs() {
			time.Sleep(3 * time.Second)
			continue
		}
		// 发送消息
		if err := wsClientMonitor.Send(SendContentVO{AppId: strconv.FormatInt(core.AppId, 10), AppName: core.AppName, Keys: dic}); err != nil {
			wsClientMonitor = nil
			flog.Warningf("[%s]监控发送消息失败：%s", core.AppName, err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		return
	}
}

// Send 添加一个监控，直接发送消息
func SendValue(appName string, key string, value any) {
	dic := collections.NewDictionary[string, any]()
	dic.Add(key, value)

	for {
		// 连接ws
		if !connectWs() {
			time.Sleep(3 * time.Second)
			continue
		}
		// 发送消息
		if err := wsClientMonitor.Send(SendContentVO{AppId: strconv.FormatInt(core.AppId, 10), AppName: appName, Keys: dic}); err != nil {
			wsClientMonitor = nil
			flog.Warningf("[%s]监控发送消息失败：%s", core.AppName, err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		return
	}
}
