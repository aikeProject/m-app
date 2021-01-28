package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"

	"github.com/wailsapp/wails"
)

type FileManager struct {
	Files   []*File
	Runtime *wails.Runtime
	Logger  *wails.CustomLogger
}

func NewFileManager() *FileManager {
	return &FileManager{}
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

	if file.MimeTye == "image/jpeg" || file.MimeTye == "image/jpg" {
		file.Image, err = jpeg.Decode(bytes.NewReader(file.Data))
		m.Files = append(m.Files, file)
		m.Logger.Infof("添加文件到管理器：%s", file.Name)
	} else if file.MimeTye == "image/png" {
		file.Image, err = png.Decode(bytes.NewReader(file.Data))
		m.Files = append(m.Files, file)
		m.Logger.Infof("添加文件到管理器：%s", file.Name)
	}

	if err != nil {
		return err
	}

	return nil
}

// 格式转换
func (m FileManager) Convert() error {
	for _, file := range m.Files {
		if err := file.Write(); err != nil {
			m.Logger.Error(fmt.Sprintf("文件转换失败: %s ", file.Name))
			return err
		}
		m.Logger.Infof("转换成功: %s", file.Name)
	}
	return nil
}
