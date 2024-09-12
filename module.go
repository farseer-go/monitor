package monitor

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/webapi"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}}
}

func (module Module) PreInitialize() {
	// 服务端配置
	defaultServer = serverVO{
		Address: configure.GetString("Fops.WsServer"),
	}

	if len(defaultServer.Address) < 1 {
		panic("调度中心的地址[Fops.WsServer]未配置")
	}

}
