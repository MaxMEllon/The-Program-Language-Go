package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"crypto/sha512"
	"encoding/hex"
)

func main() {
	t := flag.Int("type", 256, "256, 384, 512 が利用可能. 初期値: 256")
	flag.Parse()

	var h hash.Hash
	switch *t {
	case 256:
		h = sha256.New()
	case 384:
		h = sha512.New384()
	case 512:
		h = sha512.New()
	default:
		log.Fatalf("サポートされていない: %v", *t)
	}
	io.Copy(h, os.Stdin)
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
