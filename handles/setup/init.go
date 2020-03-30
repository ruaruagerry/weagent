package setup

import (
	"weagent/server"
)

func init() {
	server.RegisterGetHandleNoUserID("/setup/real/get", realGetHandle)        // 获取实名认证信息
	server.RegisterPostHandleNoUserID("/setup/real/modify", realModifyHandle) // 修改实名认证信息
}
