package image

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"

	"github.com/chai2010/webp"
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
	file.Image, err = jpeg.Decode(bytes.NewReader(file.Data))
	if err != nil {
		return file, err
	}
	if err := Write(&file); err != nil {
		return file, err
	}
	return file, nil
}

// jpeg => webp
func Write(i *File) error {
	var buf bytes.Buffer
	//var data []byte
	var err error

	// Load file data
	//if data, err = ioutil.ReadFile("./test.jpg"); err != nil {
	//	log.Println(err)
	//}
	//
	//m, err := jpeg.Decode(bytes.NewReader(data))

	// Encode lossless webp
	if err = webp.Encode(&buf, i.Image, &webp.Options{Lossless: false, Quality: 70}); err != nil {
		return err
	}
	if err = ioutil.WriteFile("test.webp", buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
