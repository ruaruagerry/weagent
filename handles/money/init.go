package money

import (
	"weagent/server"
)

func init() {
	server.RegisterGetHandle("/money/ad/see", adSeeHandle)                // 查看广告上报
	server.RegisterGetHandle("/money/ad/click", adClickHandle)            // 点击广告上报
	server.RegisterPostHandle("/money/ad/record", adRecordHandle)         // 广告收益记录
	server.RegisterGetHandle("/money/entrance", entranceHandle)           // 主界面
	server.RegisterGetHandle("/money/getout", getoutHandle)               // 提现
	server.RegisterPostHandle("/money/getout/record", getoutRecordHandle) // 提现记录
}
