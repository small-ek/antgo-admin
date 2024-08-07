package main

import "server/boot/serve"

//	@title			Swagger server API
//	@version		1.0
//	@description	server project

//	@contact.name	antgo
//	@contact.url	https://github.com/small-ek/antgo
//	@contact.email	56494565@qq.com

// @host		127.0.0.1:9001
// @BasePath	/api
func main() {
	serve.LoadSrv()
}
