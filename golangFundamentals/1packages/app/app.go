package main

import (
		"fmt"
		"time"
		)

func main(){
	start := time.Now()
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)


	fmt.Println(start)
	fmt.Println(afterTenSeconds)
	fmt.Println(afterTenMinutes)
	fmt.Println(afterTenHours)
	fmt.Println(afterTenDays)

	
	now := time.Now() 
	fmt.Println("Hello Martha")
	fmt.Println("Time : ", now)
	fmt.Println("Today : ", now.Format(time.ANSIC))
	fmt.Println("Day of the week : ", now.Weekday().String())
	nextYear := now.Add(365 * 24 * time.Hour)
	fmt.Println("Next year will be ", nextYear.Format(time.ANSIC))
	nextYearPlusOneDay := now.Add(366 * 24 * time.Hour)
	fmt.Println("Tomorrow a year from now the day will be ",nextYearPlusOneDay.Format(time.ANSIC))

}