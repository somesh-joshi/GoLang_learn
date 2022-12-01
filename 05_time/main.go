package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) 
	// if you want to format the date-month-year then it is must in 01-02-2006 format
	// Don't worry it give right output

	fmt.Println(presentTime.Format("01-02-2006 15:04:05"))
	// if you want to format the date-month-year time then it is must in 01-02-2006 15:04:05 format
	// Don't worry it give right output

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// if you want to format the date-month-year time day then it is must in 01-02-2006 15:04:05 Monday format
	// Don't worry it give right output

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday January"))
	// if you want to format the date-month-year time day month then it is must in 01-02-2006 15:04:05 Monday January format
	// Don't worry it give right output

	//🤣🤣🤣

	// for mixtime
	fmt.Println(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
}
