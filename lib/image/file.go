package image

import (
	"bytes"
	"errors"
	"image"
	"io/ioutil"
	"magick-app/lib/config"
	"magick-app/lib/jpeg"
	"magick-app/lib/png"
	"magick-app/lib/webp"
	"os"
	"path"
	"path/filepath"
)

type File struct {
	Data          []byte `json:"data"`
	Ext           string `json:"ext"`
	Id            string `json:"id"`
	MimeTye       string `json:"type"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	ConvertedFile string
	IsConverted   bool
	Image         image.Image
}

var mimes = map[string]string{
	"image/.jpg": "jpg",
	"image/jpg":  "jpg",
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/webp": "webp",
}

// 返回文件类型
func getFileType(t string) (string, error) {
	s, prs := mimes[t]
	if !prs {
		_ = errors.New("不持之的文件类型:" + t)
	}
	return s, nil
}

// 根据不同文件类型进行解析
func (f *File) Decode() error {
	fileType, err := getFileType(f.MimeTye)
	if err != nil {
		return err
	}

	switch fileType {
	case "jpg":
		f.Image, err = jpeg.DecodeJPEG(bytes.NewReader(f.Data))
	case "png":
		f.Image, err = png.DecodePNG(bytes.NewReader(f.Data))
	case "webp":
		f.Image, err = webp.DecodeWebp(bytes.NewReader(f.Data))
	}

	return nil
}

// 返回已转换文件的大小
func (f *File) GetConvertedSize() (int64, error) {
	if f.ConvertedFile == "" {
		return 0, errors.New("没有对应的转换后的文件")
	}
	s, err := os.Stat(f.ConvertedFile)
	if err != nil {
		return 0, err
	}
	return s.Size(), nil
}

// 返回原始文件大小与转换后文件大小的增量
func (f *File) GetSavings() (int64, error) {
	size, err := f.GetConvertedSize()
	if err != nil {
		return 0, err
	}
	return size, nil
}

// jpeg/png => webp
func (f *File) Write(c *config.Config) (err error) {
	var buf bytes.Buffer

	switch c.App.Target {
	case "jpg":
		buf, err = jpeg.EncodeJPEG(f.Image, c.App.JpegOpt)
	case "png":
		buf, err = png.EncodePNG(f.Image, c.App.PngOpt)
	case "webp":
		buf, err = webp.EncodeWebp(f.Image, c.App.WebpOpt)
	}

	dest := path.Join(c.App.OutDir, c.App.Prefix+f.Name+c.App.Suffix+"."+c.App.Target)

	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(dest, buf.Bytes(), 0666); err != nil {
		return err
	}
	// 已转换文件的地址
	f.ConvertedFile = filepath.Clean(dest)
	f.IsConverted = true
	return nil
}
