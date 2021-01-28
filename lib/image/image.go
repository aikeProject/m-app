package image

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"magick-app/lib/webp"
)

type File struct {
	Data    []byte `json:"data"`
	MimeTye string `json:"type"`
	Name    string `json:"name"`
	Image   image.Image
}

func HandleFile(fileJson string) (file File, err error) {
	// 解析json
	if err := json.Unmarshal([]byte(fileJson), &file); err != nil {
		return file, err
	}
	if file.MimeTye == "image/jpg" {
		file.Image, err = jpeg.Decode(bytes.NewReader(file.Data))
	} else if file.MimeTye == "image/png" {
		file.Image, err = png.Decode(bytes.NewReader(file.Data))
	}

	if err != nil {
		return file, err
	}
	if err := file.Write(); err != nil {
		return file, err
	}
	return file, nil
}

// jpeg/png => webp
func (f *File) Write() error {
	buf, err := webp.EncodeWebp(f.Image)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("test.webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
