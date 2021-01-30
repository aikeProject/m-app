package image

import (
	"encoding/json"
	"fmt"
	"magick-app/lib/config"
	"path"
	"sync"

	"github.com/wailsapp/wails"
)

type FileManager struct {
	Config  *config.Config
	Files   []*File
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

func NewFileManager(c *config.Config) *FileManager {
	return &FileManager{
		Config: c,
	}
}

func (m *FileManager) WailsInit(runtime *wails.Runtime) error {
	m.Runtime = runtime
	m.Logger = m.Runtime.Log.New("FileManager")
	m.Logger.Info("FileManager 初始化...")
	return nil
}

// 处理客户端图片
// 根据图片类型，将客户端base64解析到缓冲区
func (m *FileManager) HandleFile(fileJson string) (err error) {
	file := &File{}
	// 解析json
	if err := json.Unmarshal([]byte(fileJson), &file); err != nil {
		return err
	}

	if err := file.Decode(); err != nil {
		return err
	}
	m.Files = append(m.Files, file)
	m.Logger.Infof("添加到文件管理器: %s, type: %s", file.Name, file.MimeTye)

	return nil
}

// 格式转换
func (m *FileManager) Convert() (errs []error) {
	var wg sync.WaitGroup
	wg.Add(m.CountUnconverted())

	for _, file := range m.Files {
		f := file
		go func(w *sync.WaitGroup) {
			err := f.Write(m.Config.OutDir)
			if err != nil {
				m.Logger.Error(fmt.Sprintf("文件转换失败: %s, %v", f.Name, err))
				errs = append(errs, fmt.Errorf("文件转换失败: %s", f.Name))
			} else {
				f.IsConverted = true
				m.Logger.Infof("转换成功: %s", path.Join(m.Config.OutDir, f.Name+".webp"))
				s, err := f.GetConvertedSize()
				if err != nil {
					m.Logger.Errorf("获取不到转换文件大小：%v", err)
				}
				m.Runtime.Events.Emit("conversion:complete", map[string]interface{}{
					"id":   f.Id,
					"size": s,
				})
			}
			w.Done()
		}(&wg)
	}
	wg.Wait()

	return errs
}

// 统计未完成转换的文件
func (m *FileManager) CountUnconverted() int {
	c := 0
	for _, file := range m.Files {
		if !file.IsConverted {
			c++
		}
	}
	return c
}
