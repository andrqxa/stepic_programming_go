package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func checkTime(time string) (int, int, int) {
	fullTime := strings.TrimSpace(time)
	if fullTime == "" || len(fullTime) != 8 {
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
	return hr, mn, sc
}

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	srcStr := string(src)
	x := strings.Split(srcStr, " ")
	date, tim := x[0], x[1]
	y := strings.Split(tim, ":")
	hr, err := strconv.Atoi(y[0])
	if err != nil {
		panic("invalid argument")
	}
	if hr < 13 {
		fmt.Println(srcStr)
		return
	}
	year, month, day := checkDate(date)
	hour, minute, seconds := checkTime(tim)
	srcTime := time.Date(year, time.Month(month), day, hour, minute, seconds, 0, time.UTC)
	future := srcTime.Add(time.Hour * 24)
	fmt.Println(future.Format("02-01-2006 15:04:05"))
}
