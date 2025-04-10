package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println("Result Slice 1 : ", slice1)
	fmt.Println("Lenght Slice 1 :", len(slice1))
	fmt.Println("Capacity Slice 1 : ", cap(slice1))

	slice1[0] = "May Again"
	fmt.Println(slice1)
	fmt.Println("All Months : ", months)

	slice2 := months[10:]
	fmt.Println("Data Slice 2 of Months : ", slice2)

	slice3 := append(slice2, "New Year")
	fmt.Println("Append Of Months by slice 2 : ", slice3)

	slice3[1] = "Desember"
	fmt.Println("Months of slice 3 has changed : ", slice3)
	fmt.Println("Months of slice 2 : ", slice2)
	fmt.Println("All Months : ", months)

	fmt.Println("================================")

	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	day_slice1 := days[5:]
	fmt.Println("Day Slice of Slice", day_slice1)
	day_slice1[0] = "Sabtu"
	fmt.Println(day_slice1)
	fmt.Println("All Days : ", days)
	day_slice1[1] = "Weekend"
	fmt.Println(day_slice1)
	fmt.Println("All Days and Has Changed : ", days)

	day_slice2 := append(day_slice1, "Back To Work")
	fmt.Println("Schedule Append : ", day_slice2)

	fmt.Println("===========================")

	var new_slice = make([]string, 2, 5)
	new_slice[0] = "Muhammad"
	new_slice[1] = "Rivaldhi"

	fmt.Println("New Slice : ", new_slice)
	fmt.Println("Lenght of New Slice : ", len(new_slice))
	fmt.Println("Capacity of New Slice", cap(new_slice))

	fmt.Println("==========================================")

	var copy_slice = make([]string, len(new_slice), cap(new_slice))
	copy(copy_slice, new_slice)
	fmt.Println("Copy Slice by New Slice : ", copy_slice)

	// Different Array and Slice
	this_array := [5]int{1, 2, 3, 4, 5}
	this_slice := []int{1, 2, 3, 4, 5}

	fmt.Println("This Array : ", this_array)
	fmt.Println("This Slice : ", this_slice)
}