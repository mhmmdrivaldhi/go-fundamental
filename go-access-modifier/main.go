package main

import (
	"access-modifier/lib"
	"fmt"
)

func main() {
	fmt.Println("Hour In a Day :", lib.HourInDay)
	fmt.Println("Days In Week :", lib.DayInWeek)

	fmt.Println(lib.DaysToMinutes(3))
}

