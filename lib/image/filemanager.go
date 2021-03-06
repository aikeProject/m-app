package image

import (
	"encoding/json"
	"fmt"
	"magick-app/lib/config"
	"magick-app/lib/stat"
	"path"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails"
)

type FileManager struct {
	Files   []*File
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
	config  *config.Config
	stats   *stat.Stat
}

func NewFileManager(c *config.Config, s *stat.Stat) *FileManager {
	return &FileManager{
		config: c,
		stats:  s,
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
	m.Logger.Infof("添加到文件管理器: id:%s, name: %s, type: %s", file.Id, file.Name, file.MimeTye)

	return nil
}

// 格式转换
func (m *FileManager) Convert() (errs []error) {
	var wg sync.WaitGroup
	wg.Add(m.countUnconverted())
	c := 0
	t := time.Now().UnixNano()
	var b int64
	for _, file := range m.Files {
		m.Logger.Infof("m.File: %s", file.Id)
		f := file
		if !f.IsConverted {
			go func(w *sync.WaitGroup) {
				err := f.Write(m.config)
				if err != nil {
					m.Logger.Error(fmt.Sprintf("文件转换失败: %s, %v", f.Name, err))
					errs = append(errs, fmt.Errorf("文件转换失败: %s", f.Name))
				} else {
					f.IsConverted = true
					m.Logger.Infof("转换成功: %s", path.Join(m.config.App.OutDir, f.Name+".webp"))
					s, err := f.GetConvertedSize()
					if err != nil {
						m.Logger.Errorf("获取不到转换文件大小：%v", err)
					}
					m.Runtime.Events.Emit("conversion:complete", map[string]interface{}{
						"id":   f.Id,
						"size": s,
						// 路径转换
						"path": strings.Replace(f.ConvertedFile, "\\", "/", -1),
					})
					c++
					s, err = f.GetSavings()
					if err != nil {
						m.Logger.Errorf("无法获取文件减少量：%v", err)
					}
					b += s
				}
				w.Done()
			}(&wg)
		}
	}
	wg.Wait()
	// 保存转换量
	m.stats.SetByteCount(b)
	// 保存转换的文件数
	m.stats.SetImageCount(c)
	m.Runtime.Events.Emit("conversion:stat", map[string]interface{}{
		"count":   c,
		"savings": b,
		"time":    (time.Now().UnixNano() - t) / 1000000,
	})
	m.Clear()
	return errs
}

// 统计未完成转换的文件
func (m *FileManager) countUnconverted() int {
	c := 0
	for _, file := range m.Files {
		if !file.IsConverted {
			c++
		}
	}
	return c
}

// 清空选择的文件
func (m *FileManager) Clear() {
	m.Files = nil
	debug.FreeOSMemory()
}

// 打开文件
func (m *FileManager) OpenFile(p string) error {
	if err := m.Runtime.Browser.OpenFile(p); err != nil {
		return err
	}
	return nil
}
