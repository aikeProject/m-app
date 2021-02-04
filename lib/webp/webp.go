package webp

import (
	"bytes"
	"image"
	"io"

	"github.com/chai2010/webp"
)

type Options struct {
	Lossless bool `json:"lossless"`
	Quality  int  `json:"quality"`
}

// 解析webp图像数据
func DecodeWebp(r io.Reader) (image.Image, error) {
	d, err := webp.Decode(r)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// 将图像加载到缓冲区
func EncodeWebp(i image.Image, op *Options) (buf bytes.Buffer, err error) {
	if err = webp.Encode(&buf, i, &webp.Options{Lossless: op.Lossless, Quality: float32(op.Quality)}); err != nil {
		return buf, err
	}
	return buf, nil
}
