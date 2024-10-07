package historyImageMaker

type cHistoryData struct {
	dataList []int
}

var g_SingleHistoryData *cHistoryData = &cHistoryData{}

func GetSingleHistoryData() *cHistoryData {
	return g_SingleHistoryData
}

func (pInst *cHistoryData) PutData(value int) {
	pInst.dataList = append(pInst.dataList, value)
	iLen := len(pInst.dataList)
	if iLen > 10 {
		pInst.dataList = pInst.dataList[iLen-10:]
	}
}
func (pInst *cHistoryData) GetDataList() []int {
	return pInst.dataList
}
