package test

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/monitor"
	"github.com/farseer-go/webapi"
	"testing"
)

type startupModule struct {
}

func (module startupModule) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}, monitor.Module{}}
}

func (module startupModule) PostInitialize() {
}

// 发送测试信息
func TestSend(t *testing.T) {
	fs.Initialize[startupModule]("test monitor")
	// 创建字典
	dic := collections.NewDictionary[string, string]()
	dic.Add("cpu", "35")
	dic.Add("store", "120")
	dic.Add("total", "0")
	monitor.Send(monitor.SendContentVO{
		AppId:   "app0001",
		AppName: "监控程序",
		Keys:    dic,
	})
}
