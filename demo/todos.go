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
	watcher  *fsnotify.Watcher
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
	//t.runtime.Events.Emit("error", "来自Go程序的数据...")
	return string(bytes), err
}

// 保存列表
func (t *Todos) SaveList(todos string) error {
	t.logger.Infof("保存的数据为：%s", todos)
	return t.saveListByName(todos, t.filename)
}

// 保存
func (t *Todos) SaveAs(todos string) error {
	filename := t.runtime.Dialog.SelectSaveFile()
	t.logger.Info("Save As: " + filename)
	err := t.saveListByName(todos, filename)
	if err != nil {
		return err
	}
	return t.setFilename(filename)
}

func (t *Todos) saveListByName(todos string, filename string) error {
	return ioutil.WriteFile(filename, []byte(todos), 0600)
}

func (t *Todos) LoadNewList() error {
	var err error
	filename := t.runtime.Dialog.SelectFile()
	if len(filename) > 0 {
		err = t.setFilename(filename)
		t.runtime.Events.Emit("fileModified")
	}

	return err
}

// 监听文件变化
func (t *Todos) startWatcher() error {
	t.logger.Infof("开始监听...")
	watcher, err := fsnotify.NewWatcher()
	t.watcher = watcher
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

func (t *Todos) setFilename(filename string) error {
	var err error
	t.logger.Infof("old filename %s", t.filename)
	t.logger.Infof("new filename %s", filename)
	// Stop watching the current file and return any error
	err = t.watcher.Remove(t.filename)
	if err != nil {
		return err
	}

	// Set the filename
	t.filename = filename

	// Add the new file to the watcher and return any errors
	err = t.watcher.Add(filename)
	if err != nil {
		return err
	}
	t.logger.Info("Now watching: " + filename)
	t.runtime.Window.SetTitle(t.filename)
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
	t.runtime.Window.SetTitle(t.filename)
	t.ensureFileExists()
	return t.startWatcher()
}
