package tasks

import (
	"path/filepath"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"gitee.com/linakesi/home-cloud-server/conf"
	"github.com/fatih/color"
)

type Task struct {
	File    string
	Opera   string
	Service *conf.ServiceConf
}

func (t *Task) loadServiceConf() {
	t.Service = conf.GetServiceConf(t.File)
}

type TaskList struct {
	list      []*Task
	lock      sync.Mutex
	DoingTask chan *Task // 正在处理的任务
}

func New() *TaskList {
	return &TaskList{
		DoingTask: make(chan *Task),
	}
}

func (l *TaskList) Working() {
	for {
		if len(l.list) > 0 {

			l.lock.Lock()
			t := l.list[0]
			l.log("任务总数:" + strconv.Itoa(len(l.list)))
			l.log("处理任务:" + filepath.Base(t.File))
			l.deleteNoLock(t)
			l.lock.Unlock()

			t.loadServiceConf()
			l.DoingTask <- t
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func (l *TaskList) log(s ...string) {
	for _, l := range s {
		log.Println(color.GreenString("任务:") + l)
	}
}
func (l *TaskList) logOperate(t *Task, prefix string) {
	l.log(prefix + " " + t.Opera + " " + filepath.Base(t.File))
}

func (l *TaskList) Del(t *Task) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if -1 == l.existNoLock(t) {
		l.addNoLock(t)
		l.logOperate(t, "添加任务")
	} else {
		l.deleteNoLock(t)
		l.logOperate(t, "更新任务")
	}
	l.log("-----------------------------------")
	l.log("任务总数:" + strconv.Itoa(len(l.list)))
	l.log("处理任务:" + filepath.Base(t.File))
}

// 去除重复，并以最后一个任务为准
// (重复任务更新文件操作即可)
func (l *TaskList) Add(t *Task) {
	l.lock.Lock()
	defer l.lock.Unlock()
	Index := l.existNoLock(t)
	if -1 == Index {
		l.addNoLock(t)
	} else {
		l.updateNoLock(t, Index)
	}
}

// WARN: No listLock limit, 需要调用处获取TaskList锁
func (l *TaskList) deleteNoLock(t *Task) {
	for i, other := range l.list {
		if other.File == t.File {
			l.list = append(l.list[:i], l.list[i+1:]...)
			break
		}
	}
}

// WARN: No listLock limit, 需要调用处获取TaskList锁
func (l *TaskList) addNoLock(t *Task) {
	l.list = append(l.list, t)
	l.logOperate(t, "添加任务")
}

// WARN: No listLock limit, 需要调用处获取TaskList锁
func (l *TaskList) updateNoLock(t *Task, Index int) {
	l.list[Index].Opera = t.Opera
	l.logOperate(t, "更新任务")
}

// WARN: No listLock limit, 需要调用处获取TaskList锁
func (l *TaskList) existNoLock(t *Task) int {
	Index := -1
	for i, l := range l.list {
		if l.File == t.File {
			Index = i
			break
		}
	}
	return Index
}
