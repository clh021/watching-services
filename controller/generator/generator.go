package generator

import (
	"log"
	"sync"
	"time"

	"gitee.com/linakesi/home-cloud-server/conf"
	"github.com/fatih/color"
)

type Generator struct {
	// 生成文件目的地
	dir string
}

func New(path string) *Generator {
	return &Generator{
		dir: path,
	}
}

func (g *Generator) log(s ...string) {
	for _, l := range s {
		log.Println(color.GreenString(" GEN:") + l)
	}
}

func (g Generator) Apply(wg *sync.WaitGroup, s *conf.ServiceConf) error {
	g.log("nginx 配置 生成 开始")
	// 模拟 5秒钟后，操作完毕
	time.Sleep(time.Duration(5) * time.Second)
	g.log("nginx 配置 生成 完毕")
	wg.Done()
	return nil
}
