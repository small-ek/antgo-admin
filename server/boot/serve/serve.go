package serve

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/ant"
	_ "github.com/small-ek/antgo/frame/serve/gin"
	"server/boot/router"
)

// LoadSrv Load Api service<加载API服务>
func LoadSrv() {
	gin.ForceConsoleColor()
	// Load the configuration file<加载配置文件>
	configPath := flag.String("config", "./config/config.toml", "Configuration file path")

	flag.Parse()

	eng := ant.New(*configPath).Serve(router.Load())

	defer eng.Close()
}
