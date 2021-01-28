package webp

import (
	"bytes"
	"image"

	"github.com/chai2010/webp"
)

// 将图像加载到缓冲区
func EncodeWebp(i image.Image) (buf bytes.Buffer, err error) {
	if err = webp.Encode(&buf, i, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		return buf, err
	}
	return buf, nil
}
