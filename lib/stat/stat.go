package stat

import (
	"encoding/json"
	"magick-app/lib/localstore"

	"github.com/wailsapp/wails"
)

const filename = "stat.json"

type Stat struct {
	ByteCount  int64 `json:"byteCount"`
	ImageCount int   `json:"imageCount"`
	Runtime    *wails.Runtime
	Logger     *wails.CustomLogger
	localStore *localstore.LocalStore
}

// 创建Stat实例
func NewStat() *Stat {
	s := &Stat{localStore: localstore.NewLocalStore()}
	// 加载本地配置文件
	d, _ := s.localStore.Load(filename)
	_ = json.Unmarshal(d, &s)
	return s
}

// 初始化
func (s *Stat) WailsInit(runtime *wails.Runtime) error {
	s.Runtime = runtime
	s.Logger = s.Runtime.Log.New("Stat")
	s.Logger.Info("Stat initialized...")
	return nil
}

// 获取"stat"配置
func (s *Stat) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"byteCount":  s.ByteCount,
		"imageCount": s.ImageCount,
	}
}

// 保存"ByteCount"
func (s *Stat) SetByteCount(b int64) {
	if b < 0 {
		return
	}
	s.ByteCount += b
	if err := s.store(); err != nil {
		s.Logger.Errorf("byteCount 保存失败: %v", err)
	}
}

// 保存"ImageCount"
func (s *Stat) SetImageCount(i int) {
	if i < 0 {
		return
	}
	s.ImageCount += i
	if err := s.store(); err != nil {
		s.Logger.Errorf("ImageCount 保存失败: %v", err)
	}
}

// 保存配置到本地文件
func (s *Stat) store() error {
	js, err := json.Marshal(s.GetStats())
	if err != nil {
		return err
	}
	if err := s.localStore.Store(js, filename); err != nil {
		return err
	}
	return nil
}
