package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"image-watermark/config"

	"image-watermark/fontwater"

	"github.com/gin-gonic/gin"
)

type Param struct {
	Url  string  `form:"url"`
	Text string  `form:"text"`
	Size float64 `form:"size"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// get params of image
		var param Param
		err := c.ShouldBind(&param)
		if err != nil {
			panic(err)
		}
		SetDefault(&param)
		GetImage(param.Url)
		watermark := fontwater.FontInfo{
			Size:     param.Size,
			Message:  param.Text,
			Position: config.TopLeft,
			Dx:       37,
			Dy:       37,
			R:        config.R,
			G:        config.G,
			B:        config.B,
			A:        config.A,
		}
		mutil := make([]fontwater.FontInfo, 3, 10)
		mutil = append(mutil, watermark)
		// make watermark
		if fontwater.StaticFontWater("src", "target.jpg", mutil) != nil {
			log.Panic(err)
		}
		// return image
		c.File("target.jpg")
	})
	router.Run(":9000")
}

func GetImage(imgUrl string) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("src", data, 0644)
}

func SetDefault(param *Param) {
	if param.Url == "" {
		param.Url = "https://pic.netbian.com/uploads/allimg/161001/095746-1475287066579f.jpg"
	}
	if param.Text == "" {
		param.Text = "watermark"
	}
	if param.Size == 0 {
		param.Size = 37
	}
}
