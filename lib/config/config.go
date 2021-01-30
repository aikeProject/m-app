package config

import (
	"fmt"
	"os"
	"path"

	"github.com/wailsapp/wails"
)

// 应用程序配置
type Config struct {
	OutDir  string
	Target  string
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

// 返回"Config"实例
func NewConfig() *Config {
	c := &Config{}
	dir, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("找不到当前用户的主目录 %s", err)
	}
	od := path.Join(dir, "optimus")
	if _, err := os.Stat(od); os.IsNotExist(err) {
		if err := os.Mkdir(od, 0666); err != nil {
			od = "./"
			_ = fmt.Errorf("无法创建默认输出目录:%s", err)
		}
	}
	c.OutDir = od
	c.Target = "webp"
	return c
}

func (c *Config) WailsInit(runtime *wails.Runtime) error {
	c.Runtime = runtime
	c.Logger = c.Runtime.Log.New("Config")
	c.Logger.Info("Config 初始化...")
	return nil
}

// 打开对话框，选择输出目录
func (c *Config) SetOutDir() string {
	dir := c.Runtime.Dialog.SelectDirectory()
	c.OutDir = dir
	c.Logger.Infof("输出目录: %s", dir)
	return c.OutDir
}

// 文件转换的目标类型
func (c *Config) setTarget(t string) {
	c.Target = t
	c.Logger.Infof("文件转换类型: %s", t)
}
