package models

import "github.com/fsnotify/fsnotify"

type Trigger interface {
	Trig(event *fsnotify.Event)
}
