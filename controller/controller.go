package controller

import (
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"

	"gitee.com/linakesi/home-cloud-server/controller/generator"
	"gitee.com/linakesi/home-cloud-server/controller/operator"
	"gitee.com/linakesi/home-cloud-server/controller/tasks"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

type Controller struct {
	nginxGenerator *generator.Generator
	dockerOperator *operator.Operator
	listTask       *tasks.TaskList
	genWG          sync.WaitGroup
}

func New(g *generator.Generator, o *operator.Operator) *Controller {
	tl := tasks.New()
	ng := &Controller{
		dockerOperator: o,
		nginxGenerator: g,
		listTask:       tl,
	}
	go tl.Working()
	go ng.Working()
	return ng
}

func (c *Controller) Working() {
	for {
		c.genWG = sync.WaitGroup{}

		c.genWG.Add(1)
		c.log("等待下一步要执行的任务")
		t := <-c.listTask.DoingTask
		c.log("取得任务:" + filepath.Base(t.File) + " " + t.Opera)
		err := c.dockerOperator.Apply(&c.genWG, t.Service)
		if err != nil {
			c.log(err.Error())
		}

		c.genWG.Add(1)
		err = c.nginxGenerator.Apply(&c.genWG, t.Service)
		if err != nil {
			c.log(err.Error())
		}

		c.genWG.Wait()
	}
}

func (c *Controller) log(s ...string) {
	for _, l := range s {
		log.Debugln(color.GreenString("服务:") + l)
	}
}

func (c *Controller) Trig(event *fsnotify.Event) {
	// 重命名事件可以被认定为删除操作的说明:
	// 1. mv to recycle             可定义为删除，文件确实不存在于监控目录了
	// 2. mv oldname newname oldDir 可定义为删除，新文件名会触发 create 事件
	const delEvent = fsnotify.Remove | fsnotify.Rename

	if event.Op&fsnotify.Write != 0 {
		// 有文件被写
		c.log("更新\t: " + filepath.Base(event.Name))
		c.listTask.Add(&tasks.Task{File: event.Name, Opera: "update"})

	} else if event.Op&fsnotify.Create != 0 {
		// 有文件被创建
		c.log("添加\t: " + filepath.Base(event.Name))
		c.listTask.Add(&tasks.Task{File: event.Name, Opera: "add"})

	} else if event.Op&delEvent != 0 {
		// 有文件被删除
		c.log("删除\t: " + filepath.Base(event.Name))
		c.listTask.Del(&tasks.Task{File: event.Name})

	} else {
		// 未知的操作
		c.log(
			color.RedString("Unkown Event !!!"),
			filepath.Base(event.Name),
			event.Op.String(),
		)
	}
}
