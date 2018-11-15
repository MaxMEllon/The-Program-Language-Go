package main

import (
	"flag"
	"fmt"
	"github.com/MaxMEllon/The-Program-Language-Go/ch10/ex01/img/gif"
	"github.com/MaxMEllon/The-Program-Language-Go/ch10/ex01/img/jpeg"
	"github.com/MaxMEllon/The-Program-Language-Go/ch10/ex01/img/png"
	"log"
	"os"
)

var contentType = flag.String("f", "jpg", "content type")

func main() {
	flag.Parse()

	var err error
	switch *contentType {
	case "jpg", "jpeg":
		err = jpeg.Convert(os.Stdin, os.Stdout)
	case "png":
		err = png.Convert(os.Stdin, os.Stdout)
	case "gif":
		err = gif.Convert(os.Stdin, os.Stdout)
	default:
		err = fmt.Errorf("unknown image format")
	}
	if err != nil {
		log.Fatal(err)
	}
}
