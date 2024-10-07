package diceHistoryImage

import historyImageMaker "github.com/ancestortelegram/diceHistoryImage/internal"

func DHI_Initialize(imageData []byte, width, height int, rowCount int) error {
	return historyImageMaker.GetSingleImageMaker().Initialize(imageData, width, height, rowCount)
}
func DHI_PutNumber(value int) {
	historyImageMaker.GetSingleHistoryData().PutData(value)
}
func DHI_DrawData() {
	historyImageMaker.GetSingleImageMaker().DrawData()
}
