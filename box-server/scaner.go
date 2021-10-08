package main

import (
	"path/filepath"

	"gitee.com/linakesi/home-cloud-server/models"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

type scaner struct {
	dir     string
	watcher *fsnotify.Watcher
}

func NewScaner(path string) *scaner {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	return &scaner{
		dir:     path,
		watcher: w,
	}
}

func (w *scaner) log(s ...string) {
	for _, l := range s {
		log.Println(color.CyanString("扫描:") + l)
	}
}

// 发现扫描目录中已经存在的服务并触发
func (w *scaner) Scanning(ts ...models.Trigger) {
	w.log("扫描目录:" + w.dir)
	files, err := filepath.Glob(w.dir + "/*.yaml")
	if err != nil {
		w.log(err.Error())
	}
	for _, f := range files {
		w.log("发现服务:" + filepath.Base(f))
		for _, t := range ts {
			t.Trig(&fsnotify.Event{
				Name: f,
				Op:   fsnotify.Create,
			})
		}
	}
}

// 发现扫描到的变化触发给传递进来的触发器
func (w *scaner) Watching(ts ...models.Trigger) {
	defer w.watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				// 监听到的所有事件都打印出来
				w.log(event.Op.String() + "\t: " + filepath.Base(event.Name))
				// 部分事件才进行触发操作：写入或创建或删除
				const writeOrCreateMask = fsnotify.Write | fsnotify.Create | fsnotify.Remove | fsnotify.Rename
				if event.Op&writeOrCreateMask != 0 {
					for _, t := range ts {
						t.Trig(&event)
					}
				}
			case err, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
				w.log("error", err.Error())
			}
		}
	}()

	err := w.watcher.Add(w.dir)
	if err != nil {
		log.Fatal(err)
	}
	w.log("监听目录:" + w.dir)
	<-done
}
