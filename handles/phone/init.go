package phone

import (
	"math/rand"
	"time"
	"weagent/server"
)

func init() {
	droprand = rand.New(rand.NewSource(time.Now().UnixNano()))

	server.RegisterGetHandleNoUserID("/phone/entrance", entranceHandle) // 手机验证入口
	server.RegisterGetHandleNoUserID("/phone/getcode", getcodeHandle)   // 获取手机验证码
	server.RegisterGetHandleNoUserID("/phone/bind", bindHandle)         // 绑定手机号
}
