package demo

import (
	"fmt"
	"github.com/wailsapp/wails"
	"io/ioutil"
	"os"
	"path"
)

type Todos struct {
	filename string
	runtime *wails.Runtime
	logger *wails.CustomLogger
}

// NewTodos attempts to create a new Todo list
func NewTodos() (*Todos, error) {
	// Create new Todos instance
	result := &Todos{}
	// Try and get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Join the cwd with our todos filename
	filename := path.Join(cwd, "myList.json")
	// Set the filename member of our new Todo list
	result.filename = filename
	// Return it
	return result, nil
}

func (t *Todos) LoadList() (string, error) {
	t.logger.Infof("列表数据来自于：%s", t.filename)
	bytes, err := ioutil.ReadFile(t.filename)
	if err != nil {
		err = fmt.Errorf("Unable to open list: %s ", t.filename)
	}
	// 发送消息
	t.runtime.Events.Emit("error","来自Go程序的数据...")
	return string(bytes), err
}

func (t *Todos) SaveList(todos string) error {
	t.logger.Infof("保存的数据为：%s", todos)
	return ioutil.WriteFile(t.filename, []byte(todos), 0600)
}

// 获取"wails"的运行时环境"runtime"
func (t *Todos) WailsInit(runtime *wails.Runtime) error  {
	t.runtime = runtime
	t.logger = t.runtime.Log.New("Todos")
	t.logger.Info("哈哈哈...")
	return nil
}
