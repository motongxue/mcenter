package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/infraboard/mcenter/apps/instance/api"
	_ "github.com/infraboard/mcenter/apps/service/api"
)
