package tasks

import (
	"log"

	"gitee.com/linakesi/home-cloud-server/conf"
	"github.com/fatih/color"
)

type Task struct {
	File    string
	Opera   string
	Service *conf.ServiceConf
}

func (t *Task) log(s ...string) {
	for _, l := range s {
		log.Println(color.GreenString("任务:") + l)
	}
}

func (t *Task) loadServiceConf() {
	var err error
	t.Service, err = conf.GetServiceConf(t.File)
	if err != nil {
		t.log(err.Error())
	}
}
