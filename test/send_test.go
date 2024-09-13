package test

import (
	"context"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/monitor"
	"github.com/farseer-go/webapi"
	"testing"
	"time"
)

type startupModule struct {
}

func (module startupModule) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}, monitor.Module{}}
}

func (module startupModule) PostInitialize() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	monitor.AddMonitor(5*time.Second, func() collections.Dictionary[string, any] {
		// 创建字典
		dic := collections.NewDictionary[string, any]()
		dic.Add("cpu", "35")
		dic.Add("store", "120")
		dic.Add("total", "0")
		return dic
	}, ctx)
}

// 发送测试信息
func TestSend(t *testing.T) {
	fs.Initialize[startupModule]("test monitor")
}
