package webp

import (
	"bytes"
	"image"
	"io"

	"github.com/chai2010/webp"
)

// 解析webp图像数据
func DecodeWebp(r io.Reader) (image.Image, error) {
	d, err := webp.Decode(r)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// 将图像加载到缓冲区
func EncodeWebp(i image.Image) (buf bytes.Buffer, err error) {
	if err = webp.Encode(&buf, i, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		return buf, err
	}
	return buf, nil
}
