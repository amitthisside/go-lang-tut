package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time handling in Golang!")

	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 Mon 15:04:05")) // the values passed are the default formatting values for date and time

	//creating our own time
	createdTime := time.Date(2003, time.December, 18, 11, 15, 8, 2, time.Local)
	fmt.Println(createdTime.Format("02-01-2006 Monday MST"))

	//Formats:
	// Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// Basic short date "2006/01/02" gives "2015/02/25"
	// AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
	// No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// 0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
	// 9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
}
