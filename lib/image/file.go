package image

import (
	"image"
	"io/ioutil"
	"magick-app/lib/webp"
)

type File struct {
	Data    []byte `json:"data"`
	MimeTye string `json:"type"`
	Name    string `json:"name"`
	Image   image.Image
}

// jpeg/png => webp
func (f *File) Write() error {
	buf, err := webp.EncodeWebp(f.Image)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("convert/"+f.Name+".webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
