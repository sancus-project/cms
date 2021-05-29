package local

import (
	"context"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"

	"go.sancus.dev/cms/os/registry"
	"go.sancus.dev/cms/os/types"
)

type Filesystem struct {
	root string
	ctx  context.Context

	// watcher
	watcher *fsnotify.Watcher

	// sync
	mu     sync.Mutex
	wg     sync.WaitGroup
	closed bool
}

func NewFilesystem(ctx context.Context, root string) (types.Filesystem, error) {

	// validate root
	root, err := CleanRoot(root)
	if err != nil {
		return nil, err
	}

	dents, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	// Watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// Context
	if ctx == nil {
		ctx = context.Background()
	}

	v := &Filesystem{
		root:    root,
		ctx:     ctx,
		watcher: watcher,
	}

	// start worker
	go func() {
		defer v.Close()
		defer v.wg.Done()

		v.wg.Add(1)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				v.handleEvent(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				v.handleError(err)
			}
		}
	}()

	if err := watcher.Add(root); err != nil {
		defer v.Close()
		return nil, err
	}

	// start scanning
	go func() {
		defer v.wg.Done()
		v.wg.Add(1)

		for _, f := range dents {
			v.scan(f)

			if v.closed {
				break
			}
		}
	}()

	return v, nil
}

func (v *Filesystem) Close() error {
	v.mu.Lock()
	if v.closed {
		v.mu.Unlock()
		return nil
	}

	watcher := v.watcher
	v.watcher = nil
	v.closed = true
	v.mu.Unlock()

	watcher.Close()
	v.wg.Wait()

	return nil
}

func (v *Filesystem) Root() string {
	return v.root
}

func (v *Filesystem) Protocol() string {
	return "file"
}

func init() {
	registry.RegisterFilesystem("", NewFilesystem)
	registry.RegisterFilesystem("local:", NewFilesystem)
	registry.RegisterFilesystem("file://", NewFilesystem)
}
