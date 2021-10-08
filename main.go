package main

import (
	"fmt"
	"os"

	"gitee.com/linakesi/home-cloud-server/conf"
	"gitee.com/linakesi/home-cloud-server/controller"
	"gitee.com/linakesi/home-cloud-server/controller/generator"
	"gitee.com/linakesi/home-cloud-server/controller/operator"
	"gitee.com/linakesi/home-cloud-server/models"
	log "github.com/sirupsen/logrus"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

	conf := *conf.GetConf()

	// 触发器
	var t models.Trigger

	// Nginx配置生成器
	g := generator.New(conf.NginxConfPath)
	o := operator.New()

	// 服务维护器
	t = controller.New(g, o)

	// 服务扫描器
	s := NewScaner(conf.ServicePath)
	s.Scanning(t) // 扫描所有已经存在的服务
	s.Watching(t) // 监听所有发生变化的服务
}
