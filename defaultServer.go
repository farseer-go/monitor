package monitor

import (
	"fmt"
	"github.com/farseer-go/fs/flog"
)

var defaultServer serverVO

// 服务端配置
type serverVO struct {
	Address string
}

// 随机一个服务端地址
func (receiver *serverVO) getAddress() string {
	count := len(receiver.Address)
	if count == 0 {
		flog.Panic("./farseer.yml配置文件没有找到Fops.Server.Address的设置")
	}
	return fmt.Sprintf("%s/ws/monitor", receiver.Address)
}
