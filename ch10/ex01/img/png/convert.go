package png

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

func Convert(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	if err := png.Encode(out, img); err != nil {
		return fmt.Errorf("png: %v\n", err)
	}
	return nil
}
