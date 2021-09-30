package operator

import (
	"log"
	"sync"
	"time"

	"gitee.com/linakesi/home-cloud-server/conf"
	"github.com/fatih/color"
)

type Operator struct {
}

func New() *Operator {
	return &Operator{}
}

func (o *Operator) log(s ...string) {
	for _, l := range s {
		log.Println(color.YellowString("OPER:") + l)
	}
}

func (o *Operator) Apply(wg *sync.WaitGroup, s *conf.ServiceConf) error {
	o.log("docker容器 应用 开始")
	// 模拟 5秒钟后，操作完毕
	time.Sleep(time.Duration(5) * time.Second)
	o.log("docker容器 应用 完毕")
	// 模拟 5秒钟后，操作完毕
	wg.Done()
	// s.MapPorts = append(s.MapPorts, 8080)
	// s.MapPorts = append(s.MapPorts, 8081)
	return nil
}
