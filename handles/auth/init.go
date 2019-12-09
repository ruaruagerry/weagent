package auth

import "weagent/server"

func init() {
	server.RegisterPostHandleNoUserID("/auth/login", loginHandle)
}
