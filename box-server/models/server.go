package models

// import (
// 	"path/filepath"

// 	"github.com/fatih/color"
// 	"github.com/fsnotify/fsnotify"
// )
// type EventServer interface {
// 	Trig(event *fsnotify.Event)
// }
// type Server struct {
// }

// func (e Server) Trig(event *fsnotify.Event) {
// 	// 重命名事件可以被认定为删除操作的说明:
// 	// 1. mv to recycle             可定义为删除，文件确实不存在于监控目录了
// 	// 2. mv oldname newname oldDir 可定义为删除，新文件名会触发 create 事件
// 	const delEvent = fsnotify.Remove | fsnotify.Rename
// 	if event.Op&fsnotify.Write != 0 {
// 		// 有文件被写
// 		c.log("      发现服务修改: " + filepath.Base(event.Name))
// 		c.add(&CtrlTask{file: event.Name, opera: "update"})
// 	} else if event.Op&fsnotify.Create != 0 {
// 		// 有文件被创建
// 		c.log("      发现服务增加: " + filepath.Base(event.Name))
// 		c.add(&CtrlTask{file: event.Name, opera: "add"})
// 	} else if event.Op&delEvent != 0 {
// 		// 有文件被删除
// 		c.log("      发现服务删除: " + filepath.Base(event.Name))
// 		c.del(&CtrlTask{file: event.Name})
// 	} else {
// 		// 未知的操作
// 		c.log(
// 			color.RedString("      发现未处理的事件: !!!"),
// 			filepath.Base(event.Name),
// 			event.Op.String(),
// 		)
// 	}
// }
