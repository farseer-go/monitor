package monitor

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/modules"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return nil
}

func (module Module) PreInitialize() {
	// 服务端配置
	defaultServer = serverVO{
		Address: configure.GetString("Fops.Server.Address"),
	}

	if len(defaultServer.Address) < 1 {
		panic("调度中心的地址[Fops.Server.Address]未配置")
	}

}
