package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// вам это понадобится
const now = 1589570165

func getData() (int, int, int) {
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	x := strings.Split(s, " ")
	m, s := x[0], x[2]
	mins, err := strconv.Atoi(m)
	if err != nil {
		panic(err)
	}
	secs, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if secs < 0 || secs > 59 {
		panic("seconds must be in [0:59]")
	}
	if mins < 0 {
		panic("minutes must be positive")
	}
	hours := 0
	if mins > 59 {
		hours = mins / 60
		mins = mins % 60
	}
	return hours, mins, secs
}

func main() {
	hours, mins, secs := getData()
	dur, err := time.ParseDuration(fmt.Sprintf("%dh%dm%ds", hours, mins, secs))
	if err != nil {
		panic(err)
	}
	newTime := time.Unix(now, 0).UTC().Add(dur)
	fmt.Println(newTime.Format(time.UnixDate))
}
