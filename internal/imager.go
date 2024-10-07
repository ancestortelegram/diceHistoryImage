package historyImageMaker

import (
	"bytes"
	"image"
	"image/jpeg"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/goregular"
)

type cImageMaker struct {
	//backGroup *image.RGBA
	bg        *gg.Context
	imgWidth  int
	imgHeight int
	rowCount  int
}

var g_singleImageMaker *cImageMaker = &cImageMaker{}

func GetSingleImageMaker() *cImageMaker {
	return g_singleImageMaker
}

func (pInst *cImageMaker) Initialize(imageData []byte, width, height int, rowCount int) error {
	img1, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return err
	}
	pInst.imgWidth = width
	pInst.imgHeight = height
	pInst.rowCount = rowCount

	pInst.makeBackGround(img1)

	return nil
}

func (pInst *cImageMaker) DrawData() ([]byte, error) {

	dc := gg.NewContextForImage(pInst.bg.Image())
	//dc.DrawImage(pInst.backGroup, 0, 0)

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 32})
	dc.SetFontFace(face)

	dataLen := len(GetSingleHistoryData().dataList)
	startPos := pInst.rowCount - dataLen
	if startPos < 0 {
		startPos = 0
	}

	wSpan := pInst.imgWidth / pInst.rowCount
	hSpan := pInst.imgHeight / 6
	for iLoop := startPos; iLoop < pInst.rowCount; iLoop++ {
		number1 := GetSingleHistoryData().dataList[iLoop-startPos]

		dc.DrawCircle(float64(iLoop*wSpan)+20, float64(hSpan*(7-number1))-15, 15)

		if number1 > 3 {
			dc.SetRGB(0.9, 0.3, 0.3)
		} else {
			dc.SetRGB(0.3, 0.3, 0.9)
		}
		dc.Fill()

		dc.SetRGB(1, 1, 1)

		dc.DrawString(strconv.Itoa(number1), float64(iLoop*wSpan)+12, float64(hSpan*(7-number1))-4)

	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, dc.Image(), nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (pInst *cImageMaker) makeBackGround(img1 image.Image) {
	dc := gg.NewContext(pInst.imgWidth, pInst.imgHeight)
	iLeft := (img1.Bounds().Dx() - pInst.imgWidth) / 2
	iTop := (img1.Bounds().Dy() - pInst.imgHeight) / 2
	//draw.Draw(m, m.Bounds(), img1, image.Point{iLeft, iTop}, draw.Src)
	//dc.DrawImageAnchored(img1, iLeft, iTop, 0, 0) //float64(pInst.imgWidth), float64(pInst.imgHeight))

	dc.DrawImage(img1, 0-iLeft, 0-iTop)
	dc.SetRGB(150, 150, 150)
	hSpan := pInst.imgHeight / 6
	dc.SetLineWidth(1)
	for iLoop := 1; iLoop < 6; iLoop++ {
		iTop := hSpan * iLoop
		dc.DrawLine(0, float64(iTop), float64(pInst.imgWidth), float64(iTop))
	}
	wSpan := pInst.imgWidth / pInst.rowCount
	for iLoop := 1; iLoop < pInst.rowCount; iLoop++ {
		iWidth := wSpan * iLoop
		dc.DrawLine(float64(iWidth), 0, float64(iWidth), float64(pInst.imgHeight))
	}
	dc.Stroke()

	pInst.bg = dc
}
