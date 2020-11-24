package ddz

import (
	"github.com/spf13/cast"
)

func ToDz(i interface{}, preNum, sufNum int) string {
	str := cast.ToString(i)
	strNum := len(str)

	if preNum < 0 {
		preNum = 0
	}

	if sufNum < 0 {
		sufNum = 0
	}

	if preNum + sufNum > strNum {
		return str
	}

	mask := ""
	loopCount := strNum - preNum - sufNum
	for i := 0; i < loopCount; i++ {
		mask += "*"
	}

	return str[:preNum] + mask + str[strNum-sufNum:]
}
