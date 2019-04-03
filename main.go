package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gographics/imagick/imagick"
	"github.com/olahol/go-imageupload"
)

func main() {
	go server()

	forever := make(chan bool)
	<-forever
}

func server() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")

		if err != nil {
			panic(err)
		}

		mw := imagick.NewMagickWand()
		defer mw.Destroy()

		mw.ReadImageBlob(img.Data)

		mw.TrimImage(10)
		img.Data = mw.GetImageBlob()

		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
		if err != nil {
			panic(err)
		}

		thumb.Write(c.Writer)
	})

	r.Run(":5000")

}
