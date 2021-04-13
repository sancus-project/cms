package local

import (
	"io/fs"
	"log"

	"github.com/fsnotify/fsnotify"
)

func (v *Filesystem) handleEvent(event fsnotify.Event) {
	log.Printf("local:%s: event:%s", v.root, event)
}

func (v *Filesystem) handleError(err error) {
	log.Printf("local:%s: error:%s", v.root, err)
}

func (v *Filesystem) scan(f fs.DirEntry) error {
	log.Printf("local:%s: scan:%q", v.root, f)
	return nil
}
