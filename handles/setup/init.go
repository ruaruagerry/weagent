package setup

import (
	"weagent/server"
)

func init() {
	server.RegisterGetHandle("/setup/real/get", realGetHandle)
	server.RegisterPostHandle("/setup/real/modify", realModifyHandle)
}
