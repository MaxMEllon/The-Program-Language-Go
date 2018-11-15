package img

import (
	"image"
	"io"
)

type Converter interface {
	Convert(in io.Reader, out io.Writer) error
}
