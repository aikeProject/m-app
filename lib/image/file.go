package image

import (
	"errors"
	"image"
	"io/ioutil"
	"magick-app/lib/webp"
	"os"
	"path"
)

type File struct {
	Data          []byte `json:"data"`
	Ext           string `json:"ext"`
	Id            string `json:"id"`
	MimeTye       string `json:"type"`
	Name          string `json:"name"`
	ConvertedFile string
	IsConverted   bool
	Image         image.Image
}

// 返回已转换文件的大小
func (f File) GetConvertedSize() (int64, error) {
	if f.ConvertedFile == "" {
		return 0, errors.New("文件没有对应的转换后的文件")
	}
	s, err := os.Stat(f.ConvertedFile)
	if err != nil {
		return 0, err
	}
	return s.Size(), nil
}

// jpeg/png => webp
func (f *File) Write(dir string) error {
	buf, err := webp.EncodeWebp(f.Image)
	dest := path.Join(dir, f.Name+".webp")

	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dest, buf.Bytes(), 0666); err != nil {
		return err
	}
	// 已转换文件的地址
	f.ConvertedFile = dest
	return nil
}
