package jpeg

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

func Convert(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	if err := jpeg.Encode(out, img, &jpeg.Options{}); err != nil {
		return fmt.Errorf("jpeg: %v\n", err)
	}
	return nil
}
