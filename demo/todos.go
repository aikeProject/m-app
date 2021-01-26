package demo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails"
)

type Todos struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
}

// NewTodos attempts to create a new Todo list
func NewTodos() (*Todos, error) {
	return &Todos{}, nil
}

// 加载列表
func (t *Todos) LoadList() (string, error) {
	t.logger.Infof("列表数据来自于：%s", t.filename)
	bytes, err := ioutil.ReadFile(t.filename)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s ", t.filename)
	}
	// 发送消息
	t.runtime.Events.Emit("error", "来自Go程序的数据...")
	return string(bytes), err
}

// 保存列表
func (t *Todos) SaveList(todos string) error {
	t.logger.Infof("保存的数据为：%s", todos)
	return ioutil.WriteFile(t.filename, []byte(todos), 0600)
}

// 监听文件变化
func (t Todos) startWatcher() error {
	t.logger.Infof("开始监听...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case events, ok := <-watcher.Events:
				if !ok {
					return
				}
				if events.Op&fsnotify.Write == fsnotify.Write {
					t.logger.Infof("文件改变: %s", events.Name)
					t.runtime.Events.Emit("fileModified")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				t.logger.Error(err.Error())
			}
		}
	}()
	err = watcher.Add(t.filename)
	if err != nil {
		return err
	}
	return nil
}

// 文件不存在则创建一个
func (t *Todos) ensureFileExists() {
	// Check status of file
	_, err := os.Stat(t.filename)
	// If it doesn't exist
	if os.IsNotExist(err) {
		// Create it with a blank list
		err := ioutil.WriteFile(t.filename, []byte("[]"), 0600)
		if err != nil {
			t.logger.Error(err.Error())
		}
	}
}

// 获取"wails"的运行时环境"runtime"
func (t *Todos) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.logger = t.runtime.Log.New("Todos")
	t.logger.Info("todos is demo...")

	// Set the default filename to $HOMEDIR/todos.json
	homedir, err := runtime.FileSystem.HomeDir()
	if err != nil {
		return err
	}
	t.filename = path.Join(homedir, "Desktop/todos.json")

	t.ensureFileExists()
	return t.startWatcher()
}
