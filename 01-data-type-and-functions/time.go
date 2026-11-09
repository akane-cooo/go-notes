package ch01

import (
	"fmt"
	"time"
)

func basicTime() {
	t0 := time.Now()
	// time in second-scale
	fmt.Printf("second: %d, milli: %d, micro: %d, nano %d\n",
		t0.Unix(),
		t0.UnixMilli(),
		t0.UnixMicro(),
		t0.UnixNano(),
	)

	// duration (both methods are valid in getting duration since last time)
	// Duation is a struct
	duration1 := time.Since(t0)
	t1 := time.Now()

	duration2 := t1.Sub(t0)
	fmt.Printf("duration1 == duration2: %t\n", int(duration1.Seconds()) == int(duration2.Seconds()))

	// duration operations
	d := time.Duration(5 * time.Minute)
	t3 := t1.Add(d)
	fmt.Printf("t1 %d:%d; t2 %d:%d\n", t1.Hour(), t1.Minute(), t3.Hour(), t3.Minute())
}

func parseFormat() {
	// verbose way
	t := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d.%d\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.UnixMilli()%t.Unix())

	timeFormatA := "2006-01-02 15:04:05.000"
	fmt.Printf("%s\n", t.Format(timeFormatA))

	timeFormatB := "2006-1-2 15:4:5.000"
	fmt.Printf("%s\n", t.Format(timeFormatB))

	dateFmt := "01-02-2006"
	fmt.Printf("%s\n", t.Format(dateFmt))

	// dont use Parse, use ParseInLocation in production
	loc, _ := time.LoadLocation("America/New_York")
	parsedT, _ := time.ParseInLocation(timeFormatA, t.Format(timeFormatA), loc)
	fmt.Printf("%+v\n", parsedT)
}

func TimeRun() {
	// basicTime()
	parseFormat()
}
