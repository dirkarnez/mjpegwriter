package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/icza/mjpeg"
)

const (
	dx = 500
	dy = 200
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	aw, err := mjpeg.New("test.avi", dx, dy, 2)
	checkErr(err)

	// Acquire / initialize image, e.g.:
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			rgba.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}

	fmt.Println(rgba.At(400, 100))    //{144 100 0 255}
	fmt.Println(rgba.Bounds())        //(0,0)-(500,200)
	fmt.Println(rgba.Opaque())        //true，其完全透明
	fmt.Println(rgba.PixOffset(1, 1)) //2004
	fmt.Println(rgba.Stride)          //2000

	buf := &bytes.Buffer{}
	checkErr(jpeg.Encode(buf, rgba, nil))

	for i := 0; i < 500; i++ {
		checkErr(aw.AddFrame(buf.Bytes()))
	}

	checkErr(aw.Close())
}
