#!/bin/bash

go run main.go -f jpeg < ./static/image.png > ./static/jpeg.jpg
go run main.go -f png < ./static/jpeg.jpg > ./static/png.png
go run main.go -f gif < ./static/png.png > ./static/gif.gif

open ./static/jpeg.jpg
open ./static/png.png
open ./static/gif.gif
