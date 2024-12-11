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
	//host := flag.String("ETCD_HOST", "127.0.0.1:2379", "etcd hosts")
	//username := flag.String("ETCD_USERNAME", "", "etcd username")
	//pwd := flag.String("ETCD_PASSWORD", "", "etcd password")
	//path := flag.String("ETCD_PATH", "/local/admin.toml", "etcd path")
	flag.Parse()

	eng := ant.New(*configPath).Serve(router.Load())
	//使用etcd作为配置中心
	//eng := ant.New().Etcd([]string{*host}, []string{*path}, *username, *pwd).Serve(router.Load())

	defer eng.Close()
}
