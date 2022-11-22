package helperFunc

import (
	"fmt"
	"strings"
)

func ModifyDate(needTime string) (string, error) {
	var needDateFormat string = ""
	needTimeSplit := strings.Split(needTime, "T")
	needTimeDateSplit := strings.Split(needTimeSplit[0], "-")
	needTimeTimeSplit := strings.Split(needTimeSplit[1], ".")
	needDateFormat = fmt.Sprintf("%s.%s.%s %s", needTimeDateSplit[2], needTimeDateSplit[1], needTimeDateSplit[0], needTimeTimeSplit[0])
	return needDateFormat, nil
}
