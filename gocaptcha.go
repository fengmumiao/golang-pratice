//https://gitee.com/longfei6671/gocaptcha
package main

import (
	"image"
	"image/color"
	"image/draw"
	"io"
	"image/png"
	"image/jpeg"
	"image/gif"
	"go.intra.xiaojukeji.com/gulfstream/go-common/third_party/src/github.com/pkg/errors"
	"flag"
	"time"
	"math/rand"
	"math"
)

var (
dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
r = rand.New(rand.NewSource(time.Now().UnixNano()));
FontFamily []string = make([]string,0);
)

const txtChars = "AaCcDdEeFfGgHhJjKkLMmNnPpQqRrSsTtUuVvWwXxYtZ2346789";

const (
	//图片样式
	ImageFormatPng = iota
	ImageFormatJpeg
	ImageFormatGif

	//验证码噪点强度
	CaptchaComplexLower = iota
	CaptchaComplexMedium
	CaptchaComplexHigh
)

type CaptchaImage struct {
	nrgba *image.NRGBA
	width int
	height int
	Complex int
}

//新建图像
func NewCaptchaImage(width int, height int, bgcolor color.RGBA) (*CaptchaImage, error)  {
	m := image.NewNRGBA(image.Rect(0,o,width,height));
	draw.Draw(m, m.Bounds(), &image.Uniform{bgcolor},image.ZP,draw.Src)
	return &CaptchaImage{
		nrgba: m,
		height:height,
		width:width,
	},nil
}

//保存图片
func (this *CaptchaImage) saceImage(w io.Writer, imageFormat int) error {
	switch imageFormat {
	case ImageFormatPng:
		return png.Encode(w, this.nrgba);
	case ImageFormatJpeg:
		return jpeg.Encode(w, this.nrgba, &jpeg.Options{100});
	case ImageFormatGif:
		return gif.Encode(w, this.nrgba, &gif.Options{NumColors:256})
	}
	return errors.New("image type error")
}

//添加较粗空白直线
func (captatcha *CaptchaImage) DrawHollowLine()(*CaptchaImage)  {
	first := (captatcha.width/20);
	end := first*19;
	lineColor := color.RGBA{R:245,G:250,B:251,A:255};
	x1 := float64(r.Intn(first))
	x2 := float64(r.Intn(first)+end)

	multiple := float64(r.Intn(5)+3)/float64(5);
	if (int(multiple*19)%3 == 0) {
		multiple = multiple*-1.0;
	}
	w := captatcha.height / 20;

	for ;x1 < x2; x1 ++{

		y := math.Sin(x1*math.Pi*multiple/float64(captatcha.width)) * float64(captatcha.height/3);

		if(multiple < 0){
			y = y + float64(captatcha.height/2);
		}
		captatcha.nrgba.Set(int(x1),int(y),lineColor);

		for i:=0;i<=w;i++{
			captatcha.nrgba.Set(int(x1),int(y)+i,lineColor);
		}
	}

	return captatcha;
}

