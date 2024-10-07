package diceHistoryImage

import (
	"os"
	"testing"
)

func TestDHI_Initialize(t *testing.T) {

	data1, err := os.ReadFile("1.jpg")
	if err != nil {
		t.Log("file load error")
		return
	}

	err = DHI_Initialize(data1, 400, 200, 10)
	if err != nil {
		t.Log(err)
	}

	DHI_PutNumber(3)
	DHI_PutNumber(2)
	DHI_PutNumber(6)
	DHI_PutNumber(3)
	DHI_PutNumber(4)
	DHI_PutNumber(5)
	DHI_PutNumber(1)
	DHI_PutNumber(4)
	DHI_PutNumber(2)
	DHI_PutNumber(2)
	DHI_PutNumber(2)

	DHI_DrawData()

}
