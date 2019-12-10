package phone

import (
	"weagent/server"
)

func init() {
	server.RegisterGetHandle("/phone/getcode", getcodeHandle)             // 获取手机验证码
	server.RegisterGetHandle("/phone/bind", bindHandle)                   // 绑定手机号
	server.RegisterGetHandleNoUserID("/phone/wechat/getcode", bindHandle) // 公众号获取手机验证码
	server.RegisterGetHandleNoUserID("/phone/wechat/bind", bindHandle)    // 公众号绑定手机号
}
