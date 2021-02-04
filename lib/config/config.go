package config

import (
	"encoding/json"
	"fmt"
	"magick-app/lib/jpeg"
	"magick-app/lib/localstore"
	"magick-app/lib/png"
	"magick-app/lib/webp"
	"os"
	"path"
	"path/filepath"

	"github.com/wailsapp/wails"
)

const filename = "conf.json"

// 本地配置
// OutDir 文件保存目录
// Target 文件目标类型 png/jpg/webp...
// Prefix 文件名前缀
// Suffix 文件名结尾，不是扩展名
type App struct {
	OutDir  string        `json:"outDir"`
	Target  string        `json:"target"`
	Prefix  string        `json:"prefix"`
	Suffix  string        `json:"suffix"`
	JpegOpt *jpeg.Options `json:"jpegOpt"`
	PngOpt  *png.Options  `json:"pngOpt"`
	WebpOpt *webp.Options `json:"webpOpt"`
}

// 应用程序配置
type Config struct {
	App        *App
	LocalStore *localstore.LocalStore
	Runtime    *wails.Runtime
	Logger     *wails.CustomLogger
}

// 返回"Config"实例
func NewConfig() *Config {
	c := &Config{}
	c.LocalStore = localstore.NewLocalStore()
	file, err := c.LocalStore.Load(filename)
	if err != nil {
		c.App, _ = defaults()
	}
	if err := json.Unmarshal(file, &c.App); err != nil {
		fmt.Printf("error %v", err)
	}
	return c
}

// 获取配置
func (c *Config) GetAppConfig() map[string]interface{} {
	return map[string]interface{}{
		"outDir":  c.App.OutDir,
		"target":  c.App.Target,
		"prefix":  c.App.Prefix,
		"suffix":  c.App.Suffix,
		"jpegOpt": c.App.JpegOpt,
		"pngOpt":  c.App.PngOpt,
		"webpOpt": c.App.WebpOpt,
	}
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
	if dir != "" {
		c.App.OutDir = dir
		c.Logger.Infof("输出目录: %s", dir)
		if err := c.store(); err != nil {
			c.Logger.Errorf("配置保存失败：%v", err)
		}
	}
	return c.App.OutDir
}

// 打开输出目录
func (c *Config) OpenOutputDir() error {
	if err := c.Runtime.Browser.OpenURL(c.App.OutDir); err != nil {
		return err
	}
	return nil
}

// SetConfig sets and stores the given configuration.
func (c *Config) SetConfig(cfg string) error {
	a := &App{}
	if err := json.Unmarshal([]byte(cfg), &a); err != nil {
		c.Logger.Errorf("failed to unmarshal config: %v", err)
		return err
	}
	c.App = a
	if err := c.store(); err != nil {
		c.Logger.Errorf("failed to store config: %v", err)
		return err
	}
	return nil
}

// 保存配置到配置文件
func (c *Config) store() error {
	js, err := json.Marshal(c.GetAppConfig())
	if err != nil {
		c.Logger.Errorf("应用配置解析失败: %v", err)
		return err
	}
	if err := c.LocalStore.Store(js, filename); err != nil {
		c.Logger.Errorf("应用配置保存失败: %v", err)
		return err
	}
	return nil
}

// 重置为默认配置
func (c Config) RestoreDefaults() (err error) {
	app, err := defaults()
	if err != nil {
		return err
	}
	c.App = app
	if err := c.store(); err != nil {
		return err
	}
	return nil
}

// 默认配置
func defaults() (*App, error) {
	a := &App{
		Target:  "webp",
		JpegOpt: &jpeg.Options{Quality: 80},
		PngOpt:  &png.Options{Quality: 80},
		WebpOpt: &webp.Options{Lossless: false, Quality: 80},
	}
	ud, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("failed to get user directory: %v", err)
		return nil, err
	}

	od := path.Join(ud, "Optimus")
	cp := filepath.Clean(od)

	if _, err := os.Stat(od); os.IsNotExist(err) {
		if err := os.Mkdir(od, 0777); err != nil {
			od = "./"
			fmt.Printf("failed to create default output directory: %v", err)
			return nil, err
		}
	}
	a.OutDir = cp
	return a, nil
}
