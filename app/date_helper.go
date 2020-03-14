package main

import (
	"strings"
	"time"
)

const thisYearDateLayout = "02 Jan, 15:04"
const anotherYearDateLayout = "02 Jan 2006, 15:04"

var localMonths = map[string]string{
	"янв": "Jan",
	"фев": "Feb",
	"мар": "Mar",
	"апр": "Apr",
	"май": "May",
	"июн": "Jun",
	"июл": "Jul",
	"авг": "Aug",
	"сен": "Sep",
	"окт": "Oct",
	"ноя": "Nov",
	"дек": "Dec",
}

func replaceMonth(dateString string) string {
	lowerDateString := strings.ToLower(dateString)
	for loc, eng := range localMonths {
		if index := strings.Index(lowerDateString, loc); index != -1 {
			return dateString[:index] + eng + dateString[index+len(loc):]
		}
	}
	return dateString
}

func parseTime(dateString string) (time.Time, error) {
	dateStringEng := replaceMonth(dateString)
	result, err1 := time.Parse(thisYearDateLayout, dateStringEng)
	if err1 != nil {
		return time.Parse(anotherYearDateLayout, dateStringEng)
	}
	return result.AddDate(time.Now().Year(), 0, 0), nil
}
