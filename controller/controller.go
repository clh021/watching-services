package controller

import (
	"log"
	"path/filepath"
	"strconv"
	"sync"

	"gitee.com/linakesi/home-cloud-server/conf"
	"gitee.com/linakesi/home-cloud-server/controller/generator"
	"gitee.com/linakesi/home-cloud-server/controller/operator"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

type CtrlTask struct {
	file    string
	opera   string
	service *conf.ServiceConf
}

func (t *CtrlTask) loadSericeConf() {
	t.service = conf.GetServiceConf(t.file)
}

type Controller struct {
	nginxGenerator *generator.Generator
	dockerOperator *operator.Operator
	listTask       []*CtrlTask
	listLocked     bool //避免判断列表中存在同服务变化任务时操作失败
	genWG          sync.WaitGroup
}

func New(g *generator.Generator, o *operator.Operator) *Controller {
	ng := &Controller{
		dockerOperator: o,
		nginxGenerator: g,
		listLocked:     false,
	}
	go ng.Start()
	return ng
}
func (c *Controller) prepare(wg *sync.WaitGroup, t *CtrlTask) {
	c.log("当前任务总数:" + strconv.Itoa(len(c.listTask)))
	c.log("处理任务[" + t.opera + "]: " + filepath.Base(t.file))

	for {
		// listLocked 避免判断列表中存在同服务变化任务时操作失败
		if !c.listLocked {
			c.listLocked = true
			t.loadSericeConf()
			c.del(t)
			c.listLocked = false
			wg.Done()
			break
		}
	}
}
func (c *Controller) Start() {
	for {
		if len(c.listTask) > 0 {
			c.genWG = sync.WaitGroup{}

			c.genWG.Add(1)
			// 配置准备
			t := c.listTask[0]
			c.prepare(&c.genWG, t)

			c.genWG.Add(1)
			// 检查 ctrlTaskQueue 是否有内容，有就取出第一个
			// 发送到处理队列 queue
			err := c.dockerOperator.Apply(&c.genWG, t.service)
			if err != nil {
				c.log(err.Error())
			}

			c.genWG.Add(1)
			err = c.nginxGenerator.Apply(&c.genWG, t.service)
			if err != nil {
				c.log(err.Error())
			}

			c.genWG.Wait()
		}
	}
}

func (c *Controller) del(t *CtrlTask) {
	for i, other := range c.listTask {
		if other.file == t.file {
			c.listTask = append(c.listTask[:i], c.listTask[i+1:]...)
			break
		}
	}
}

// 去除重复，并以最后一个任务为准
// (重复任务更新文件操作即可)
func (c *Controller) add(t *CtrlTask) {
	Index := c.exist(t)
	if -1 == Index {
		c.listTask = append(c.listTask, t)
	} else {
		c.listTask[Index].opera = t.opera
	}
}

func (c *Controller) exist(s *CtrlTask) int {
	for number, t := range c.listTask {
		if t.file == s.file {
			return number
		}
	}
	return -1
}

func (c *Controller) log(s ...string) {
	for _, l := range s {
		log.Println(color.GreenString("CTRL:") + l)
	}
}

func (c *Controller) Trig(event *fsnotify.Event) {
	// 重命名事件可以被认定为删除操作的说明:
	// 1. mv to recycle             可定义为删除，文件确实不存在于监控目录了
	// 2. mv oldname newname oldDir 可定义为删除，新文件名会触发 create 事件
	const delEvent = fsnotify.Remove | fsnotify.Rename
	if event.Op&fsnotify.Write != 0 {
		// 有文件被写
		c.log("      发现服务修改: " + filepath.Base(event.Name))
		c.add(&CtrlTask{file: event.Name, opera: "update"})
	} else if event.Op&fsnotify.Create != 0 {
		// 有文件被创建
		c.log("      发现服务增加: " + filepath.Base(event.Name))
		c.add(&CtrlTask{file: event.Name, opera: "add"})
	} else if event.Op&delEvent != 0 {
		// 有文件被删除
		c.log("      发现服务删除: " + filepath.Base(event.Name))
		c.del(&CtrlTask{file: event.Name})
	} else {
		// 未知的操作
		c.log(
			color.RedString("      发现未处理的事件: !!!"),
			filepath.Base(event.Name),
			event.Op.String(),
		)
	}
}
