package gif

import (
	"fmt"
	"image"
	"image/gif"
	"io"
)

func Convert(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	if err := gif.Encode(out, img, &gif.Options{}); err != nil {
		return fmt.Errorf("gif: %v\n", err)
	}
	return nil
}
