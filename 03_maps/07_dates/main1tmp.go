/*
На стандартный ввод подается строковое представление даты и времени в следующем
формате:

1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

Wed Apr 16 05:20:00 +0600 1986

Для более эффективной работы рекомендуется ознакомиться с документацией о
содержащихся в модуле time константах.

Sample Input:

1986-04-16T05:20:00+06:00
Sample Output:

Wed Apr 16 05:20:00 +0600 1986
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func checkDate(date string) (int, int, int) {
	date = strings.TrimSpace(date)
	splitten := strings.Split(date, "-")
	if len(splitten) != 3 {
		panic("invalid arguments")
	}
	year, month, day := splitten[0], splitten[1], splitten[2]
	year = strings.TrimSpace(year)
	month = strings.TrimSpace(month)
	day = strings.TrimSpace(day)
	if year == "" || month == "" || day == "" {
		panic("invalid arguments")
	}
	if len(year) != 4 || len(month) != 2 || len(day) != 2 {
		panic("invalid arguments")
	}
	yr, err := strconv.Atoi(year)
	if err != nil || yr < 0 {
		panic("year is invalid")
	}
	mn, err := strconv.Atoi(month)
	if err != nil || mn < 1 || mn > 12 {
		panic("month is invalid")
	}
	dy, err := strconv.Atoi(day)
	if err != nil || dy < 1 || dy > 31 {
		panic("day is invalid")
	}
	return yr, mn, dy
}

func checkTime(time string) (int, int, int, int, int, string) {
	time = strings.TrimSpace(time)
	sign := "+"
	if strings.Contains(time, "-") {
		sign = "-"
	}
	spl1 := strings.Split(time, sign)
	if len(spl1) != 2 {
		panic("invalid arguments")
	}
	fullTime, fullOffset := spl1[0], spl1[1]
	fullTime = strings.TrimSpace(fullTime)
	fullOffset = strings.TrimSpace(fullOffset)
	if fullTime == "" || fullOffset == "" || len(fullTime) != 8 || len(fullOffset) != 5 {
		panic("invalid arguments")
	}
	spl2 := strings.Split(fullTime, ":")
	if len(spl2) != 3 {
		panic("invalid arguments")
	}
	hour, minute, seconds := spl2[0], spl2[1], spl2[2]
	if hour == "" || minute == "" || seconds == "" {
		panic("invalid arguments")
	}
	hr, err := strconv.Atoi(hour)
	if err != nil || hr < 0 || hr > 23 {
		panic("hour is invalid")
	}
	mn, err := strconv.Atoi(minute)
	if err != nil || mn < 0 || mn > 59 {
		panic("minute is invalid")
	}
	sc, err := strconv.Atoi(seconds)
	if err != nil || sc < 0 || sc > 59 {
		panic("seconds is invalid")
	}
	spl3 := strings.Split(fullOffset, ":")
	if len(spl3) != 2 {
		panic("invalid arguments")
	}
	offsetHour, offsetMinute := spl3[0], spl3[1]
	if offsetHour == "" || offsetMinute == "" {
		panic("invalid arguments")
	}
	offHr, err := strconv.Atoi(offsetHour)
	if err != nil || offHr < 0 || offHr > 23 {
		panic("offset hour is invalid")
	}
	offMn, err := strconv.Atoi(offsetMinute)
	if err != nil || offMn < 0 || offMn > 59 {
		panic("offset minute is invalid")
	}
	return hr, mn, sc, offHr, offMn, sign
}

func main() {
	var src string = "1986-04-16T05:20:00-06:00"
	// fmt.Scan(&src)
	src = strings.TrimSpace(src)
	splitten := strings.Split(src, "T")
	if len(splitten) != 2 {
		panic("invalid arguments")
	}
	date, tim := splitten[0], splitten[1]
	if date == "" || tim == "" {
		panic("invalid arguments")
	}
	year, month, day := checkDate(date)
	hr, min, sec, offHr, offMn, sign := checkTime(tim)
	// fmt.Printf("year=%d, month=%d, day=%d\n", year, month, day)
	// fmt.Printf("hour=%d, min=%d, sec=%d, offHr=%d, offMn=%d\n", hr, min, sec, offHr, offMn)

	fixZoneName := fmt.Sprintf("UTC%s%d", sign, offHr)
	var loc *time.Location
	switch sign {
	case "-":
		loc = time.FixedZone(fixZoneName, -offHr*60*60)
	case "+":
		loc = time.FixedZone(fixZoneName, offHr*60*60)
	default:
		panic("invalid sign")
	}
	t := time.Date(year, time.Month(month), day, hr, min, sec, 1, loc)
	unixTime := time.Unix(t.Unix(), 1)
	fmt.Println(unixTime.Format(time.UnixDate))
}
