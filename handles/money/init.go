package money

import (
	"weagent/server"
)

func init() {
	server.RegisterGetHandleNoUserID("/money/ad/see", adSeeHandle)                // 查看广告上报
	server.RegisterGetHandleNoUserID("/money/ad/click", adClickHandle)            // 点击广告上报
	server.RegisterPostHandleNoUserID("/money/ad/record", adRecordHandle)         // 查看广告收益记录
	server.RegisterGetHandleNoUserID("/money/entrance", entranceHandle)           // 主界面
	server.RegisterPostHandleNoUserID("/money/getout/apply", getoutApplyHandle)   // 提现申请
	server.RegisterPostHandleNoUserID("/money/getout/record", getoutRecordHandle) // 查看提现记录
}
