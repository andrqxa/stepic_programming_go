package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var m, s int
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	duration, err := time.ParseDuration(fmt.Sprintf("%dm%ds", m, s))
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Unix(now, 0).Add(duration).Format(time.UnixDate))
}
