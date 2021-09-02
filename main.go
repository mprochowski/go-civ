package main

import (
	"flag"
	"fmt"
	"github.com/mccutchen/palettor"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main()  {
	filePath := flag.String("image", "","Image file path")
	flag.Parse()

	log.Println(*filePath)

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	i, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	c := getConsole()
	b := i.Bounds()
	log.Println(b.Max.X, b.Max.Y)
	log.Println(getConsole().Width, getConsole().Height)

	columns := c.Width
	lines := c.Height-2

	var proportion float32


	if columns < lines *2 {
		proportion = float32(b.Max.Y) / float32(columns)
		log.Println(proportion)
		//for y := 0; y < lines/2; y++ {
		//	for x := 0; x < (b.Max.X / int(proportion) *2) -1 ; x++ {
		//
		//		subImg := i.(SubImager).SubImage(image.Rect(x*int(proportion/2),int(float32(y)*proportion),x*int(proportion)/2+int(proportion), int(float32(y)*proportion+proportion))) // 0 0 0+4 y+4
		//		c, _ := palettor.Extract(1, 500, subImg)
		//
		//		cr, cg, cb, _ := c.Colors()[0].RGBA()
		//
		//		fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m", uint8(cr), uint8(cg), uint8(cb))
		//	}
		//	fmt.Printf("\n")
		//}
	} else {
		proportion = float32(b.Max.Y) / float32(lines)
		log.Println(proportion)
		for y := 0; y < lines; y++ {
			for x := 0; x < (b.Max.X / int(proportion) * 2) -1 ; x++ {

				subImg := i.(SubImager).SubImage(image.Rect(x*int(proportion/2),int(float32(y)*proportion),x*int(proportion)/2+int(proportion), int(float32(y)*proportion+proportion))) // 0 0 0+4 y+4
				c, _ := palettor.Extract(1, 500, subImg)

				cr, cg, cb, _ := c.Colors()[0].RGBA()

				fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m", uint8(cr), uint8(cg), uint8(cb))
			}
			fmt.Printf("\n")
		}
	}
}
