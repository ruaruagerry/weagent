package phone

import (
	"math/rand"
	"time"
	"weagent/server"
)

func init() {
	droprand = rand.New(rand.NewSource(time.Now().UnixNano()))

	server.RegisterGetHandle("/phone/entrance", entranceHandle)           // 手机验证码入口
	server.RegisterGetHandle("/phone/getcode", getcodeHandle)             // 获取手机验证码
	server.RegisterGetHandle("/phone/bind", bindHandle)                   // 绑定手机号
	server.RegisterGetHandleNoUserID("/phone/wechat/getcode", bindHandle) // 公众号获取手机验证码
	server.RegisterGetHandleNoUserID("/phone/wechat/bind", bindHandle)    // 公众号绑定手机号
}
