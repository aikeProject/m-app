package png

import (
	"bytes"
	"image"
	"image/png"
	"io"

	"github.com/foobaz/lossypng/lossypng"
)

const qMax = 20

type Options struct {
	Quality int `json:"quality"`
}

// Decode a PNG file and return an image.
func DecodePNG(r io.Reader) (image.Image, error) {
	i, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// EncodePNG encodes an image into PNG and returns a buffer.
func EncodePNG(i image.Image, op *Options) (buf bytes.Buffer, err error) {
	c := lossypng.Compress(i, -1, qualityFactor(op.Quality))
	err = png.Encode(&buf, c)
	return buf, err
}

// qualityFactor normalizes the PNG quality factor from a max of 20, where 0 is
// no conversion.
func qualityFactor(q int) int {
	f := q / 100
	return qMax - (f * qMax)
}
