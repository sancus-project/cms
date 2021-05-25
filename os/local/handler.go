package local

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/fsnotify/fsnotify"
)

func (v *Filesystem) String() string {
	return fmt.Sprintf("%s: %s:%s", "go.sancus.dev/cms", "local", v.root)
}

func (v *Filesystem) handleEvent(event fsnotify.Event) {
	log.Printf("%v event:%s", v, event)
}

func (v *Filesystem) handleError(err error) {
	log.Printf("%v error:%s", v, err)
}

func (v *Filesystem) scan(f fs.DirEntry) error {
	log.Printf("%v scan:%#v", v, f)
	return nil
}
